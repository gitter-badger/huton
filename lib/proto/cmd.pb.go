// Code generated by protoc-gen-go.
// source: cmd.proto
// DO NOT EDIT!

package huton_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Command struct {
	Type             *uint32 `protobuf:"varint,1,req,name=type" json:"type,omitempty"`
	Body             []byte  `protobuf:"bytes,2,req,name=body" json:"body,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Command) GetType() uint32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *Command) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Response struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func init() {
	proto.RegisterType((*Command)(nil), "huton_proto.Command")
	proto.RegisterType((*Response)(nil), "huton_proto.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Reciever service

type RecieverClient interface {
	OnCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Response, error)
}

type recieverClient struct {
	cc *grpc.ClientConn
}

func NewRecieverClient(cc *grpc.ClientConn) RecieverClient {
	return &recieverClient{cc}
}

func (c *recieverClient) OnCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/huton_proto.Reciever/OnCommand", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Reciever service

type RecieverServer interface {
	OnCommand(context.Context, *Command) (*Response, error)
}

func RegisterRecieverServer(s *grpc.Server, srv RecieverServer) {
	s.RegisterService(&_Reciever_serviceDesc, srv)
}

func _Reciever_OnCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecieverServer).OnCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/huton_proto.Reciever/OnCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecieverServer).OnCommand(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

var _Reciever_serviceDesc = grpc.ServiceDesc{
	ServiceName: "huton_proto.Reciever",
	HandlerType: (*RecieverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnCommand",
			Handler:    _Reciever_OnCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cmd.proto",
}

func init() { proto.RegisterFile("cmd.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 127 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0xce, 0x4d, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xce, 0x28, 0x2d, 0xc9, 0xcf, 0x8b, 0x07, 0x73, 0x94,
	0x54, 0xb9, 0xd8, 0x9d, 0xf3, 0x73, 0x73, 0x13, 0xf3, 0x52, 0x84, 0x78, 0xb8, 0x58, 0x4a, 0x2a,
	0x0b, 0x52, 0x25, 0x18, 0x15, 0x98, 0x34, 0x78, 0x41, 0xbc, 0xa4, 0xfc, 0x94, 0x4a, 0x09, 0x26,
	0x05, 0x26, 0x0d, 0x1e, 0x25, 0x2e, 0x2e, 0x8e, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54,
	0x23, 0x37, 0x10, 0x3b, 0x39, 0x33, 0xb5, 0x2c, 0xb5, 0x48, 0xc8, 0x8a, 0x8b, 0xd3, 0x3f, 0x0f,
	0x66, 0x80, 0x88, 0x1e, 0x92, 0xc9, 0x7a, 0x50, 0x51, 0x29, 0x51, 0x14, 0x51, 0x98, 0x29, 0x4a,
	0x0c, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x93, 0x16, 0xfd, 0x93, 0x00, 0x00, 0x00,
}
