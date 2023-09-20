package coderunners

import (
	"context"
	"flag"
	"golang.yacloud.eu/yatools/repomodifier"
)

var (
	enable_static_check = flag.Bool("enable_static_check", false, "if true, also run go static check tool")
)

type staticcheck struct {
}

func (g *staticcheck) Run(ctx context.Context, b brunner) error {
	if !*enable_static_check {
		b.Printf("static check disabled\n")
		return nil
	}
	b.Printf("static check...\n")
	rc, err := repomodifier.NewRepoChangerFromDir(b.GetRepoPath())
	if err != nil {
		return err
	}
	defer rc.Close()
	err = rc.RunCommandInPackageDirs([]string{
		"/bin/bash",
		"-c",
		"staticcheck `go list -m`",
	})
	if err != nil {
		return err
	}
	return nil
}
