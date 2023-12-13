package main

import (
	"context"
	"flag"
	"fmt"
	"golang.conradwood.net/apis/common"
	pb "golang.conradwood.net/apis/gitbuilder"
	_ "golang.conradwood.net/gitbuilder/appinfo"
	"golang.conradwood.net/gitbuilder/builder"
	"golang.conradwood.net/gitbuilder/git"
	"golang.conradwood.net/go-easyops/auth"
	"golang.conradwood.net/go-easyops/errors"
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
	sd.SetPort(*port)
	sd.SetRegister(server.Register(
		func(server *grpc.Server) error {
			e := new(echoServer)
			pb.RegisterGitBuilderServer(server, e)
			return nil
		},
	))
	err = server.ServerStartup(sd)
	utils.Bail("Unable to start server", err)
	os.Exit(0)
}

/************************************
* grpc functions
************************************/
func (e *echoServer) GetBuildScripts(ctx context.Context, req *common.Void) (*pb.BuildScriptList, error) {
	res := &pb.BuildScriptList{
		Names: builder.GetBuildScriptNames(),
	}
	return res, nil
}
func (e *echoServer) GetLocalRepos(ctx context.Context, req *common.Void) (*pb.LocalRepoList, error) {
	return git.GetLocalRepos(), nil
}
func (e *echoServer) Build(req *pb.BuildRequest, srv pb.GitBuilder_BuildServer) error {
	u := auth.GetUser(srv.Context())
	fmt.Printf("Building (as user %s):\n", auth.UserIDString(u))
	if u == nil {
		fmt.Printf("WARNING!!! Building without user account. (from service %s)\n", auth.UserIDString(auth.GetService(srv.Context())))
		return errors.Unauthenticated(srv.Context(), "User account required to build")
	}
	fmt.Printf("#for gitbuilder-client: -url=\"%s\" -commitid=\"%s\" -build=%d -repoid=%d -name=%s -artefactid=%d\n", req.GitURL, req.CommitID, req.BuildNumber, req.RepositoryID, req.RepoName, req.ArtefactID)
	fmt.Printf("  Reponame: \"%s\", Artefactname: \"%s\"\n", req.RepoName, req.ArtefactName)

	ctx := srv.Context()
	sw := &serverwriter{srv: srv}
	lr, err := git.GetLocalRepo(ctx, req.GitURL, req.FetchURLs, !req.RequiresDeepClone, sw)
	if err != nil {
		return err
	}
	defer lr.Release()

	// this is done as part of "GetLocalRepo")
	/*
		for _, fu := range req.FetchURLs {
			fmt.Printf("Fetching %s %s\n", fu.URL, fu.RefSpec)
			err = lr.Fetch(ctx, fu)
			if err != nil {
				return err
			}
		}
	*/
	err = lr.Checkout(ctx, req.CommitID)
	if err != nil {
		return err
	}

	logmessage, err := lr.GetLogMessage(ctx)
	if err != nil {
		return err
	}

	bd, err := builder.NewBuilder(lr.GitRepoPath(), sw, req.BuildNumber, &builder.StandardBuildInfo{
		Req: req,
		/*
			Commit:       req.CommitID,
			RepoID:       req.RepositoryID,
			RepoName:     req.RepoName,
			ArtefactName: req.ArtefactName,
			Build:        req.BuildNumber,
		*/
	})
	if err != nil {
		return err
	}

	berr := bd.BuildAll(ctx)
	br := &pb.BuildResponse{
		Complete:   true,
		Success:    true,
		LogMessage: logmessage,
	}
	if berr != nil {
		br.Success = false
		br.ResultMessage = utils.ErrorString(berr)
	}
	err = srv.Send(br)
	if err != nil {
		return err
	}

	if berr != nil {
		return berr
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




