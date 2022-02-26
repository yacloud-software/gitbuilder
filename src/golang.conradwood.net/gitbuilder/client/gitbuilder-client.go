package main

import (
	"context"
	"flag"
	"fmt"
	"golang.conradwood.net/apis/common"
	pb "golang.conradwood.net/apis/gitbuilder"
	_ "golang.conradwood.net/gitbuilder/appinfo"
	"golang.conradwood.net/gitbuilder/builder"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/utils"
	"io"
	"os"
	"strings"
	"time"
)

var (
	tags          = flag.String("tags", "", "routing tags to choose")
	echoClient    pb.GitBuilderClient
	f_url         = flag.String("url", "", "git url to build")
	f_dir         = flag.String("dir", "", "if not-empty, use this local working copy to build locally instead of calling the server")
	f_buildnumber = flag.Uint("build", 0, "buildnumber to use for this build")
	f_commitid    = flag.String("commitid", "", "commit id to set repository at for build")
	f_name        = flag.String("name", "", "repo and artefact name")
	f_repoid      = flag.Uint("repoid", 0, "repository id for scripts")
	status        = flag.Bool("status", false, "print status of gitbuilder server")
)

func main() {
	flag.Parse()

	echoClient = pb.GetGitBuilderClient()

	// a context with authentication
	authremote.Context()
	ctx := authremote.ContextWithTimeout(time.Duration(5) * time.Minute)
	ctx = addTags(ctx)
	if *status {
		printStatus(ctx)
		os.Exit(0)
	}
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
func addTags(ctx context.Context) context.Context {
	if *tags == "" {
		return ctx
	}
	rtags := make(map[string]string)
	vals := strings.Split(*tags, ",")
	for _, v := range vals {
		kv := strings.SplitN(v, "=", 2)
		if len(kv) != 2 {
			s := fmt.Sprintf("Invalid keyvalue tag: \"%s\" - it splits into %d parts instead of 2\n", v, len(kv))
			panic(s)
		}
		tk := kv[0]
		tv := kv[1]
		fmt.Printf("Adding tag \"%s\" with value \"%s\"\n", tk, tv)
		rtags[tk] = tv
	}
	return authremote.DerivedContextWithRouting(ctx, rtags)
}
func printStatus(ctx context.Context) {
	repolist, err := echoClient.GetLocalRepos(ctx, &common.Void{})
	utils.Bail("failed to get repos", err)
	t := &utils.Table{}
	t.AddHeaders("WorkDir")
	for _, repo := range repolist.Repos {
		t.AddString(repo.WorkDir)
		t.NewRow()
	}
	fmt.Println(t.ToPrettyString())
}
