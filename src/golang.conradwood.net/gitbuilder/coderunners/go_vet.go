package coderunners

import (
	"context"
	"flag"
	"fmt"
	"strings"
	"time"

	"golang.conradwood.net/gitbuilder/common"
	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/linux"
	"golang.conradwood.net/go-easyops/utils"
)

var (
	fail_with_govet    = flag.Bool("fail_go_vet", true, "if false, go vet failures will not result in errors")
	known_non_vettable = []string{
		"github.com",
		"subs",
		"/vendor",
	}
)

type go_vet struct {
}

/*
in every directory with a go.mod file,
start "go vet" in each sub directory containing at least one .go file
*/
func (g go_vet) Run(ctx context.Context, b brunner) error {
	gocompiler := common.FindExecutable("go")
	if gocompiler == "" {
		return errors.Errorf("no go compiler found")
	}
	fmt.Printf("go-vet using go at: %s\n", gocompiler)
	var gomods []string
	non_vettable := known_non_vettable
	err := utils.DirWalk(b.GetRepoPath(), func(root, rel string) error {
		// ignore if non go.mod
		if !strings.HasSuffix(rel, "go.mod") {
			return nil
		}

		// ignore patterns in non_vettable
		nv := false
		for _, nvs := range non_vettable {
			if strings.Contains(rel, nvs) {
				nv = true
				break
			}
		}
		if nv {
			return nil
		}

		// add to list of go.mod files
		gomods = append(gomods, rel)
		return nil
	})
	if err != nil {
		return err
	}
	failed_at_least_one := false
	env := common.StdEnv(ctx, b)
	for _, gomod := range gomods {
		gomoddir := strings.TrimSuffix(gomod, "/go.mod")

		subdirs := make(map[string]bool)
		root := b.GetRepoPath() + "/" + gomoddir
		utils.DirWalk(root, func(root, rel string) error {
			if !strings.HasSuffix(rel, ".go") {
				return nil
			}
			idx := strings.LastIndex(rel, "/")
			if idx == -1 {
				return nil
			}
			d := rel[:idx]
			subdirs[d] = true
			return nil
		})

		b.Printf("go.mod file %s has %d subdirs:\n", gomod, len(subdirs))
		for subdir, _ := range subdirs {
			ffname := root + "/" + subdir
			exists := utils.FileExists(ffname)
			if !exists {
				b.Printf("%s does not exist\n", ffname)
				continue
			}
			// ignore patterns in non_vettable
			nv := ""
			for _, nvs := range non_vettable {
				if strings.Contains(ffname, nvs) {
					nv = nvs
					break
				}
			}
			if nv != "" {
				b.Printf("%s is non_vettable (contains \"%s\")\n", ffname, nv)
				continue
			}
			l := linux.New()
			l.SetMaxRuntime(time.Duration(180) * time.Second)
			res := "PASSED"

			com := []string{gocompiler, "vet"}
			vdir := ffname
			l.SetEnvironment(env)
			out, err := l.SafelyExecuteWithDir(com, vdir, nil)
			if err != nil {
				failed_at_least_one = true
				res = "FAILED"
				b.Printf("go vet in %s failed (%s):\n%s\n", vdir, err, out)
				//				return errors.Errorf("vet failed for \"%s\"\n", subdir)
			}
			b.Printf("%s   %s\n", res, subdir)

		}
	}

	b.Printf("go-vet overall pass: failed=%v\n", failed_at_least_one)
	if failed_at_least_one && *fail_with_govet {
		return errors.Errorf("GO VET failed")
	}
	return nil
}
