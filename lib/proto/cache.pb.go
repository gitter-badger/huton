// Code generated by protoc-gen-go.
// source: cache.proto
// DO NOT EDIT!

/*
Package huton_proto is a generated protocol buffer package.

It is generated from these files:
	cache.proto
	cmd.proto

It has these top-level messages:
	CachePutCommand
	CacheDeleteCommand
	Command
	Response
*/
package huton_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CachePutCommand struct {
	CacheName        *string `protobuf:"bytes,1,req,name=cache_name" json:"cache_name,omitempty"`
	Key              []byte  `protobuf:"bytes,2,req,name=key" json:"key,omitempty"`
	Value            []byte  `protobuf:"bytes,3,req,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CachePutCommand) Reset()                    { *m = CachePutCommand{} }
func (m *CachePutCommand) String() string            { return proto.CompactTextString(m) }
func (*CachePutCommand) ProtoMessage()               {}
func (*CachePutCommand) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CachePutCommand) GetCacheName() string {
	if m != nil && m.CacheName != nil {
		return *m.CacheName
	}
	return ""
}

func (m *CachePutCommand) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *CachePutCommand) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type CacheDeleteCommand struct {
	CacheName        *string `protobuf:"bytes,1,req,name=cache_name" json:"cache_name,omitempty"`
	Key              []byte  `protobuf:"bytes,2,req,name=key" json:"key,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CacheDeleteCommand) Reset()                    { *m = CacheDeleteCommand{} }
func (m *CacheDeleteCommand) String() string            { return proto.CompactTextString(m) }
func (*CacheDeleteCommand) ProtoMessage()               {}
func (*CacheDeleteCommand) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CacheDeleteCommand) GetCacheName() string {
	if m != nil && m.CacheName != nil {
		return *m.CacheName
	}
	return ""
}

func (m *CacheDeleteCommand) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func init() {
	proto.RegisterType((*CachePutCommand)(nil), "huton_proto.CachePutCommand")
	proto.RegisterType((*CacheDeleteCommand)(nil), "huton_proto.CacheDeleteCommand")
}

func init() { proto.RegisterFile("cache.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 121 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x4e, 0x4c, 0xce,
	0x48, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xce, 0x28, 0x2d, 0xc9, 0xcf, 0x8b, 0x07,
	0x73, 0x94, 0x1c, 0xb9, 0xf8, 0x9d, 0x41, 0x72, 0x01, 0xa5, 0x25, 0xce, 0xf9, 0xb9, 0xb9, 0x89,
	0x79, 0x29, 0x42, 0x42, 0x5c, 0x5c, 0x60, 0xe5, 0xf1, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a,
	0x4c, 0x1a, 0x9c, 0x42, 0xdc, 0x5c, 0xcc, 0xd9, 0xa9, 0x95, 0x12, 0x4c, 0x0a, 0x4c, 0x1a, 0x3c,
	0x42, 0xbc, 0x5c, 0xac, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x12, 0xcc, 0x20, 0xae, 0x92, 0x29, 0x97,
	0x10, 0xd8, 0x08, 0x97, 0xd4, 0x9c, 0xd4, 0x92, 0x54, 0x62, 0x4d, 0x01, 0x04, 0x00, 0x00, 0xff,
	0xff, 0xa6, 0x47, 0xfd, 0x77, 0x94, 0x00, 0x00, 0x00,
}
