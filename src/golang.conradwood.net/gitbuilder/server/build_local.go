package main

import (
	"fmt"
	pb "golang.conradwood.net/apis/gitbuilder"
	"golang.conradwood.net/gitbuilder/builder"
	"golang.conradwood.net/gitbuilder/filetransfer"
	"golang.conradwood.net/go-easyops/utils"
	"sync"
)

var (
	builddirlock sync.Mutex
)

func get_local_build_dir() string {
	builddirlock.Lock()
	defer builddirlock.Unlock()
	return "/tmp/x/buildlocal"
}

func (e *echoServer) BuildFromLocalFiles(srv pb.GitBuilder_BuildFromLocalFilesServer) error {
	build_dir := get_local_build_dir()
	t, err := filetransfer.New(build_dir)
	if err != nil {
		return err
	}
	var blr *pb.BuildLocalRequest
	for {
		r, err := srv.Recv()
		if err != nil {
			fmt.Printf("Failed to recv: %s\n", err)
			return err
		}
		if r.FileTransfer != nil {
			err = t.Receive(r.FileTransfer)
			if err != nil {
				fmt.Printf("Failed to store file:%s", err)
				return err
			}
		}
		if r.Request != nil {
			blr = r.Request
			t.Close()
			break
		}
	}
	if blr == nil {
		return fmt.Errorf("No build request received")
	}
	fmt.Printf("Starting build for \"%s\"\n", blr.RepoName)
	sw := &localwriter{srv: srv}
	req := &pb.BuildRequest{
		GitURL:              "http://localhost/localbuild",
		CommitID:            "local",
		BuildNumber:         blr.BuildNumber,
		RepositoryID:        blr.RepositoryID,
		RepoName:            blr.RepoName,
		ArtefactName:        blr.ArtefactName,
		ArtefactID:          blr.ArtefactID,
		ExcludeBuildScripts: []string{"DIST"},
	}
	bd, err := builder.NewBuilder(build_dir, sw, req.BuildNumber, &builder.StandardBuildInfo{Req: req})

	if err != nil {
		return err
	}
	ctx := srv.Context()
	berr := bd.BuildAll(ctx)
	br := &pb.BuildResponse{
		Complete:   true,
		Success:    true,
		LogMessage: "logmessage",
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

type localwriter struct {
	srv pb.GitBuilder_BuildServer
}

func (s *localwriter) Write(buf []byte) (int, error) {
	err := s.srv.Send(&pb.BuildResponse{Stdout: buf})
	if err != nil {
		return 0, err
	}
	return len(buf), nil
}
