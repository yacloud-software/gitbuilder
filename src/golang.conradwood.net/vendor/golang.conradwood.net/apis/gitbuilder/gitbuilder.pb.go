// Code generated by protoc-gen-go.
// source: golang.conradwood.net/apis/gitbuilder/gitbuilder.proto
// DO NOT EDIT!

/*
Package gitbuilder is a generated protocol buffer package.

It is generated from these files:
	golang.conradwood.net/apis/gitbuilder/gitbuilder.proto

It has these top-level messages:
	PingResponse
	BuildRequest
	BuildResponse
*/
package gitbuilder

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "golang.conradwood.net/apis/common"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// comment: message pingresponse
type PingResponse struct {
	// comment: field pingresponse.response
	Response string `protobuf:"bytes,1,opt,name=Response" json:"Response,omitempty"`
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PingResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

type BuildRequest struct {
	GitURL    string   `protobuf:"bytes,1,opt,name=GitURL" json:"GitURL,omitempty"`
	FetchURLS []string `protobuf:"bytes,2,rep,name=FetchURLS" json:"FetchURLS,omitempty"`
	CommitID  string   `protobuf:"bytes,3,opt,name=CommitID" json:"CommitID,omitempty"`
}

func (m *BuildRequest) Reset()                    { *m = BuildRequest{} }
func (m *BuildRequest) String() string            { return proto.CompactTextString(m) }
func (*BuildRequest) ProtoMessage()               {}
func (*BuildRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BuildRequest) GetGitURL() string {
	if m != nil {
		return m.GitURL
	}
	return ""
}

func (m *BuildRequest) GetFetchURLS() []string {
	if m != nil {
		return m.FetchURLS
	}
	return nil
}

func (m *BuildRequest) GetCommitID() string {
	if m != nil {
		return m.CommitID
	}
	return ""
}

type BuildResponse struct {
	Stdout        []byte `protobuf:"bytes,1,opt,name=Stdout,proto3" json:"Stdout,omitempty"`
	Complete      bool   `protobuf:"varint,2,opt,name=Complete" json:"Complete,omitempty"`
	ResultMessage string `protobuf:"bytes,3,opt,name=ResultMessage" json:"ResultMessage,omitempty"`
	Success       bool   `protobuf:"varint,4,opt,name=Success" json:"Success,omitempty"`
}

func (m *BuildResponse) Reset()                    { *m = BuildResponse{} }
func (m *BuildResponse) String() string            { return proto.CompactTextString(m) }
func (*BuildResponse) ProtoMessage()               {}
func (*BuildResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BuildResponse) GetStdout() []byte {
	if m != nil {
		return m.Stdout
	}
	return nil
}

func (m *BuildResponse) GetComplete() bool {
	if m != nil {
		return m.Complete
	}
	return false
}

func (m *BuildResponse) GetResultMessage() string {
	if m != nil {
		return m.ResultMessage
	}
	return ""
}

func (m *BuildResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*PingResponse)(nil), "gitbuilder.PingResponse")
	proto.RegisterType((*BuildRequest)(nil), "gitbuilder.BuildRequest")
	proto.RegisterType((*BuildResponse)(nil), "gitbuilder.BuildResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GitBuilder service

type GitBuilderClient interface {
	// comment: rpc ping
	Ping(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*PingResponse, error)
	// build something. Note that this RPC might take several minutes to complete
	Build(ctx context.Context, in *BuildRequest, opts ...grpc.CallOption) (GitBuilder_BuildClient, error)
}

type gitBuilderClient struct {
	cc *grpc.ClientConn
}

func NewGitBuilderClient(cc *grpc.ClientConn) GitBuilderClient {
	return &gitBuilderClient{cc}
}

func (c *gitBuilderClient) Ping(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/gitbuilder.GitBuilder/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitBuilderClient) Build(ctx context.Context, in *BuildRequest, opts ...grpc.CallOption) (GitBuilder_BuildClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GitBuilder_serviceDesc.Streams[0], c.cc, "/gitbuilder.GitBuilder/Build", opts...)
	if err != nil {
		return nil, err
	}
	x := &gitBuilderBuildClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GitBuilder_BuildClient interface {
	Recv() (*BuildResponse, error)
	grpc.ClientStream
}

type gitBuilderBuildClient struct {
	grpc.ClientStream
}

func (x *gitBuilderBuildClient) Recv() (*BuildResponse, error) {
	m := new(BuildResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for GitBuilder service

type GitBuilderServer interface {
	// comment: rpc ping
	Ping(context.Context, *common.Void) (*PingResponse, error)
	// build something. Note that this RPC might take several minutes to complete
	Build(*BuildRequest, GitBuilder_BuildServer) error
}

func RegisterGitBuilderServer(s *grpc.Server, srv GitBuilderServer) {
	s.RegisterService(&_GitBuilder_serviceDesc, srv)
}

func _GitBuilder_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitBuilderServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitbuilder.GitBuilder/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitBuilderServer).Ping(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitBuilder_Build_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BuildRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GitBuilderServer).Build(m, &gitBuilderBuildServer{stream})
}

type GitBuilder_BuildServer interface {
	Send(*BuildResponse) error
	grpc.ServerStream
}

type gitBuilderBuildServer struct {
	grpc.ServerStream
}

func (x *gitBuilderBuildServer) Send(m *BuildResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _GitBuilder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitbuilder.GitBuilder",
	HandlerType: (*GitBuilderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _GitBuilder_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Build",
			Handler:       _GitBuilder_Build_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "golang.conradwood.net/apis/gitbuilder/gitbuilder.proto",
}

func init() {
	proto.RegisterFile("golang.conradwood.net/apis/gitbuilder/gitbuilder.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x51, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x25, 0x6d, 0xad, 0xed, 0x90, 0x5e, 0xf6, 0x50, 0xd6, 0x20, 0x12, 0x8a, 0x87, 0xe2, 0x61,
	0x2b, 0x0a, 0x1e, 0x3d, 0x54, 0xb1, 0x08, 0x15, 0x64, 0x4b, 0x3d, 0x9b, 0x26, 0x43, 0x5c, 0x48,
	0x76, 0x63, 0x77, 0x82, 0x17, 0x6f, 0xfe, 0xb8, 0x24, 0xd9, 0xb6, 0x29, 0x88, 0xa7, 0xcc, 0x9b,
	0xcc, 0x7b, 0x6f, 0x79, 0x0f, 0xee, 0x52, 0x93, 0x45, 0x3a, 0x15, 0xb1, 0xd1, 0xdb, 0x28, 0xf9,
	0x32, 0x26, 0x11, 0x1a, 0x69, 0x16, 0x15, 0xca, 0xce, 0x52, 0x45, 0x9b, 0x52, 0x65, 0x09, 0x6e,
	0x5b, 0xa3, 0x28, 0xb6, 0x86, 0x0c, 0x83, 0xc3, 0x26, 0x10, 0xff, 0x68, 0xc4, 0x26, 0xcf, 0x8d,
	0x76, 0x9f, 0x86, 0x3b, 0xb9, 0x02, 0xff, 0x55, 0xe9, 0x54, 0xa2, 0x2d, 0x8c, 0xb6, 0xc8, 0x02,
	0x18, 0xec, 0x66, 0xee, 0x85, 0xde, 0x74, 0x28, 0xf7, 0x78, 0xf2, 0x0e, 0xfe, 0xbc, 0xb2, 0x91,
	0xf8, 0x59, 0xa2, 0x25, 0x36, 0x86, 0xfe, 0x42, 0xd1, 0x5a, 0x2e, 0xdd, 0xa5, 0x43, 0xec, 0x1c,
	0x86, 0x4f, 0x48, 0xf1, 0xc7, 0x5a, 0x2e, 0x57, 0xbc, 0x13, 0x76, 0xa7, 0x43, 0x79, 0x58, 0x54,
	0x0e, 0x0f, 0x26, 0xcf, 0x15, 0x3d, 0x3f, 0xf2, 0x6e, 0xe3, 0xb0, 0xc3, 0x93, 0x1f, 0x0f, 0x46,
	0xce, 0xc2, 0xbd, 0x67, 0x0c, 0xfd, 0x15, 0x25, 0xa6, 0xa4, 0xda, 0xc3, 0x97, 0x0e, 0x39, 0x95,
	0x22, 0x43, 0x42, 0xde, 0x09, 0xbd, 0xe9, 0x40, 0xee, 0x31, 0xbb, 0x84, 0x91, 0x44, 0x5b, 0x66,
	0xf4, 0x82, 0xd6, 0x46, 0x29, 0x3a, 0x9b, 0xe3, 0x25, 0xe3, 0x70, 0xba, 0x2a, 0xe3, 0x18, 0xad,
	0xe5, 0xbd, 0x5a, 0x60, 0x07, 0x6f, 0xbe, 0x01, 0x16, 0x8a, 0xe6, 0x4d, 0xa2, 0x4c, 0x40, 0xaf,
	0x4a, 0x88, 0xf9, 0xc2, 0x05, 0xf7, 0x66, 0x54, 0x12, 0x70, 0xd1, 0xaa, 0xe1, 0x28, 0xc1, 0x7b,
	0x38, 0xa9, 0xa9, 0xec, 0xe8, 0xa4, 0x1d, 0x5c, 0x70, 0xf6, 0xc7, 0x9f, 0x86, 0x7d, 0xed, 0xcd,
	0x43, 0xb8, 0xd0, 0x48, 0xed, 0x02, 0xab, 0xf2, 0x5a, 0x8c, 0x4d, 0xbf, 0xae, 0xee, 0xf6, 0x37,
	0x00, 0x00, 0xff, 0xff, 0x21, 0x33, 0x99, 0x0c, 0x30, 0x02, 0x00, 0x00,
}
