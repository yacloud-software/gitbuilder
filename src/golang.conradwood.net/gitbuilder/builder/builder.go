package builder

import (
	"fmt"
	"golang.conradwood.net/gitbuilder/buildinfo"
	"golang.conradwood.net/gitbuilder/buildrules"
	"golang.conradwood.net/gitbuilder/common"
	"io"
	"os"
	"time"
)

type Builder struct {
	buildrules *buildrules.BuildRules
	path       string // path containing .git
	stdout     io.Writer
	buildid    uint64 // for the build management system
	timestamp  time.Time
	buildinfo  buildinfo.BuildInfo // for the scripts and coderunners
	printer    *common.LinePrinter
}

func (b *Builder) BuildInfo() buildinfo.BuildInfo {
	return b.buildinfo
}

func builder_start() {
	l, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("os.Getwd: %s", err))
	}
	locpath = l
}

func NewBuilder(repopath string, stdout io.Writer, buildid uint64, bi buildinfo.BuildInfo) (*Builder, error) {
	builder_start()
	b := &Builder{path: repopath,
		stdout:    stdout,
		buildid:   buildid,
		buildinfo: bi,
		timestamp: time.Now(),
		printer:   &common.LinePrinter{MaxLineLength: 256, Prefix: fmt.Sprintf("[builder %s] ", bi.RepositoryName())},
	}
	br, err := buildrules.Read(b, b.GetRepoPath()+"/BUILD_RULES")
	if err != nil {
		return nil, err
	}
	b.buildrules = br
	return b, nil
}

// get the names of all buildscripts we currently know
func GetBuildScriptNames() []string {
	var res []string
	for k, _ := range buildrules.BUILD_SCRIPTS {
		res = append(res, k)
	}
	return res
}
func (b *Builder) Printf(txt string, args ...interface{}) {
	s := fmt.Sprintf("[builder %s] ", b.buildinfo.RepositoryName())
	s = fmt.Sprintf(s+txt, args...)
	b.printer.Printf(txt, args...)
	if b.stdout != nil {
		b.stdout.Write([]byte(s))
	} else {
		if len(s) > 500 {
			s = s[:500]
		}
		fmt.Print(s)
	}
}

// get the directory containing '.git'
func (b *Builder) GetRepoPath() string {
	return b.path
}
