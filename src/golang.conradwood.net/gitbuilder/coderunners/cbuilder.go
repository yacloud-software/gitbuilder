package coderunners

import (
	"context"
	"golang.conradwood.net/go-easyops/linux"
	"golang.conradwood.net/go-easyops/utils"
	"io/fs"
	"io/ioutil"
	"time"
)

type cbuilder struct{}

func (c *cbuilder) Run(ctx context.Context, builder brunner) error {
	srcdir := builder.GetRepoPath() + "/c"
	distdir := builder.GetRepoPath() + "/dist/"
	if !utils.FileExists(srcdir) {
		builder.Printf("WARNING - cbuilder invoked, but directory %s does not exist", srcdir)
		return nil
	}
	subdirs, err := ioutil.ReadDir(srcdir)
	if err != nil {
		return err
	}

	var cdirs []fs.FileInfo
	for _, c := range subdirs {
		if !c.IsDir() {
			continue
		}
		cdirs = append(cdirs, c)

	}
	if len(cdirs) == 0 {
		builder.Printf("WARNING - cbuilder invoked, but directory %s has 0 subdirectories", srcdir)
		return nil
	}

	builder.Printf("C-Builder dist: \"%s\"\n", distdir)
	for _, c := range subdirs {
		ffname := srcdir + "/" + c.Name()
		l := linux.New()
		l.SetMaxRuntime(time.Duration(5) * time.Minute)
		com := []string{
			"make",
			"all",
			"DIST=" + distdir,
		}
		relname := "c/" + c.Name()
		builder.Printf("Compiling \"%s\"...\n", relname)
		b, err := l.SafelyExecuteWithDir(com, ffname, nil)
		if err != nil {
			builder.Printf("Compile %s failed:%s\n", relname, b)
			return err
		}
		builder.Printf("Compiled  \"%s\"\n", relname)
	}
	return nil
}