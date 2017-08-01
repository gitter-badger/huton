package huton

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"path/filepath"
	"sync"
	"time"

	"github.com/boltdb/bolt"
	"github.com/hashicorp/raft"
	"github.com/hashicorp/raft-boltdb"
	"github.com/hashicorp/serf/serf"
	"github.com/juju/errors"
	"google.golang.org/grpc"
)

const (
	raftRemoveGracePeriod = 5 * time.Second
)

var (
	// ErrNoName is an error used when and instance name is not provided
	ErrNoName = errors.New("no instance name provided")
)

// Instance is an interface for the Huton instance.
type Instance interface {
	// Bucket returns the bucket in the off-heap database with the given name.
	// If the bucket doesn't exist, it is automatically created.
	Bucket(name string) (Bucket, error)
	// Peers returns the current list of cluster peers. The list includes the local peer.
	Peers() []*Peer
	// Local returns the local peer.
	Local() *Peer
	// IsLeader returns true if this instance is the cluster leader.
	IsLeader() bool
	// Join joins and existing cluster.
	Join(addrs []string) (int, error)
	// Leave gracefully leaves the cluster.
	Leave() error
	// Shutdown forcefully shuts down the instance.
	Shutdown() error
}

type instance struct {
	name             string
	serf             *serf.Serf
	raft             *raft.Raft
	raftBoltStore    *raftboltdb.BoltStore
	raftTransport    *raft.NetworkTransport
	dbFilePath       string
	db               *bolt.DB
	rpcListener      net.Listener
	rpc              *grpc.Server
	serfEventChannel chan serf.Event
	shutdownCh       chan struct{}
	peersMu          sync.Mutex
	peers            map[string]*Peer
	dbMu             sync.Mutex
	config           *Config
	buckets          map[string]*bucket
	logger           *log.Logger
	raftNotifyCh     chan bool
	shutdownLock     sync.Mutex
	shutdown         bool
}

func (i *instance) Bucket(name string) (Bucket, error) {
	i.dbMu.Lock()
	defer i.dbMu.Unlock()
	if c, ok := i.buckets[name]; ok {
		return c, nil
	}
	return newBucket(i.db, name, i)
}

func (i *instance) Join(addrs []string) (int, error) {
	return i.serf.Join(addrs, true)
}

func (i *instance) IsLeader() bool {
	if i.raft == nil {
		return false
	}
	return i.raft.State() == raft.Leader
}

func (i *instance) Shutdown() error {
	i.logger.Println("Shutting down instance...")
	i.shutdownLock.Lock()
	defer i.shutdownLock.Unlock()
	if i.shutdown {
		return nil
	}
	i.shutdown = true
	close(i.shutdownCh)
	if i.serf != nil {
		i.serf.Shutdown()
	}
	if i.raft != nil {
		i.raftTransport.Close()
		i.raft.Shutdown().Error()
		if i.raftBoltStore != nil {
			i.raftBoltStore.Close()
		}
	}
	if i.rpcListener != nil {
		if err := i.rpcListener.Close(); err != nil {
			return fmt.Errorf("Failed to close RPC Listener: %s", err)
		}
	}
	if i.db != nil {
		if err := i.db.Close(); err != nil {
			return fmt.Errorf("Failed to close datastore: %s", err)
		}
	}
	return nil
}

func (i *instance) Leave() error {
	numPeers, err := i.numPeers()
	if err != nil {
		return err
	}
	addr := i.raftTransport.LocalAddr()
	isLeader := i.IsLeader()
	if isLeader && numPeers > 1 {
		future := i.raft.RemoveServer(raft.ServerID(i.name), 0, 0)
		if err := future.Error(); err != nil {
			i.logger.Printf("[ERR] failed to remove ourself as raft peer: %v", err)
		}
	}
	if i.serf != nil {
		if err := i.serf.Leave(); err != nil {
			i.logger.Printf("[ERR] failed to leave serf cluster: %v", err)
		}
	}
	// If we were not leader, wait to be safely removed from the cluster. We
	// must wait to allow the raft replication to take place, otherwise an
	// immediate shutdown could cause a loss of quorum.
	if !isLeader {
		var left bool
		limit := time.Now().Add(raftRemoveGracePeriod)
		for !left && time.Now().Before(limit) {
			time.Sleep(50 * time.Millisecond)
			future := i.raft.GetConfiguration()
			if err := future.Error(); err != nil {
				i.logger.Printf("[ERR] failed to get raft configuration: %v", err)
				break
			}
			left = true
			for _, server := range future.Configuration().Servers {
				if server.Address == addr {
					left = false
					break
				}
			}
		}
		if !left {
			i.logger.Printf("[WARN] failed to leave raft configuration gracefully, timeout")
		}
	}
	return nil
}

func (i *instance) numPeers() (int, error) {
	future := i.raft.GetConfiguration()
	if err := future.Error(); err != nil {
		return 0, err
	}
	configuration := future.Configuration()
	return len(configuration.Servers), nil
}

// NewInstance creates a new Huton instance and initializes it and all of its sub-components, such as Serf, Raft, and
// GRPC server, with the provided configuration.
//
// If this function returns successfully, the instance should be considered started and ready for use.
func NewInstance(name string, config *Config) (Instance, error) {
	if name == "" {
		return nil, ErrNoName
	}
	if config.LogOutput == nil {
		config.LogOutput = ioutil.Discard
	}
	i := &instance{
		name:       name,
		shutdownCh: make(chan struct{}),
		peers:      make(map[string]*Peer),
		config:     config,
		buckets:    make(map[string]*bucket),
		dbFilePath: filepath.Join(config.BaseDir, name, "store.db"),
		logger:     log.New(config.LogOutput, "huton", log.LstdFlags),
	}
	i.logger.Println("Initializing datastore...")
	if err := i.setupDB(); err != nil {
		return i, err
	}
	i.logger.Println("Initializing RPC server...")
	if err := i.setupRPC(); err != nil {
		i.Shutdown()
		return i, err
	}
	i.logger.Println("Initializing Raft cluster...")
	if err := i.setupRaft(); err != nil {
		i.Shutdown()
		return i, err
	}
	ip := net.ParseIP(config.BindAddr)
	raftAddr := &net.TCPAddr{
		IP:   ip,
		Port: config.BindPort + 1,
	}
	rpcAddr := &net.TCPAddr{
		IP:   ip,
		Port: config.BindPort + 2,
	}

	i.logger.Println("Initializing Serf cluster...")
	if err := i.setupSerf(raftAddr, rpcAddr); err != nil {
		i.Shutdown()
		return i, err
	}
	go i.handleEvents()
	return i, nil
}

func (i *instance) setupDB() error {
	timeout, err := time.ParseDuration(i.config.CacheDBTimeout)
	if err != nil {
		return fmt.Errorf("Failed to parse duration: %s", err)
	}
	cachesDB, err := bolt.Open(i.dbFilePath, 0644, &bolt.Options{
		Timeout: timeout,
	})
	if err != nil {
		return fmt.Errorf("Failed to open DB file %s: %s", i.dbFilePath, err)
	}
	i.db = cachesDB
	return nil
}

func (i *instance) handleEvents() {
	for {
		select {
		case e := <-i.config.SerfEventChannel:
			i.handleSerfEvent(e)
		case <-i.serf.ShutdownCh():
			i.Shutdown()
		case <-i.shutdownCh:
			return
		}
	}
}
