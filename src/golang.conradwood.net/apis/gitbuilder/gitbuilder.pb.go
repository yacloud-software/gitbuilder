// Code generated by protoc-gen-go.
// source: golang.conradwood.net/apis/gitbuilder/gitbuilder.proto
// DO NOT EDIT!

/*
Package gitbuilder is a generated protocol buffer package.

It is generated from these files:
	golang.conradwood.net/apis/gitbuilder/gitbuilder.proto

It has these top-level messages:
	FetchURL
	BuildRequest
	BuildResponse
	LocalRepo
	LocalRepoList
	BuildScriptList
	FileTransferPart
	BuildLocalRequest
	BuildLocalFiles
	BuildRuleDef_Go
	BuildRuleDef_C
	BuildRuleDef
	BuildRules
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

type ReturnSet int32

const (
	ReturnSet_RETURN_ALL_CHANGED    ReturnSet = 0
	ReturnSet_RETURN_DIST           ReturnSet = 1
	ReturnSet_RETURN_BUNDLE         ReturnSet = 2
	ReturnSet_RETURN_BUNDLE_CHANGED ReturnSet = 3
)

var ReturnSet_name = map[int32]string{
	0: "RETURN_ALL_CHANGED",
	1: "RETURN_DIST",
	2: "RETURN_BUNDLE",
	3: "RETURN_BUNDLE_CHANGED",
}
var ReturnSet_value = map[string]int32{
	"RETURN_ALL_CHANGED":    0,
	"RETURN_DIST":           1,
	"RETURN_BUNDLE":         2,
	"RETURN_BUNDLE_CHANGED": 3,
}

func (x ReturnSet) String() string {
	return proto.EnumName(ReturnSet_name, int32(x))
}
func (ReturnSet) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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
	Stdout        []byte            `protobuf:"bytes,1,opt,name=Stdout,proto3" json:"Stdout,omitempty"`
	Complete      bool              `protobuf:"varint,2,opt,name=Complete" json:"Complete,omitempty"`
	ResultMessage string            `protobuf:"bytes,3,opt,name=ResultMessage" json:"ResultMessage,omitempty"`
	Success       bool              `protobuf:"varint,4,opt,name=Success" json:"Success,omitempty"`
	LogMessage    string            `protobuf:"bytes,5,opt,name=LogMessage" json:"LogMessage,omitempty"`
	FileTransfer  *FileTransferPart `protobuf:"bytes,6,opt,name=FileTransfer" json:"FileTransfer,omitempty"`
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

func (m *BuildResponse) GetFileTransfer() *FileTransferPart {
	if m != nil {
		return m.FileTransfer
	}
	return nil
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

type FileTransferPart struct {
	Filename    string `protobuf:"bytes,1,opt,name=Filename" json:"Filename,omitempty"`
	Data        []byte `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	Permissions uint32 `protobuf:"varint,3,opt,name=Permissions" json:"Permissions,omitempty"`
}

func (m *FileTransferPart) Reset()                    { *m = FileTransferPart{} }
func (m *FileTransferPart) String() string            { return proto.CompactTextString(m) }
func (*FileTransferPart) ProtoMessage()               {}
func (*FileTransferPart) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *FileTransferPart) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *FileTransferPart) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *FileTransferPart) GetPermissions() uint32 {
	if m != nil {
		return m.Permissions
	}
	return 0
}

type BuildLocalRequest struct {
	RepositoryID        uint64    `protobuf:"varint,1,opt,name=RepositoryID" json:"RepositoryID,omitempty"`
	RepoName            string    `protobuf:"bytes,2,opt,name=RepoName" json:"RepoName,omitempty"`
	ExcludeBuildScripts []string  `protobuf:"bytes,3,rep,name=ExcludeBuildScripts" json:"ExcludeBuildScripts,omitempty"`
	IncludeBuildScripts []string  `protobuf:"bytes,4,rep,name=IncludeBuildScripts" json:"IncludeBuildScripts,omitempty"`
	ArtefactID          uint64    `protobuf:"varint,5,opt,name=ArtefactID" json:"ArtefactID,omitempty"`
	ArtefactName        string    `protobuf:"bytes,6,opt,name=ArtefactName" json:"ArtefactName,omitempty"`
	BuildNumber         uint64    `protobuf:"varint,7,opt,name=BuildNumber" json:"BuildNumber,omitempty"`
	Return              ReturnSet `protobuf:"varint,8,opt,name=Return,enum=gitbuilder.ReturnSet" json:"Return,omitempty"`
}

func (m *BuildLocalRequest) Reset()                    { *m = BuildLocalRequest{} }
func (m *BuildLocalRequest) String() string            { return proto.CompactTextString(m) }
func (*BuildLocalRequest) ProtoMessage()               {}
func (*BuildLocalRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *BuildLocalRequest) GetRepositoryID() uint64 {
	if m != nil {
		return m.RepositoryID
	}
	return 0
}

func (m *BuildLocalRequest) GetRepoName() string {
	if m != nil {
		return m.RepoName
	}
	return ""
}

func (m *BuildLocalRequest) GetExcludeBuildScripts() []string {
	if m != nil {
		return m.ExcludeBuildScripts
	}
	return nil
}

func (m *BuildLocalRequest) GetIncludeBuildScripts() []string {
	if m != nil {
		return m.IncludeBuildScripts
	}
	return nil
}

func (m *BuildLocalRequest) GetArtefactID() uint64 {
	if m != nil {
		return m.ArtefactID
	}
	return 0
}

func (m *BuildLocalRequest) GetArtefactName() string {
	if m != nil {
		return m.ArtefactName
	}
	return ""
}

func (m *BuildLocalRequest) GetBuildNumber() uint64 {
	if m != nil {
		return m.BuildNumber
	}
	return 0
}

func (m *BuildLocalRequest) GetReturn() ReturnSet {
	if m != nil {
		return m.Return
	}
	return ReturnSet_RETURN_ALL_CHANGED
}

type BuildLocalFiles struct {
	FileTransfer *FileTransferPart  `protobuf:"bytes,1,opt,name=FileTransfer" json:"FileTransfer,omitempty"`
	Request      *BuildLocalRequest `protobuf:"bytes,2,opt,name=Request" json:"Request,omitempty"`
}

func (m *BuildLocalFiles) Reset()                    { *m = BuildLocalFiles{} }
func (m *BuildLocalFiles) String() string            { return proto.CompactTextString(m) }
func (*BuildLocalFiles) ProtoMessage()               {}
func (*BuildLocalFiles) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *BuildLocalFiles) GetFileTransfer() *FileTransferPart {
	if m != nil {
		return m.FileTransfer
	}
	return nil
}

func (m *BuildLocalFiles) GetRequest() *BuildLocalRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

type BuildRuleDef_Go struct {
	CGOEnabled  bool     `protobuf:"varint,1,opt,name=CGOEnabled" json:"CGOEnabled,omitempty"`
	ExcludeDirs []string `protobuf:"bytes,2,rep,name=ExcludeDirs" json:"ExcludeDirs,omitempty"`
}

func (m *BuildRuleDef_Go) Reset()                    { *m = BuildRuleDef_Go{} }
func (m *BuildRuleDef_Go) String() string            { return proto.CompactTextString(m) }
func (*BuildRuleDef_Go) ProtoMessage()               {}
func (*BuildRuleDef_Go) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *BuildRuleDef_Go) GetCGOEnabled() bool {
	if m != nil {
		return m.CGOEnabled
	}
	return false
}

func (m *BuildRuleDef_Go) GetExcludeDirs() []string {
	if m != nil {
		return m.ExcludeDirs
	}
	return nil
}

type BuildRuleDef_C struct {
	CompilerVersion string `protobuf:"bytes,2,opt,name=CompilerVersion" json:"CompilerVersion,omitempty"`
	OS              string `protobuf:"bytes,3,opt,name=OS" json:"OS,omitempty"`
	CPU             string `protobuf:"bytes,4,opt,name=CPU" json:"CPU,omitempty"`
}

func (m *BuildRuleDef_C) Reset()                    { *m = BuildRuleDef_C{} }
func (m *BuildRuleDef_C) String() string            { return proto.CompactTextString(m) }
func (*BuildRuleDef_C) ProtoMessage()               {}
func (*BuildRuleDef_C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *BuildRuleDef_C) GetCompilerVersion() string {
	if m != nil {
		return m.CompilerVersion
	}
	return ""
}

func (m *BuildRuleDef_C) GetOS() string {
	if m != nil {
		return m.OS
	}
	return ""
}

func (m *BuildRuleDef_C) GetCPU() string {
	if m != nil {
		return m.CPU
	}
	return ""
}

type BuildRuleDef struct {
	BuildType string           `protobuf:"bytes,1,opt,name=BuildType" json:"BuildType,omitempty"`
	BuildOS   string           `protobuf:"bytes,2,opt,name=BuildOS" json:"BuildOS,omitempty"`
	Go        *BuildRuleDef_Go `protobuf:"bytes,3,opt,name=Go" json:"Go,omitempty"`
	C         *BuildRuleDef_C  `protobuf:"bytes,4,opt,name=C" json:"C,omitempty"`
}

func (m *BuildRuleDef) Reset()                    { *m = BuildRuleDef{} }
func (m *BuildRuleDef) String() string            { return proto.CompactTextString(m) }
func (*BuildRuleDef) ProtoMessage()               {}
func (*BuildRuleDef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *BuildRuleDef) GetBuildType() string {
	if m != nil {
		return m.BuildType
	}
	return ""
}

func (m *BuildRuleDef) GetBuildOS() string {
	if m != nil {
		return m.BuildOS
	}
	return ""
}

func (m *BuildRuleDef) GetGo() *BuildRuleDef_Go {
	if m != nil {
		return m.Go
	}
	return nil
}

func (m *BuildRuleDef) GetC() *BuildRuleDef_C {
	if m != nil {
		return m.C
	}
	return nil
}

// yaml file defining build rules
type BuildRules struct {
	PreBuild   string          `protobuf:"bytes,1,opt,name=PreBuild" json:"PreBuild,omitempty"`
	PostCommit string          `protobuf:"bytes,2,opt,name=PostCommit" json:"PostCommit,omitempty"`
	Rules      []*BuildRuleDef `protobuf:"bytes,3,rep,name=Rules" json:"Rules,omitempty"`
}

func (m *BuildRules) Reset()                    { *m = BuildRules{} }
func (m *BuildRules) String() string            { return proto.CompactTextString(m) }
func (*BuildRules) ProtoMessage()               {}
func (*BuildRules) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *BuildRules) GetPreBuild() string {
	if m != nil {
		return m.PreBuild
	}
	return ""
}

func (m *BuildRules) GetPostCommit() string {
	if m != nil {
		return m.PostCommit
	}
	return ""
}

func (m *BuildRules) GetRules() []*BuildRuleDef {
	if m != nil {
		return m.Rules
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
	proto.RegisterType((*FileTransferPart)(nil), "gitbuilder.FileTransferPart")
	proto.RegisterType((*BuildLocalRequest)(nil), "gitbuilder.BuildLocalRequest")
	proto.RegisterType((*BuildLocalFiles)(nil), "gitbuilder.BuildLocalFiles")
	proto.RegisterType((*BuildRuleDef_Go)(nil), "gitbuilder.BuildRuleDef_Go")
	proto.RegisterType((*BuildRuleDef_C)(nil), "gitbuilder.BuildRuleDef_C")
	proto.RegisterType((*BuildRuleDef)(nil), "gitbuilder.BuildRuleDef")
	proto.RegisterType((*BuildRules)(nil), "gitbuilder.BuildRules")
	proto.RegisterEnum("gitbuilder.ReturnSet", ReturnSet_name, ReturnSet_value)
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
	// build something that exists only on the local file system. return the files that would otherwise be uploaded to buildrepo to the client instead
	BuildFromLocalFiles(ctx context.Context, opts ...grpc.CallOption) (GitBuilder_BuildFromLocalFilesClient, error)
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

func (c *gitBuilderClient) BuildFromLocalFiles(ctx context.Context, opts ...grpc.CallOption) (GitBuilder_BuildFromLocalFilesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GitBuilder_serviceDesc.Streams[1], c.cc, "/gitbuilder.GitBuilder/BuildFromLocalFiles", opts...)
	if err != nil {
		return nil, err
	}
	x := &gitBuilderBuildFromLocalFilesClient{stream}
	return x, nil
}

type GitBuilder_BuildFromLocalFilesClient interface {
	Send(*BuildLocalFiles) error
	Recv() (*BuildResponse, error)
	grpc.ClientStream
}

type gitBuilderBuildFromLocalFilesClient struct {
	grpc.ClientStream
}

func (x *gitBuilderBuildFromLocalFilesClient) Send(m *BuildLocalFiles) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gitBuilderBuildFromLocalFilesClient) Recv() (*BuildResponse, error) {
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
	// build something that exists only on the local file system. return the files that would otherwise be uploaded to buildrepo to the client instead
	BuildFromLocalFiles(GitBuilder_BuildFromLocalFilesServer) error
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

func _GitBuilder_BuildFromLocalFiles_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GitBuilderServer).BuildFromLocalFiles(&gitBuilderBuildFromLocalFilesServer{stream})
}

type GitBuilder_BuildFromLocalFilesServer interface {
	Send(*BuildResponse) error
	Recv() (*BuildLocalFiles, error)
	grpc.ServerStream
}

type gitBuilderBuildFromLocalFilesServer struct {
	grpc.ServerStream
}

func (x *gitBuilderBuildFromLocalFilesServer) Send(m *BuildResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gitBuilderBuildFromLocalFilesServer) Recv() (*BuildLocalFiles, error) {
	m := new(BuildLocalFiles)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
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
		{
			StreamName:    "BuildFromLocalFiles",
			Handler:       _GitBuilder_BuildFromLocalFiles_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "golang.conradwood.net/apis/gitbuilder/gitbuilder.proto",
}

func init() {
	proto.RegisterFile("golang.conradwood.net/apis/gitbuilder/gitbuilder.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 1053 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x56, 0xef, 0x6e, 0x1b, 0x45,
	0x10, 0xe7, 0xec, 0x38, 0xb5, 0xc7, 0x76, 0xfe, 0x6c, 0xd3, 0xea, 0x9a, 0x96, 0xca, 0x3a, 0x81,
	0xb0, 0x28, 0xb8, 0x91, 0x91, 0x0a, 0x1f, 0x2a, 0x44, 0xe2, 0x4b, 0x8c, 0x25, 0x93, 0x84, 0xbd,
	0xb8, 0x48, 0x08, 0x29, 0xba, 0xdc, 0x4d, 0xcc, 0x89, 0xf3, 0xad, 0xd9, 0x5d, 0x8b, 0xf6, 0x1d,
	0xf8, 0xce, 0x27, 0x5e, 0x81, 0xaf, 0xbc, 0x04, 0xaf, 0xc1, 0x7b, 0xa0, 0xdd, 0xbd, 0x3b, 0xdf,
	0xd9, 0x4e, 0xa1, 0x7c, 0xb2, 0xe7, 0x37, 0x3b, 0xbb, 0xb3, 0x33, 0xbf, 0xdf, 0xdc, 0xc2, 0x8b,
	0x29, 0x8b, 0xfd, 0x64, 0xda, 0x0b, 0x58, 0xc2, 0xfd, 0xf0, 0x17, 0xc6, 0xc2, 0x5e, 0x82, 0xf2,
	0xb9, 0x3f, 0x8f, 0xc4, 0xf3, 0x69, 0x24, 0x6f, 0x16, 0x51, 0x1c, 0x22, 0x2f, 0xfc, 0xed, 0xcd,
	0x39, 0x93, 0x8c, 0xc0, 0x12, 0x39, 0xec, 0xbd, 0x65, 0x8f, 0x80, 0xcd, 0x66, 0x2c, 0x49, 0x7f,
	0x4c, 0xac, 0xf3, 0x02, 0xea, 0x67, 0x28, 0x83, 0x1f, 0x27, 0x74, 0x4c, 0xf6, 0xa0, 0x3a, 0xa1,
	0x63, 0xdb, 0xea, 0x58, 0xdd, 0x06, 0x55, 0x7f, 0x89, 0x0d, 0xf7, 0x28, 0xde, 0x7a, 0x73, 0x0c,
	0xec, 0x8a, 0x46, 0x33, 0xd3, 0xf9, 0xa3, 0x0a, 0xad, 0x13, 0x75, 0x26, 0xc5, 0x9f, 0x17, 0x28,
	0x24, 0x79, 0x08, 0xdb, 0xc3, 0x48, 0x2e, 0xe3, 0x53, 0x8b, 0xf4, 0xa1, 0x91, 0x1d, 0x20, 0xec,
	0x4a, 0xa7, 0xda, 0x6d, 0xf6, 0x0f, 0x7a, 0x85, 0x2b, 0x64, 0x4e, 0xba, 0x5c, 0x46, 0x0e, 0xa1,
	0x3e, 0x60, 0xb3, 0x59, 0x24, 0x47, 0xae, 0x5d, 0xd5, 0xbb, 0xe5, 0x36, 0xe9, 0x40, 0x53, 0x9f,
	0x7b, 0xbe, 0x98, 0xdd, 0x20, 0xb7, 0xb7, 0x3a, 0x56, 0x77, 0x8b, 0x16, 0x21, 0xe2, 0x40, 0x8b,
	0xe2, 0x9c, 0x89, 0x48, 0x32, 0xfe, 0x66, 0xe4, 0xda, 0x35, 0xbd, 0xa4, 0x84, 0xa9, 0x13, 0x94,
	0x7d, 0xee, 0xcf, 0xd0, 0xde, 0x36, 0x27, 0x64, 0xb6, 0x8a, 0x3f, 0xe6, 0x12, 0x6f, 0xfd, 0x40,
	0x6a, 0xff, 0x3d, 0xed, 0x2f, 0x61, 0xe4, 0x08, 0xee, 0x9f, 0xbe, 0x0e, 0xe2, 0x45, 0x88, 0xfa,
	0x64, 0x2f, 0xe0, 0xd1, 0x5c, 0x0a, 0xbb, 0xde, 0xa9, 0x76, 0x1b, 0x74, 0x93, 0x4b, 0x45, 0x8c,
	0x92, 0xf5, 0x88, 0x86, 0x89, 0xd8, 0xe0, 0x22, 0x9f, 0xc0, 0xbe, 0x2a, 0x6e, 0xc4, 0x51, 0xb8,
	0x88, 0xf3, 0x41, 0xcc, 0x12, 0xb4, 0xa1, 0x63, 0x75, 0xeb, 0x74, 0xdd, 0x41, 0x9e, 0x02, 0x64,
	0x19, 0x8e, 0x5c, 0xbb, 0xa9, 0xef, 0x5c, 0x40, 0x9c, 0xbf, 0x2d, 0x68, 0xa7, 0x0d, 0x13, 0x73,
	0x96, 0x08, 0x54, 0x1d, 0xf3, 0x64, 0xc8, 0x16, 0x52, 0x77, 0xac, 0x45, 0x53, 0x2b, 0xad, 0xfe,
	0x3c, 0x46, 0x89, 0xba, 0xeb, 0x75, 0x9a, 0xdb, 0xe4, 0x03, 0x68, 0x53, 0x14, 0x8b, 0x58, 0x7e,
	0x83, 0x42, 0xf8, 0x53, 0x4c, 0xdb, 0x53, 0x06, 0x15, 0x6d, 0xbc, 0x45, 0x10, 0xa0, 0x10, 0xba,
	0x3f, 0x75, 0x9a, 0x99, 0x2a, 0xcb, 0x31, 0x9b, 0x66, 0xc1, 0x35, 0x1d, 0x5c, 0x40, 0xc8, 0x57,
	0xd0, 0x3a, 0x8b, 0x62, 0xbc, 0xe2, 0x7e, 0x22, 0x6e, 0x91, 0xeb, 0xde, 0x34, 0xfb, 0x4f, 0x4a,
	0x84, 0x29, 0xf8, 0x2f, 0x7d, 0x2e, 0x69, 0x29, 0xc2, 0xf9, 0xd3, 0x82, 0xc6, 0x98, 0x05, 0x7e,
	0xac, 0xfa, 0xb9, 0x81, 0xd2, 0xff, 0x87, 0x8f, 0x07, 0x50, 0x1b, 0x25, 0x13, 0x61, 0x6e, 0x5b,
	0xa7, 0xc6, 0x50, 0xb7, 0xfc, 0x8e, 0xf1, 0x9f, 0xdc, 0xc8, 0xb0, 0xb0, 0x41, 0x33, 0x53, 0x79,
	0x06, 0x1c, 0x7d, 0x89, 0xa1, 0xbe, 0x62, 0x9b, 0x66, 0xa6, 0xe1, 0x5d, 0x8c, 0xbe, 0xc0, 0x50,
	0xdf, 0xad, 0x4d, 0x73, 0xdb, 0x79, 0x09, 0xed, 0x3c, 0xf1, 0x71, 0x24, 0x24, 0x79, 0x06, 0x35,
	0x4d, 0x5a, 0xdb, 0xd2, 0x69, 0x3e, 0x28, 0xa6, 0x99, 0xaf, 0xa4, 0x66, 0x8d, 0xf3, 0x11, 0xec,
	0x16, 0xd8, 0xa3, 0xe3, 0x0f, 0xa0, 0xa6, 0xc8, 0x6a, 0xe2, 0x1b, 0xd4, 0x18, 0x4e, 0x08, 0x7b,
	0xab, 0x25, 0x54, 0x69, 0x29, 0x2c, 0x51, 0x74, 0x37, 0xb5, 0xca, 0x6d, 0x42, 0x60, 0xcb, 0xf5,
	0xa5, 0xaf, 0xa9, 0xd0, 0xa2, 0xfa, 0xbf, 0x12, 0xe1, 0x25, 0xf2, 0x59, 0x24, 0x44, 0xc4, 0x12,
	0xa1, 0xcb, 0xd2, 0xa6, 0x45, 0xc8, 0xf9, 0xab, 0x02, 0xfb, 0x3a, 0x9f, 0x34, 0x51, 0x33, 0x24,
	0x56, 0xa5, 0x69, 0xfd, 0x8b, 0x34, 0x2b, 0x2b, 0xd2, 0xbc, 0x43, 0x76, 0xd5, 0x77, 0x96, 0xdd,
	0xd6, 0xdd, 0xb2, 0x2b, 0x0b, 0xa9, 0xb6, 0x2a, 0xa4, 0xb5, 0xf1, 0xb0, 0xbd, 0x61, 0x3c, 0xac,
	0x0c, 0xa9, 0x7b, 0xeb, 0x43, 0xea, 0x53, 0xd8, 0xa6, 0x28, 0x17, 0x3c, 0xb1, 0xeb, 0x1d, 0xab,
	0xbb, 0x53, 0x6e, 0xae, 0xf1, 0x78, 0x28, 0x69, 0xba, 0xc8, 0xf9, 0xd5, 0x4a, 0xdb, 0xab, 0xcb,
	0xa9, 0x7a, 0x23, 0xd6, 0xb4, 0x62, 0xbd, 0xab, 0x56, 0xc8, 0xe7, 0x6a, 0xbc, 0xeb, 0xce, 0xe8,
	0x4a, 0x37, 0xfb, 0xef, 0x17, 0x83, 0xd7, 0xda, 0x47, 0xb3, 0xd5, 0x8e, 0x97, 0x66, 0x43, 0x17,
	0x31, 0xba, 0x78, 0x7b, 0x3d, 0x64, 0xaa, 0x6c, 0x83, 0xe1, 0xc5, 0x69, 0xe2, 0xdf, 0xc4, 0x18,
	0xea, 0x5c, 0xea, 0xb4, 0x80, 0xa8, 0x92, 0xa4, 0xfd, 0x71, 0x23, 0x6e, 0x94, 0xd7, 0xa0, 0x45,
	0xc8, 0xf9, 0x01, 0x76, 0x4a, 0x9b, 0x0e, 0x48, 0x17, 0x76, 0xd5, 0xe4, 0x89, 0x62, 0xe4, 0xaf,
	0x90, 0x2b, 0x62, 0xa5, 0x8c, 0x58, 0x85, 0xc9, 0x0e, 0x54, 0x2e, 0xbc, 0x74, 0x18, 0x55, 0x2e,
	0x3c, 0xa5, 0xfb, 0xc1, 0xe5, 0x24, 0xd5, 0xa5, 0xfa, 0xeb, 0xfc, 0x6e, 0x65, 0x1f, 0x2c, 0xb3,
	0x3d, 0x79, 0x02, 0x0d, 0x6d, 0x5f, 0xbd, 0x99, 0x67, 0xa4, 0x5f, 0x02, 0x4a, 0xc2, 0xda, 0xb8,
	0xf0, 0xb2, 0x2f, 0x5f, 0x6a, 0x92, 0x67, 0x50, 0x19, 0x32, 0x7d, 0x54, 0xb3, 0xff, 0x78, 0xad,
	0x5e, 0xcb, 0x8a, 0xd0, 0xca, 0x90, 0x91, 0x2e, 0x58, 0x03, 0x9d, 0x45, 0xb3, 0x7f, 0x78, 0xe7,
	0xda, 0x01, 0xb5, 0x06, 0xce, 0x6b, 0x80, 0x1c, 0xd4, 0x5f, 0xc0, 0x4b, 0x6e, 0x78, 0x99, 0x09,
	0x32, 0xb3, 0x55, 0xa5, 0x2f, 0x99, 0x90, 0xe6, 0x8b, 0x98, 0x66, 0x57, 0x40, 0x48, 0x0f, 0x6a,
	0x7a, 0x13, 0x2d, 0x8b, 0x66, 0xdf, 0xbe, 0xeb, 0x5c, 0x6a, 0x96, 0x7d, 0x7c, 0x03, 0x8d, 0x9c,
	0x70, 0xe4, 0x21, 0x10, 0x7a, 0x7a, 0x35, 0xa1, 0xe7, 0xd7, 0xc7, 0xe3, 0xf1, 0xf5, 0xe0, 0xeb,
	0xe3, 0xf3, 0xe1, 0xa9, 0xbb, 0xf7, 0x1e, 0xd9, 0x85, 0x66, 0x8a, 0xbb, 0x23, 0xef, 0x6a, 0xcf,
	0x22, 0xfb, 0xd0, 0x4e, 0x81, 0x93, 0xc9, 0xb9, 0x3b, 0x3e, 0xdd, 0xab, 0x90, 0x47, 0xf0, 0xa0,
	0x04, 0xe5, 0xe1, 0xd5, 0xfe, 0x6f, 0x15, 0x80, 0x61, 0x24, 0x4f, 0x4c, 0x1a, 0xe4, 0x4b, 0xa8,
	0x99, 0xbb, 0x6c, 0x48, 0xce, 0x50, 0xec, 0xf0, 0xd1, 0x06, 0x8f, 0xf9, 0x70, 0x1d, 0x59, 0xe4,
	0x5b, 0xb8, 0xaf, 0xa1, 0x33, 0xce, 0x66, 0x05, 0x45, 0x3c, 0xde, 0x4c, 0x5f, 0xed, 0x7c, 0xcb,
	0x86, 0x5d, 0xeb, 0xc8, 0x22, 0x5f, 0x40, 0x7b, 0x88, 0x32, 0x1f, 0xab, 0x82, 0xb4, 0x7a, 0xe9,
	0x43, 0xe9, 0x15, 0x8b, 0xc2, 0x72, 0x74, 0x79, 0x4c, 0xbf, 0x84, 0xdd, 0x21, 0xca, 0xd2, 0x0c,
	0x29, 0xc7, 0xae, 0xa7, 0xb5, 0x1c, 0xd2, 0x27, 0x43, 0x78, 0x9a, 0xa0, 0x2c, 0xbe, 0xd7, 0xd4,
	0x5b, 0xad, 0x10, 0xf1, 0xfd, 0x87, 0xff, 0xe9, 0x59, 0x78, 0xb3, 0xad, 0x1f, 0x74, 0x9f, 0xfd,
	0x13, 0x00, 0x00, 0xff, 0xff, 0x13, 0xd3, 0x21, 0xde, 0x46, 0x0a, 0x00, 0x00,
}
