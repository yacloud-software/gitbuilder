package common

import (
	"flag"
)

var (
	DistDirs = []string{"configs", "scripts", "templates", "extra", "lib", "dist"}
	workdir  = flag.String("workdir", "/tmp/gitbuilder", "workdir for repos")
)

func WorkDir() string {
	return *workdir
}
