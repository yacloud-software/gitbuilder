// Code generated by protoc-gen-go.
// source: protos/golang.conradwood.net/apis/gitbuilder/gitbuilder.proto
// DO NOT EDIT!

/*
Package gitbuilder is a generated protocol buffer package.

It is generated from these files:
	protos/golang.conradwood.net/apis/gitbuilder/gitbuilder.proto

It has these top-level messages:
	FetchURL
	BuildRequest
	BuildResponse
	LocalRepo
	LocalRepoList
	BuildScriptList
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

type FetchURL struct {
	URL     string `protobuf:"bytes,1,opt,name=URL" json:"URL,omitempty"`
	RefSpec string `protobuf:"bytes,2,opt,name=RefSpec" json:"RefSpec,omitempty"`
}

func (m *FetchURL) Reset()                    { *m = FetchURL{} }
func (m *FetchURL) String() string            { return proto.CompactTextString(m) }
func (*FetchURL) ProtoMessage()               {}
func (*FetchURL) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FetchURL) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

func (m *FetchURL) GetRefSpec() string {
	if m != nil {
		return m.RefSpec
	}
	return ""
}

type BuildRequest struct {
	GitURL              string      `protobuf:"bytes,1,opt,name=GitURL" json:"GitURL,omitempty"`
	FetchURLs           []*FetchURL `protobuf:"bytes,2,rep,name=FetchURLs" json:"FetchURLs,omitempty"`
	CommitID            string      `protobuf:"bytes,3,opt,name=CommitID" json:"CommitID,omitempty"`
	BuildNumber         uint64      `protobuf:"varint,4,opt,name=BuildNumber" json:"BuildNumber,omitempty"`
	RepositoryID        uint64      `protobuf:"varint,5,opt,name=RepositoryID" json:"RepositoryID,omitempty"`
	RepoName            string      `protobuf:"bytes,6,opt,name=RepoName" json:"RepoName,omitempty"`
	ArtefactName        string      `protobuf:"bytes,7,opt,name=ArtefactName" json:"ArtefactName,omitempty"`
	ExcludeBuildScripts []string    `protobuf:"bytes,8,rep,name=ExcludeBuildScripts" json:"ExcludeBuildScripts,omitempty"`
	IncludeBuildScripts []string    `protobuf:"bytes,9,rep,name=IncludeBuildScripts" json:"IncludeBuildScripts,omitempty"`
	RequiresDeepClone   bool        `protobuf:"varint,10,opt,name=RequiresDeepClone" json:"RequiresDeepClone,omitempty"`
	ArtefactID          uint64      `protobuf:"varint,11,opt,name=ArtefactID" json:"ArtefactID,omitempty"`
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

func (m *BuildRequest) GetFetchURLs() []*FetchURL {
	if m != nil {
		return m.FetchURLs
	}
	return nil
}

func (m *BuildRequest) GetCommitID() string {
	if m != nil {
		return m.CommitID
	}
	return ""
}

func (m *BuildRequest) GetBuildNumber() uint64 {
	if m != nil {
		return m.BuildNumber
	}
	return 0
}

func (m *BuildRequest) GetRepositoryID() uint64 {
	if m != nil {
		return m.RepositoryID
	}
	return 0
}

func (m *BuildRequest) GetRepoName() string {
	if m != nil {
		return m.RepoName
	}
	return ""
}

func (m *BuildRequest) GetArtefactName() string {
	if m != nil {
		return m.ArtefactName
	}
	return ""
}

func (m *BuildRequest) GetExcludeBuildScripts() []string {
	if m != nil {
		return m.ExcludeBuildScripts
	}
	return nil
}

func (m *BuildRequest) GetIncludeBuildScripts() []string {
	if m != nil {
		return m.IncludeBuildScripts
	}
	return nil
}

func (m *BuildRequest) GetRequiresDeepClone() bool {
	if m != nil {
		return m.RequiresDeepClone
	}
	return false
}

func (m *BuildRequest) GetArtefactID() uint64 {
	if m != nil {
		return m.ArtefactID
	}
	return 0
}

type BuildResponse struct {
	Stdout        []byte `protobuf:"bytes,1,opt,name=Stdout,proto3" json:"Stdout,omitempty"`
	Complete      bool   `protobuf:"varint,2,opt,name=Complete" json:"Complete,omitempty"`
	ResultMessage string `protobuf:"bytes,3,opt,name=ResultMessage" json:"ResultMessage,omitempty"`
	Success       bool   `protobuf:"varint,4,opt,name=Success" json:"Success,omitempty"`
	LogMessage    string `protobuf:"bytes,5,opt,name=LogMessage" json:"LogMessage,omitempty"`
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

func (m *BuildResponse) GetLogMessage() string {
	if m != nil {
		return m.LogMessage
	}
	return ""
}

type LocalRepo struct {
	URL       string      `protobuf:"bytes,1,opt,name=URL" json:"URL,omitempty"`
	FetchURLs []*FetchURL `protobuf:"bytes,2,rep,name=FetchURLs" json:"FetchURLs,omitempty"`
	InUse     bool        `protobuf:"varint,3,opt,name=InUse" json:"InUse,omitempty"`
	WorkDir   string      `protobuf:"bytes,4,opt,name=WorkDir" json:"WorkDir,omitempty"`
	Created   uint32      `protobuf:"varint,5,opt,name=Created" json:"Created,omitempty"`
	Released  uint32      `protobuf:"varint,6,opt,name=Released" json:"Released,omitempty"`
}

func (m *LocalRepo) Reset()                    { *m = LocalRepo{} }
func (m *LocalRepo) String() string            { return proto.CompactTextString(m) }
func (*LocalRepo) ProtoMessage()               {}
func (*LocalRepo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LocalRepo) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

func (m *LocalRepo) GetFetchURLs() []*FetchURL {
	if m != nil {
		return m.FetchURLs
	}
	return nil
}

func (m *LocalRepo) GetInUse() bool {
	if m != nil {
		return m.InUse
	}
	return false
}

func (m *LocalRepo) GetWorkDir() string {
	if m != nil {
		return m.WorkDir
	}
	return ""
}

func (m *LocalRepo) GetCreated() uint32 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *LocalRepo) GetReleased() uint32 {
	if m != nil {
		return m.Released
	}
	return 0
}

type LocalRepoList struct {
	Repos []*LocalRepo `protobuf:"bytes,1,rep,name=Repos" json:"Repos,omitempty"`
}

func (m *LocalRepoList) Reset()                    { *m = LocalRepoList{} }
func (m *LocalRepoList) String() string            { return proto.CompactTextString(m) }
func (*LocalRepoList) ProtoMessage()               {}
func (*LocalRepoList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *LocalRepoList) GetRepos() []*LocalRepo {
	if m != nil {
		return m.Repos
	}
	return nil
}

type BuildScriptList struct {
	Names []string `protobuf:"bytes,1,rep,name=Names" json:"Names,omitempty"`
}

func (m *BuildScriptList) Reset()                    { *m = BuildScriptList{} }
func (m *BuildScriptList) String() string            { return proto.CompactTextString(m) }
func (*BuildScriptList) ProtoMessage()               {}
func (*BuildScriptList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *BuildScriptList) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func init() {
	proto.RegisterType((*FetchURL)(nil), "gitbuilder.FetchURL")
	proto.RegisterType((*BuildRequest)(nil), "gitbuilder.BuildRequest")
	proto.RegisterType((*BuildResponse)(nil), "gitbuilder.BuildResponse")
	proto.RegisterType((*LocalRepo)(nil), "gitbuilder.LocalRepo")
	proto.RegisterType((*LocalRepoList)(nil), "gitbuilder.LocalRepoList")
	proto.RegisterType((*BuildScriptList)(nil), "gitbuilder.BuildScriptList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GitBuilder service

type GitBuilderClient interface {
	// build something. Note that this RPC might take several minutes to complete
	Build(ctx context.Context, in *BuildRequest, opts ...grpc.CallOption) (GitBuilder_BuildClient, error)
	// get information about the repos on disk
	GetLocalRepos(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*LocalRepoList, error)
	// get "known" build scripts
	GetBuildScripts(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*BuildScriptList, error)
}

type gitBuilderClient struct {
	cc *grpc.ClientConn
}

func NewGitBuilderClient(cc *grpc.ClientConn) GitBuilderClient {
	return &gitBuilderClient{cc}
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

func (c *gitBuilderClient) GetLocalRepos(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*LocalRepoList, error) {
	out := new(LocalRepoList)
	err := grpc.Invoke(ctx, "/gitbuilder.GitBuilder/GetLocalRepos", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitBuilderClient) GetBuildScripts(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*BuildScriptList, error) {
	out := new(BuildScriptList)
	err := grpc.Invoke(ctx, "/gitbuilder.GitBuilder/GetBuildScripts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GitBuilder service

type GitBuilderServer interface {
	// build something. Note that this RPC might take several minutes to complete
	Build(*BuildRequest, GitBuilder_BuildServer) error
	// get information about the repos on disk
	GetLocalRepos(context.Context, *common.Void) (*LocalRepoList, error)
	// get "known" build scripts
	GetBuildScripts(context.Context, *common.Void) (*BuildScriptList, error)
}

func RegisterGitBuilderServer(s *grpc.Server, srv GitBuilderServer) {
	s.RegisterService(&_GitBuilder_serviceDesc, srv)
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

func _GitBuilder_GetLocalRepos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitBuilderServer).GetLocalRepos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitbuilder.GitBuilder/GetLocalRepos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitBuilderServer).GetLocalRepos(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitBuilder_GetBuildScripts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitBuilderServer).GetBuildScripts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitbuilder.GitBuilder/GetBuildScripts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitBuilderServer).GetBuildScripts(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _GitBuilder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitbuilder.GitBuilder",
	HandlerType: (*GitBuilderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLocalRepos",
			Handler:    _GitBuilder_GetLocalRepos_Handler,
		},
		{
			MethodName: "GetBuildScripts",
			Handler:    _GitBuilder_GetBuildScripts_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Build",
			Handler:       _GitBuilder_Build_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/golang.conradwood.net/apis/gitbuilder/gitbuilder.proto",
}

func init() {
	proto.RegisterFile("protos/golang.conradwood.net/apis/gitbuilder/gitbuilder.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 609 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0x51, 0x6b, 0xd4, 0x40,
	0x10, 0x26, 0xbd, 0x5e, 0x9b, 0x4c, 0xef, 0xa8, 0xae, 0x55, 0xd6, 0x0a, 0xe5, 0x08, 0x8a, 0x07,
	0x4a, 0x5a, 0x2a, 0x88, 0x0f, 0x55, 0xb0, 0x3d, 0x3d, 0x0e, 0xce, 0x3e, 0xec, 0x51, 0x05, 0xdf,
	0xd2, 0x64, 0x7a, 0x2e, 0x26, 0xd9, 0x98, 0xdd, 0xa0, 0xfe, 0x1e, 0xdf, 0x7d, 0xf5, 0x0f, 0xf8,
	0xc3, 0x64, 0x77, 0x93, 0x5c, 0xc2, 0x1d, 0x22, 0x3e, 0xdd, 0x7e, 0x33, 0xfb, 0xed, 0x7c, 0x93,
	0xf9, 0x6e, 0xe0, 0x65, 0x5e, 0x08, 0x25, 0xe4, 0xf1, 0x52, 0x24, 0x61, 0xb6, 0x0c, 0x22, 0x91,
	0x15, 0x61, 0xfc, 0x55, 0x88, 0x38, 0xc8, 0x50, 0x1d, 0x87, 0x39, 0x97, 0xc7, 0x4b, 0xae, 0xae,
	0x4b, 0x9e, 0xc4, 0x58, 0xb4, 0x8e, 0x81, 0xe1, 0x11, 0x58, 0x45, 0x0e, 0x83, 0xbf, 0xbc, 0x11,
	0x89, 0x34, 0x15, 0x59, 0xf5, 0x63, 0xb9, 0xfe, 0x73, 0x70, 0xdf, 0xa2, 0x8a, 0x3e, 0x5d, 0xb1,
	0x39, 0xb9, 0x05, 0xbd, 0x2b, 0x36, 0xa7, 0xce, 0xc8, 0x19, 0x7b, 0x4c, 0x1f, 0x09, 0x85, 0x5d,
	0x86, 0x37, 0x8b, 0x1c, 0x23, 0xba, 0x65, 0xa2, 0x35, 0xf4, 0x7f, 0xf6, 0x60, 0x70, 0xae, 0x6b,
	0x32, 0xfc, 0x52, 0xa2, 0x54, 0xe4, 0x1e, 0xec, 0x4c, 0xb9, 0x5a, 0xf1, 0x2b, 0x44, 0x4e, 0xc1,
	0xab, 0x0b, 0x48, 0xba, 0x35, 0xea, 0x8d, 0xf7, 0x4e, 0x0f, 0x82, 0x56, 0x0b, 0x75, 0x92, 0xad,
	0xae, 0x91, 0x43, 0x70, 0x2f, 0x44, 0x9a, 0x72, 0x35, 0x9b, 0xd0, 0x9e, 0x79, 0xad, 0xc1, 0x64,
	0x04, 0x7b, 0xa6, 0xee, 0x65, 0x99, 0x5e, 0x63, 0x41, 0xb7, 0x47, 0xce, 0x78, 0x9b, 0xb5, 0x43,
	0xc4, 0x87, 0x01, 0xc3, 0x5c, 0x48, 0xae, 0x44, 0xf1, 0x7d, 0x36, 0xa1, 0x7d, 0x73, 0xa5, 0x13,
	0xd3, 0x15, 0x34, 0xbe, 0x0c, 0x53, 0xa4, 0x3b, 0xb6, 0x42, 0x8d, 0x35, 0xff, 0x75, 0xa1, 0xf0,
	0x26, 0x8c, 0x94, 0xc9, 0xef, 0x9a, 0x7c, 0x27, 0x46, 0x4e, 0xe0, 0xce, 0x9b, 0x6f, 0x51, 0x52,
	0xc6, 0x68, 0x2a, 0x2f, 0xa2, 0x82, 0xe7, 0x4a, 0x52, 0x77, 0xd4, 0x1b, 0x7b, 0x6c, 0x53, 0x4a,
	0x33, 0x66, 0xd9, 0x3a, 0xc3, 0xb3, 0x8c, 0x0d, 0x29, 0xf2, 0x14, 0x6e, 0xeb, 0x8f, 0xcb, 0x0b,
	0x94, 0x13, 0xc4, 0xfc, 0x22, 0x11, 0x19, 0x52, 0x18, 0x39, 0x63, 0x97, 0xad, 0x27, 0xc8, 0x11,
	0x40, 0xad, 0x70, 0x36, 0xa1, 0x7b, 0xa6, 0xe7, 0x56, 0xc4, 0xff, 0xe1, 0xc0, 0xb0, 0x1a, 0x98,
	0xcc, 0x45, 0x26, 0x51, 0x4f, 0x6c, 0xa1, 0x62, 0x51, 0x2a, 0x33, 0xb1, 0x01, 0xab, 0x50, 0xf5,
	0xf5, 0xf3, 0x04, 0x15, 0x9a, 0xa9, 0xbb, 0xac, 0xc1, 0xe4, 0x21, 0x0c, 0x19, 0xca, 0x32, 0x51,
	0xef, 0x50, 0xca, 0x70, 0x89, 0xd5, 0x78, 0xba, 0x41, 0x6d, 0x9b, 0x45, 0x19, 0x45, 0x28, 0xa5,
	0x99, 0x8f, 0xcb, 0x6a, 0xa8, 0x55, 0xce, 0xc5, 0xb2, 0x26, 0xf7, 0x0d, 0xb9, 0x15, 0xf1, 0x7f,
	0x39, 0xe0, 0xcd, 0x45, 0x14, 0x26, 0x7a, 0x1a, 0x1b, 0x0c, 0xf9, 0x3f, 0x6e, 0x3a, 0x80, 0xfe,
	0x2c, 0xbb, 0x92, 0x56, 0xab, 0xcb, 0x2c, 0xd0, 0x1a, 0x3f, 0x88, 0xe2, 0xf3, 0x84, 0x5b, 0x0f,
	0x79, 0xac, 0x86, 0x3a, 0x73, 0x51, 0x60, 0xa8, 0x30, 0x36, 0x02, 0x87, 0xac, 0x86, 0xd6, 0x35,
	0x09, 0x86, 0x12, 0x63, 0xe3, 0x9a, 0x21, 0x6b, 0xb0, 0x7f, 0x06, 0xc3, 0x46, 0xf8, 0x9c, 0x4b,
	0x45, 0x9e, 0x40, 0xdf, 0x58, 0x8e, 0x3a, 0x46, 0xe6, 0xdd, 0xb6, 0xcc, 0xe6, 0x26, 0xb3, 0x77,
	0xfc, 0xc7, 0xb0, 0xdf, 0x9a, 0xbd, 0xe1, 0x1f, 0x40, 0x5f, 0x5b, 0xcd, 0xf2, 0x3d, 0x66, 0xc1,
	0xe9, 0x6f, 0x07, 0x60, 0xca, 0xd5, 0xb9, 0x7d, 0x88, 0xbc, 0x82, 0xbe, 0x39, 0x12, 0xda, 0x7e,
	0xbe, 0xfd, 0xc7, 0x3c, 0xbc, 0xbf, 0x21, 0x63, 0x1d, 0x70, 0xe2, 0x90, 0x17, 0x30, 0x9c, 0xa2,
	0x6a, 0xe4, 0x48, 0x32, 0x08, 0xaa, 0xf5, 0xf0, 0x5e, 0xf0, 0xb8, 0xcb, 0xed, 0xb6, 0x77, 0x06,
	0xfb, 0x53, 0x54, 0x1d, 0xc3, 0x76, 0xb9, 0x0f, 0xd6, 0xea, 0xae, 0x9a, 0x3b, 0x9f, 0xc2, 0x51,
	0x86, 0xaa, 0xbd, 0xa5, 0xf4, 0x86, 0x6a, 0x31, 0x3e, 0x3e, 0xfa, 0xa7, 0x65, 0x78, 0xbd, 0x63,
	0xd6, 0xd8, 0xb3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x37, 0xa5, 0x60, 0xa1, 0x43, 0x05, 0x00,
	0x00,
}




