package common

import (
	"flag"
)

var (
	DistDirs       = []string{"configs", "scripts", "templates", "extra", "lib", "dist"}
	workdir        = flag.String("workdir", "/tmp/gitbuilder", "workdir for repos")
	maxprocs       = flag.Int("maxprocs", 4, "max processes during compile/check")
	xgocache       = flag.String("override_gocache", "", "if set use this as gocache. do not use in production")
	xgoproxyhost   = flag.String("goproxyhost", "goproxy.conradwood.net", "set the goproxy to this host (e.g. golang.conradwood.net)")
	xgoproxydirect = flag.Bool("goproxy_direct", false, "if true, add ',direct' to goproxy. (e.g. golang.conradwood.net needs to fallback to retrieve directly)")
)

func GetGoProxyHost() string {
	return *xgoproxyhost
}
func GetGoProxyDirect() bool {
	return *xgoproxydirect
}

func GetGoCache() string {
	return *xgocache
}

func GetMaxProcs() int {
	return *maxprocs
}
func WorkDir() string {
	return *workdir
}
