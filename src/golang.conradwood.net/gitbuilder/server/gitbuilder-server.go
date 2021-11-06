package main

import (
	"flag"
	"fmt"
	pb "golang.conradwood.net/apis/gitbuilder"
	"golang.conradwood.net/git"
	"golang.conradwood.net/gitbuilder/builder"
	"golang.conradwood.net/go-easyops/server"
	"golang.conradwood.net/go-easyops/utils"
	"google.golang.org/grpc"
	"os"
)

var (
	port = flag.Int("port", 4100, "The grpc server port")
)

type echoServer struct {
}

func main() {
	var err error
	flag.Parse()
	fmt.Printf("Starting GitBuilderServer...\n")

	sd := server.NewServerDef()
	sd.Port = *port
	sd.Register = server.Register(
		func(server *grpc.Server) error {
			e := new(echoServer)
			pb.RegisterGitBuilderServer(server, e)
			return nil
		},
	)
	err = server.ServerStartup(sd)
	utils.Bail("Unable to start server", err)
	os.Exit(0)
}

/************************************
* grpc functions
************************************/

func (e *echoServer) Build(req *pb.BuildRequest, srv pb.GitBuilder_BuildServer) error {
	fmt.Printf("Building: %#v\n", req)

	ctx := srv.Context()
	sw := &serverwriter{srv: srv}
	lr, err := git.GetLocalRepo(ctx, req.GitURL, nil, sw)
	if err != nil {
		return err
	}
	err = lr.Checkout(ctx, req.CommitID)
	if err != nil {
		return err
	}
	defer lr.Release()

	bd, err := builder.NewBuilder(lr.GitRepoPath(), sw, req.BuildNumber, &builder.StandardBuildInfo{
		Commit:       req.CommitID,
		RepoID:       req.RepositoryID,
		RepoName:     req.RepoName,
		ArtefactName: req.ArtefactName,
	})
	if err != nil {
		return err
	}
	err = bd.BuildAll(ctx)
	if err != nil {
		return err
	}
	message := `
this is a test message ,
because the server is not yet fully implemented.

thanks for trying it out
`
	err = srv.Send(&pb.BuildResponse{Stdout: []byte(message)})
	if err != nil {
		return err
	}
	return nil
}

type serverwriter struct {
	srv pb.GitBuilder_BuildServer
}

func (s *serverwriter) Write(buf []byte) (int, error) {
	err := s.srv.Send(&pb.BuildResponse{Stdout: buf})
	if err != nil {
		return 0, err
	}
	return len(buf), nil
}
