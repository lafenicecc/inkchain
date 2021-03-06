// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/peer.proto

package peer

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

type PeerID struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *PeerID) Reset()                    { *m = PeerID{} }
func (m *PeerID) String() string            { return proto.CompactTextString(m) }
func (*PeerID) ProtoMessage()               {}
func (*PeerID) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *PeerID) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PeerEndpoint struct {
	Id      *PeerID `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Address string  `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
}

func (m *PeerEndpoint) Reset()                    { *m = PeerEndpoint{} }
func (m *PeerEndpoint) String() string            { return proto.CompactTextString(m) }
func (*PeerEndpoint) ProtoMessage()               {}
func (*PeerEndpoint) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func (m *PeerEndpoint) GetId() *PeerID {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *PeerEndpoint) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*PeerID)(nil), "protos.PeerID")
	proto.RegisterType((*PeerEndpoint)(nil), "protos.PeerEndpoint")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Endorser service

type EndorserClient interface {
	ProcessProposal(ctx context.Context, in *SignedProposal, opts ...grpc.CallOption) (*ProposalResponse, error)
}

type endorserClient struct {
	cc *grpc.ClientConn
}

func NewEndorserClient(cc *grpc.ClientConn) EndorserClient {
	return &endorserClient{cc}
}

func (c *endorserClient) ProcessProposal(ctx context.Context, in *SignedProposal, opts ...grpc.CallOption) (*ProposalResponse, error) {
	out := new(ProposalResponse)
	err := grpc.Invoke(ctx, "/protos.Endorser/ProcessProposal", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Endorser service

type EndorserServer interface {
	ProcessProposal(context.Context, *SignedProposal) (*ProposalResponse, error)
}

func RegisterEndorserServer(s *grpc.Server, srv EndorserServer) {
	s.RegisterService(&_Endorser_serviceDesc, srv)
}

func _Endorser_ProcessProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignedProposal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndorserServer).ProcessProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Endorser/ProcessProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndorserServer).ProcessProposal(ctx, req.(*SignedProposal))
	}
	return interceptor(ctx, in, info, handler)
}

var _Endorser_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Endorser",
	HandlerType: (*EndorserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessProposal",
			Handler:    _Endorser_ProcessProposal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "peer/peer.proto",
}

func init() { proto.RegisterFile("peer/peer.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcf, 0x4b, 0xfc, 0x30,
	0x10, 0xc5, 0xbf, 0x5b, 0xbe, 0xac, 0x1a, 0xc5, 0x85, 0x08, 0x52, 0xca, 0x22, 0xd2, 0x93, 0x78,
	0x68, 0x71, 0xfd, 0x0f, 0xc4, 0x82, 0xde, 0x6a, 0xbd, 0x79, 0x59, 0xd2, 0x66, 0xec, 0x06, 0x77,
	0x67, 0xca, 0x4c, 0xf7, 0xff, 0x97, 0x26, 0x8d, 0x20, 0x5e, 0xf2, 0xe3, 0xbd, 0x7c, 0x5e, 0x86,
	0xa7, 0x56, 0x03, 0x00, 0x97, 0xd3, 0x52, 0x0c, 0x4c, 0x23, 0xe9, 0xa5, 0xdf, 0x24, 0xbb, 0x0a,
	0x06, 0xd3, 0x40, 0x62, 0xf6, 0xc1, 0xcc, 0xd6, 0xbf, 0xc4, 0x2d, 0x83, 0x0c, 0x84, 0x02, 0xc1,
	0xcd, 0xd7, 0x6a, 0x59, 0x03, 0xf0, 0xeb, 0xb3, 0xd6, 0xea, 0x3f, 0x9a, 0x03, 0xa4, 0x8b, 0xdb,
	0xc5, 0xdd, 0x59, 0xe3, 0xcf, 0xf9, 0x8b, 0xba, 0x98, 0xdc, 0x0a, 0xed, 0x40, 0x0e, 0x47, 0x7d,
	0xa3, 0x12, 0x67, 0xfd, 0x8b, 0xf3, 0xcd, 0x65, 0x48, 0x90, 0x22, 0xf0, 0x4d, 0xe2, 0xac, 0x4e,
	0xd5, 0x89, 0xb1, 0x96, 0x41, 0x24, 0x4d, 0x7c, 0x4c, 0xbc, 0x6e, 0xde, 0xd4, 0x69, 0x85, 0x96,
	0x58, 0x80, 0x75, 0xa5, 0x56, 0x35, 0x53, 0x07, 0x22, 0xf5, 0x3c, 0x95, 0xbe, 0x8e, 0x61, 0xef,
	0xae, 0x47, 0xb0, 0x51, 0xcf, 0xd2, 0x9f, 0x4f, 0x66, 0xa5, 0x99, 0xc7, 0xcf, 0xff, 0x3d, 0x6d,
	0xd5, 0x3d, 0x71, 0x5f, 0x38, 0xfc, 0xda, 0x9b, 0x56, 0x3e, 0xe9, 0x88, 0xd6, 0x8c, 0x8e, 0x70,
	0x52, 0xba, 0x9d, 0x71, 0x18, 0xd9, 0xa9, 0x80, 0x8f, 0x87, 0xde, 0x8d, 0xbb, 0x63, 0x5b, 0x74,
	0x74, 0x28, 0xff, 0x20, 0x65, 0x44, 0xca, 0x80, 0xf8, 0x72, 0xdb, 0x50, 0xeb, 0xe3, 0x77, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xc6, 0xec, 0xc5, 0x4d, 0x70, 0x01, 0x00, 0x00,
}
