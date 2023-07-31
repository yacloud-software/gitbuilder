package coderunners

import (
	"context"
	"golang.conradwood.net/go-easyops/linux"
	"golang.conradwood.net/go-easyops/utils"
	"io/fs"
	"io/ioutil"
)

type cbuilder struct{}

func (c *cbuilder) Run(ctx context.Context, builder brunner) error {
	srcdir := builder.GetRepoPath() + "/c"
	distdir := builder.GetRepoPath() + "/dist"
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

	for _, c := range subdirs {
		ffname := srcdir + "/" + c.Name()
		l := linux.New()
		com := []string{
			"make",
			"all",
			"DIST=" + distdir,
		}
		b, err := l.SafelyExecuteWithDir(com, ffname, nil)
		if err != nil {
			builder.Printf("Compile failed:\n", b)
			return err
		}
	}
	return nil
}
