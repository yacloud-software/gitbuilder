package main

import (
	"flag"
	"fmt"
	pb "golang.conradwood.net/apis/gitbuilder"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/utils"
	"io"
	"os"
)

var (
	echoClient    pb.GitBuilderClient
	f_url         = flag.String("url", "", "git url to build")
	f_buildnumber = flag.Uint("build", 0, "buildnumber to use for this build")
	f_commitid    = flag.String("commitid", "", "commit id to set repository at for build")
)

func main() {
	flag.Parse()

	echoClient = pb.GetGitBuilderClient()

	// a context with authentication
	ctx := authremote.Context()

	empty := &pb.BuildRequest{
		GitURL:      *f_url,
		CommitID:    *f_commitid,
		BuildNumber: uint64(*f_buildnumber),
	}
	stream, err := echoClient.Build(ctx, empty)
	utils.Bail("Failed to ping server", err)
	for {
		pl, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			utils.Bail("failed to build", err)
		}
		if pl.Complete {
			fmt.Printf("Complete. Success=%v, Message: %s\n", pl.Success, pl.ResultMessage)
		}
		if len(pl.Stdout) > 0 {
			fmt.Printf(string(pl.Stdout))
		}
	}

	fmt.Printf("Done.\n")
	os.Exit(0)
}
