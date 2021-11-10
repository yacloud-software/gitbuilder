package coderunners

import (
	"context"
	"fmt"
	"golang.conradwood.net/go-easyops/utils"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"
)

type goversion struct {
	brunner brunner
	bd      time.Time
}

func (g goversion) Run(ctx context.Context, b brunner) error {
	var err error
	g.brunner = b
	g.bd = time.Now()
	st := g.findFilesVendor()
	g.brunner.Printf("found %d files to modify (vendor mode)\n", len(st))
	for _, f := range st {
		err = g.buildfile(f)
		if err != nil {
			return err
		}
	}
	st, err = g.findFilesModule(ctx)
	if err != nil {
		return err
	}
	g.brunner.Printf("found %d files to modify (module mode)\n", len(st))
	src := g.brunner.GetRepoPath() + "/src"
	for _, f := range st {
		fname := src + "/" + f
		err = g.buildfile(fname)
		if err != nil {
			return err
		}
	}
	return nil
}
func (g goversion) BuildDate() time.Time {
	return g.bd
}

// find files as used by the modules, basically all files under src/ called "go_easyops_appversion.go"
func (g goversion) findFilesModule(ctx context.Context) ([]string, error) {
	src := g.brunner.GetRepoPath() + "/src"
	g.brunner.Printf("Finding files in \"%s\"\n", src)
	s, err := FindFiles(ctx, src, "go_easyops_appversion.go")
	if err != nil {
		return nil, err
	}
	return s, nil
}

// find files as used by the non-module vendor style repositories
func (g goversion) findFilesVendor() []string {
	src := g.brunner.GetRepoPath() + "/src"
	if !utils.FileExists(src) {
		return nil
	}
	tests := []string{
		"src/golang.conradwood.net/go-easyops/cmdline/myversion.go",
		"src/golang.conradwood.net/vendor/golang.conradwood.net/go-easyops/cmdline/appversion.go",
		"src/golang.singingcat.net/vendor/golang.conradwood.net/go-easyops/cmdline/appversion.go",
		"vendor/golang.conradwood.net/go-easyops/cmdline/appversion.go",
	}
	var res []string
	for _, rt := range tests {
		if utils.FileExists(rt) {
			res = append(res, rt)
		} else if utils.FileExists(g.brunner.GetRepoPath() + "/" + rt) {
			res = append(res, g.brunner.GetRepoPath()+"/"+rt)
		}
	}
	files, err := ioutil.ReadDir(src)
	utils.Bail("failed to read directory", err)
	for _, f := range files {
		for _, t := range tests {
			fname := src + "/" + f.Name() + "/" + t
			if utils.FileExists(fname) {
				res = append(res, fname)
			}
		}
	}
	return res
}

func (g goversion) buildfile(filename string) error {
	if filename == "" {
		fmt.Println("No filename!")
		return fmt.Errorf("No filename")
	}
	g.brunner.Printf("Modifying file: \"%s\"\n", filename)
	b := g.brunner.BuildInfo()
	sb, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	s := string(sb)
	// regexes must match original file as well as updated file
	// (a git repository might get patched and re-used)
	bnr := regexp.MustCompile("^(.*BUILD_NUMBER.*=.*) \\d+ (.*)$")
	bdesc := regexp.MustCompile(`^(.*BUILD_DESCRIPTION.*=.*").*(".*)$`)
	bts := regexp.MustCompile("^(.*BUILD_TIMESTAMP.*=.*) \\d+ (.*)$")
	brepoid := regexp.MustCompile("^(.*BUILD_REPOSITORY_ID.*=.*) \\d+ (.*)$")
	brepo := regexp.MustCompile(`^(.*BUILD_REPOSITORY .*=.*").*(".*)$`)
	comm := regexp.MustCompile(`^(.*BUILD_COMMIT.*=.*").*(".*)$`)
	ns := ""
	for _, line := range strings.Split(s, "\n") {

		repl := fmt.Sprintf("${1} %d ${2}", b.BuildNumber())
		line = bnr.ReplaceAllString(line, repl)

		repl = fmt.Sprintf("${1}Build #%d of %s at %s on host %s${2}",
			b.BuildNumber(), b.RepositoryName(), g.BuildDate(), os.Getenv("HOSTNAME"))
		line = bdesc.ReplaceAllString(line, repl)

		repl = fmt.Sprintf("${1} %d ${2}", g.BuildDate().Unix())
		line = bts.ReplaceAllString(line, repl)

		repl = fmt.Sprintf("${1} %d ${2}", b.RepositoryID())
		line = brepoid.ReplaceAllString(line, repl)

		repl = fmt.Sprintf("${1}%s${2}", b.RepositoryName())
		line = brepo.ReplaceAllString(line, repl)

		repl = fmt.Sprintf("${1}%s${2}", b.CommitID())
		line = comm.ReplaceAllString(line, repl)

		ns = ns + line + "\n"
	}
	err = ioutil.WriteFile(filename, []byte(ns), 0755)
	if err != nil {
		return err
	}
	return nil
}
