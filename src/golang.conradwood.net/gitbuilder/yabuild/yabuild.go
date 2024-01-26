package main

import (
	"flag"
	"fmt"
	pb "golang.conradwood.net/apis/gitbuilder"
	"golang.conradwood.net/gitbuilder/filetransfer"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/utils"
	"io"
	"os"
	"path/filepath"
	"time"
)

var (
	output_dir  = flag.String("output", "", "if set write compiled files to `directory`, otherwise build dir is used")
	ctx_timeout = flag.Duration("timeout", time.Duration(15)*time.Minute, "timeout for file-transfer and build")
)

func main() {
	flag.Parse()
	builddir := ""
	if len(flag.Args()) != 0 {
		builddir = flag.Args()[0]
	} else {
		builddir = utils.WorkingDir()
	}
	topdir, err := find_top_of_git_dir(builddir)
	utils.Bail("cannot determine git dir", err)

	fmt.Printf("Yabuilding %s...\n", topdir)

	if *output_dir != "" {
		utils.Bail("failed to create output dir", utils.RecreateSafely(*output_dir))
	}

	rdir := *output_dir
	if rdir == "" {
		rdir = topdir
	}

	rec, err := filetransfer.NewReceiver(rdir)
	utils.Bail("failed to create receiver", err)

	ctx := authremote.ContextWithTimeout(*ctx_timeout)
	srv, err := pb.GetGitBuilderClient().BuildFromLocalFiles(ctx)
	utils.Bail("failed to start build: %s\n", err)

	fmt.Printf("Sending files to server...\n")
	sender := filetransfer.NewSender(srv, send_function)
	err = sender.SendFiles(topdir)
	utils.Bail("failed to transfer files to server", err)
	fmt.Printf("Sent %d files to server\n", sender.FilesSent())

	fmt.Printf("Starting build...\n")
	blr := &pb.BuildLocalRequest{RepositoryID: 2,
		RepoName:    "local_build",
		BuildNumber: 5,
		ArtefactID:  2,
	}
	err = srv.Send(&pb.BuildLocalFiles{Request: blr})
	utils.Bail("failed to start build", err)

	printed := false
	for {
		r, err := srv.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			utils.Bail("failed to receive response from server", err)
		}
		if len(r.Stdout) != 0 {
			fmt.Printf("%s", string(r.Stdout))
		}
		if len(r.LogMessage) != 0 {
			fmt.Printf("%s", string(r.LogMessage))
		}
		if r.Success {
			fmt.Printf("**** Build successful *******\n")
		}
		ft := r.FileTransfer
		if ft != nil {
			if !printed {
				fmt.Printf("Receiving files (storing into %s)\n", rdir)
				printed = true
			}
			// receiving a file
			//fmt.Printf("receiving file %s\n", ft.Filename)
			err = rec.Receive(ft)
			utils.Bail("failed to receive file", err)
		}
	}
	fmt.Printf("Stored %d files in %s\n", rec.FilesReceived(), rdir)
}
func find_top_of_git_dir(builddir string) (string, error) {
	s := builddir
	for !utils.FileExists(s + "/.git/config") {
		s = filepath.Dir(s)
	}
	return s, nil
}
func send_function(opaque interface{}, filename string, fi os.FileInfo, data []byte) error {
	br := &pb.BuildLocalFiles{
		FileTransfer: &pb.FileTransferPart{Filename: filename, Data: data, Permissions: uint32(fi.Mode().Perm())},
	}
	err := opaque.(pb.GitBuilder_BuildFromLocalFilesClient).Send(br)
	return err
}
