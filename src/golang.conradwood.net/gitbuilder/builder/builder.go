package builder

import (
	"fmt"
	"golang.conradwood.net/gitbuilder/buildinfo"
	"io"
	"time"
)

var (
	// either name of scripts or coderunners. order of the array matters
	BUILD_SCRIPTS = map[string][]string{
		"STANDARD_PROTOS": []string{"protos-build.sh"},
		"STANDARD_GO":     []string{"coderunner-go-version", "go-build.sh"},
		"KICAD":           []string{"kicad-build.sh"},
		"STANDARD_JAVA":   []string{"java-build.sh"},
		"AUTOBUILD_SH":    []string{"autobuild.sh"},
		"CLEAN":           []string{"clean-build.sh"},
		"DIST":            []string{"dist.sh"},
		"GO_VERSION":      []string{"coderunner-go-version"},
		"GO_MODULES":      []string{"coderunner-gomodule"},
	}
)

type Builder struct {
	buildrules *BuildRules
	path       string // path containing .git
	stdout     io.Writer
	buildid    uint64 // for the build management system
	timestamp  time.Time
	buildinfo  buildinfo.BuildInfo // for the scripts and coderunners
}

func (b *Builder) BuildInfo() buildinfo.BuildInfo {
	return b.buildinfo
}
func NewBuilder(repopath string, stdout io.Writer, buildid uint64, bi buildinfo.BuildInfo) (*Builder, error) {
	b := &Builder{path: repopath,
		stdout:    stdout,
		buildid:   buildid,
		buildinfo: bi,
		timestamp: time.Now(),
	}
	err := b.readBuildrules()
	if err != nil {
		return nil, err
	}
	return b, nil
}
func (b *Builder) Printf(txt string, args ...interface{}) {
	s := fmt.Sprintf("[builder %s] ", b.buildinfo.RepositoryName())
	s = fmt.Sprintf(s+txt, args...)
	fmt.Print(s)
	if b.stdout != nil {
		b.stdout.Write([]byte(s))
	}
}

// get the directory containing '.git'
func (b *Builder) GetRepoPath() string {
	return b.path
}
