package builder

import (
	"fmt"
	"io"
	"time"
)

var (
	BUILD_SCRIPTS = map[string]string{
		"STANDARD_PROTOS": "protos-build.sh",
		"STANDARD_GO":     "go-build.sh",
		"KICAD":           "kicad-build.sh",
		"STANDARD_JAVA":   "java-build.sh",
	}
)

type BuildInfo interface {
	CommitID() string
	RepositoryID() uint64
	RepositoryName() string
	RepositoryArtefactName() string
}

type Builder struct {
	buildrules *BuildRules
	path       string // path containing .git
	stdout     io.Writer
	buildid    uint64
	timestamp  time.Time
	buildinfo  BuildInfo
}

func NewBuilder(repopath string, stdout io.Writer, buildid uint64, bi BuildInfo) (*Builder, error) {
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
	s := fmt.Sprintf("[builder] ")
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
