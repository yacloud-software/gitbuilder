package coderunners

import (
	"context"
	"flag"
	"fmt"
	"golang.conradwood.net/gitbuilder/buildinfo"
)

var (
	use_internal_proto_builder = flag.Bool("use_internal_proto_builder", true, "if false, use script protobuild.sh")
)

type brunner interface {
	Printf(txt string, args ...interface{})
	GetRepoPath() string
	BuildInfo() buildinfo.BuildInfo
}
type runner interface {
	Run(ctx context.Context, builder brunner) error
}

// returns true if it is a coderunner, false if it is not a coderunner
func Run(ctx context.Context, builder brunner, name string) (bool, error) {
	fmt.Printf("[coderunner ] %s\n", name)
	var g runner
	if name == "coderunner-gomodule" {
		g = gomodule{}
	} else if name == "coderunner-go-version" {
		g = goversion{}
	} else if name == "protos-build.sh" && *use_internal_proto_builder {
		g = protobuilder{}
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
