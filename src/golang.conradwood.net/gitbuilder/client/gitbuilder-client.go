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
	foo           = flag.Bool("foo", false, "if so submit a funny proto")
	inc_scripts   = flag.String("inc_scripts", "", "comma delimited list of scripts to include in run")
	ex_scripts    = flag.String("exc_scripts", "", "comma delimited list of scripts to exclude from run")
	tags          = flag.String("tags", "", "routing tags to choose")
	echoClient    pb.GitBuilderClient
	f_url         = flag.String("url", "", "git url to build")
	f_dir         = flag.String("dir", "", "if not-empty, use this local working copy to build locally instead of calling the server")
	f_buildnumber = flag.Uint("build", 0, "buildnumber to use for this build")
	f_commitid    = flag.String("commitid", "", "commit id to set repository at for build")
	f_name        = flag.String("name", "", "repo and artefact name")
	f_repoid      = flag.Uint("repoid", 0, "repository id for scripts")
	scripts       = flag.Bool("scripts", false, "print script names of all known scripts on builder server")
	status        = flag.Bool("status", false, "print status of gitbuilder server")
)

func main() {
	flag.Parse()

	echoClient = pb.GetGitBuilderClient()

	// a context with authentication
	authremote.Context()
	if *foo {
		utils.Bail("failed to foo", Foo())
		os.Exit(0)
	}
	ctx := authremote.ContextWithTimeout(time.Duration(5) * time.Minute)
	ctx = addTags(ctx)
	if *scripts {
		printScripts(ctx)
		os.Exit(0)
	}
	if *status {
		printStatus(ctx)
		os.Exit(0)
	}
	if *f_dir != "" {
		b, err := builder.NewBuilder(*f_dir, nil, uint64(*f_buildnumber),
			&builder.StandardBuildInfo{
				Req: &pb.BuildRequest{
					CommitID:     *f_commitid,
					RepositoryID: 1,
					RepoName:     "test_reponame",
					ArtefactName: "test_artefact",
					BuildNumber:  uint64(*f_buildnumber),
				},
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
		RepoName:            *f_name,
		ArtefactName:        *f_name,
		GitURL:              *f_url,
		CommitID:            *f_commitid,
		BuildNumber:         uint64(*f_buildnumber),
		RepositoryID:        uint64(*f_repoid),
		IncludeBuildScripts: include_build_scripts(),
		ExcludeBuildScripts: exclude_build_scripts(),
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
	return authremote.DerivedContextWithRouting(ctx, rtags, false)
}
func printStatus(ctx context.Context) {
	repolist, err := echoClient.GetLocalRepos(ctx, &common.Void{})
	utils.Bail("failed to get repos", err)
	t := &utils.Table{}
	t.AddHeaders("WorkDir", "inuse", "created", "released")
	for _, repo := range repolist.Repos {
		t.AddString(repo.WorkDir)
		t.AddBool(repo.InUse)
		t.AddTimestamp(repo.Created)
		if repo.InUse {
			t.AddString("---")
		} else {
			t.AddTimestamp(repo.Released)
		}
		t.NewRow()
	}
	fmt.Println(t.ToPrettyString())
}
func printScripts(ctx context.Context) {
	sn, err := echoClient.GetBuildScripts(ctx, &common.Void{})
	utils.Bail("failed to get buildscripts", err)
	for i, name := range sn.Names {
		fmt.Printf("%02d. %s\n", i+1, name)
	}

}

func exclude_build_scripts() []string {
	return cdl(*ex_scripts)
}
func include_build_scripts() []string {
	return cdl(*inc_scripts)
}

func cdl(cdl string) []string {
	var res []string
	if cdl == "" {
		return res
	}
	for _, s := range strings.Split(cdl, ",") {
		res = append(res, strings.Trim(s, " "))
	}
	return res
}
func Foo() error {
	br := &pb.BuildRequest{
		GitURL:              "https://apps.planetaryprocessing.io/gerrit/a/prober",
		CommitID:            "7062b41e1edc7d4f1416587682e85c676c50fbeb",
		ArtefactName:        "prober",
		RepoName:            "prober",
		ArtefactID:          0,
		BuildNumber:         0,
		ExcludeBuildScripts: []string{"DIST"},
		FetchURLS: []*pb.FetchURL{
			&pb.FetchURL{URL: "https://apps.planetaryprocessing.io/gerrit/a/prober", RefSpec: "refs/changes/92/92/1"},
		},
	}
	ctx := authremote.ContextWithTimeout(time.Duration(300) * time.Second)
	srv, err := pb.GetGitBuilderClient().Build(ctx, br)
	if err != nil {
		return err
	}
	for {
		bs, err := srv.Recv()
		if bs != nil && bs.Stdout != nil && len(bs.Stdout) > 0 {
			fmt.Printf("%s", string(bs.Stdout))
		} else {
			fmt.Printf("%#v\n", bs)
		}
		if err != nil {
			break
		}
	}
	return nil
}
