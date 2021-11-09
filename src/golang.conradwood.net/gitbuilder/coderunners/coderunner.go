package coderunners

import (
	"context"
	"fmt"
	"golang.conradwood.net/gitbuilder/buildinfo"
)

type brunner interface {
	Printf(txt string, args ...interface{})
	GetRepoPath() string
	BuildInfo() buildinfo.BuildInfo
}
type runner interface {
	Run(ctx context.Context, builder brunner) error
}

func Run(ctx context.Context, builder brunner, name string) (bool, error) {
	fmt.Printf("[coderunner ] %s\n", name)
	var g runner
	if name == "coderunner-gomodule" {
		g = gomodule{}
	} else if name == "coderunner-go-version" {
		g = goversion{}
	}
	if g == nil {
		return false, nil
	}
	err := g.Run(ctx, builder)
	if err != nil {
		return true, err
	}
	return true, nil
}
