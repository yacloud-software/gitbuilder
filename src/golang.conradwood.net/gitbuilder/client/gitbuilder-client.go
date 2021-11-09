package main

import (
	"flag"
	"fmt"
	pb "golang.conradwood.net/apis/gitbuilder"
	_ "golang.conradwood.net/gitbuilder/appinfo"
	"golang.conradwood.net/gitbuilder/builder"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/utils"
	"io"
	"os"
	"time"
)

var (
	echoClient    pb.GitBuilderClient
	f_url         = flag.String("url", "", "git url to build")
	f_dir         = flag.String("dir", "", "if not-empty, use this local working copy to build locally instead of calling the server")
	f_buildnumber = flag.Uint("build", 0, "buildnumber to use for this build")
	f_commitid    = flag.String("commitid", "", "commit id to set repository at for build")
	f_name        = flag.String("name", "", "repo and artefact name")
	f_repoid      = flag.Uint("repoid", 0, "repository id for scripts")
)

func main() {
	flag.Parse()

	echoClient = pb.GetGitBuilderClient()

	// a context with authentication
	authremote.Context()
	ctx := authremote.ContextWithTimeout(time.Duration(5) * time.Minute)

	if *f_dir != "" {
		b, err := builder.NewBuilder(*f_dir, nil, uint64(*f_buildnumber),
			&builder.StandardBuildInfo{
				Commit:       *f_commitid,
				RepoID:       1,
				RepoName:     "test_reponame",
				ArtefactName: "test_artefact",
				Build:        uint64(*f_buildnumber),
			},
		)
		br := &builder.BuildRules{
			Builds: []string{"STANDARD_GO"},
		}
		//br.Builds = []string{"GO_MODULES"}
		utils.Bail("failed to get builder", err)
		err = b.BuildWithRules(ctx, br)
		utils.Bail("failed to build", err)
		os.Exit(0)
	}
	empty := &pb.BuildRequest{
		RepoName:     *f_name,
		ArtefactName: *f_name,
		GitURL:       *f_url,
		CommitID:     *f_commitid,
		BuildNumber:  uint64(*f_buildnumber),
		RepositoryID: uint64(*f_repoid),
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
