package coderunners

import (
	"fmt"
	"golang.conradwood.net/gitbuilder/buildinfo"
)

type brunner_proxy struct {
	prefix string
	brun   brunner
}

func (b *brunner_proxy) Printf(txt string, args ...interface{}) {
	s := fmt.Sprintf(txt, args...)
	s = b.prefix + s
	b.brun.Printf("%s", s)
}
func (b *brunner_proxy) GetRepoPath() string {
	return b.brun.GetRepoPath()
}
func (b *brunner_proxy) BuildInfo() buildinfo.BuildInfo {
	return b.brun.BuildInfo()
}
