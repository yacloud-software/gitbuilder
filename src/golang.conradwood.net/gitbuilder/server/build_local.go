package main

import (
	"fmt"
	pb "golang.conradwood.net/apis/gitbuilder"
	"golang.conradwood.net/gitbuilder/builder"
	"golang.conradwood.net/gitbuilder/common"
	"golang.conradwood.net/gitbuilder/filetransfer"
	"golang.conradwood.net/go-easyops/utils"
	"os"
	"sync"
	"time"
)

var (
	builddirlock sync.Mutex
	build_dirs   []*build_local_dir
)

type build_local_dir struct {
	dirname          string
	created          time.Time
	seq              int
	please_remove_me bool
	delete_me        bool
	deleted          bool
}

func init() {
	go build_local_cleaner()
}
func get_local_build_dir() string {
	builddirlock.Lock()
	defer builddirlock.Unlock()
	i := 0
	dirname := ""
	for {
		dirname = fmt.Sprintf("%s/buildlocal/%d", common.WorkDir(), i)
		if !utils.FileExists(dirname) {
			break
		}
		i++
	}
	bld := &build_local_dir{dirname: dirname, created: time.Now(), seq: i}
	build_dirs = append(build_dirs, bld)
	return bld.dirname
}
func remove_build_dir(dir string) {
	builddirlock.Lock()
	defer builddirlock.Unlock()
	for _, bd := range build_dirs {
		if bd.dirname == dir {
			bd.please_remove_me = true
		}
	}

}

func (e *echoServer) BuildFromLocalFiles(srv pb.GitBuilder_BuildFromLocalFilesServer) error {
	build_dir := get_local_build_dir()
	defer remove_build_dir(build_dir)
	err := utils.RecreateSafely(build_dir)
	if err != nil {
		return err
	}
	dd := &Dirdiff{build_dir}
	t, err := filetransfer.NewReceiver(build_dir)
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
	hint, err := dd.Remember()
	if err != nil {
		return err
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
		LogMessage: "",
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

	fmt.Printf("Now sending back the result...\n")
	changed_files, err := dd.ChangedFiles(hint)
	if err != nil {
		return err
	}
	var filenames_to_send []string
	for _, cf := range changed_files {
		filenames_to_send = append(filenames_to_send, cf.RelativeFilename())
	}
	sender := filetransfer.NewSender(srv, send_function)
	err = sender.SendSomeFiles(build_dir, filenames_to_send)
	if err != nil {
		return err
	}
	fmt.Printf("Buildlocal Done\n")
	return nil
}

func send_function(opaque interface{}, filename string, data []byte) error {
	br := &pb.BuildResponse{
		FileTransfer: &pb.FileTransferPart{Filename: filename, Data: data},
	}
	err := opaque.(pb.GitBuilder_BuildFromLocalFilesServer).Send(br)
	return err
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

func build_local_cleaner() {
	for {
		time.Sleep(time.Duration(1) * time.Minute)

		builddirlock.Lock()
		for _, bd := range build_dirs {
			if !bd.please_remove_me {
				continue
			}
			err := os.Rename(bd.dirname, bd.dirname+"_DEL")
			if err == nil {
				fmt.Printf("build local dir %s marked for deletion\n", bd.dirname)
				bd.dirname = bd.dirname + "_DEL"
				bd.delete_me = true
			} else {
				err = utils.RemoveAll(bd.dirname)
				if err != nil {
					fmt.Printf("Failed to delete %s: %s\n", bd.dirname, err)
				} else {
					bd.deleted = true
				}
			}

		}
		builddirlock.Unlock()

		// delete those which were renamed and marked as such
		for _, bd := range build_dirs {
			if !bd.delete_me {
				continue
			}
			err := utils.RemoveAll(bd.dirname)
			if err != nil {
				fmt.Printf("Failed to delete %s: %s\n", bd.dirname, err)
			} else {
				bd.deleted = true
			}
		}

		builddirlock.Lock()

		// remove the deleted ones from the list
		var nbuild []*build_local_dir
		for _, bd := range build_dirs {
			if bd.deleted {
				continue
			}
			nbuild = append(nbuild, bd)
		}
		build_dirs = nbuild

		builddirlock.Unlock()

	}
}
