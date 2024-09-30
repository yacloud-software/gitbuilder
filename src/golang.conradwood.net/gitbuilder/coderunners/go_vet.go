package coderunners

import (
	"context"
	"strings"

	"golang.conradwood.net/go-easyops/linux"
	"golang.conradwood.net/go-easyops/utils"
)

type go_vet struct {
}

/*
in every directory with a go.mod file,
start "go vet" in each sub directory containing at least one .go file
*/
func (g go_vet) Run(ctx context.Context, b brunner) error {
	var gomods []string
	utils.DirWalk(b.GetRepoPath(), func(root, rel string) error {
		if !strings.HasSuffix(rel, "go.mod") {
			return nil
		}
		gomods = append(gomods, rel)
		return nil
	})

	failed_at_least_one := false
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
			d := rel[idx+1:]
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
			l := linux.New()
			res := "PASSED"
			com := []string{"go", "vet"}
			out, err := l.SafelyExecute(com, nil)
			if err != nil {
				failed_at_least_one = true
				res = "FAILED"
				b.Printf("go vet failed:\n%s\n", out)
				//				return errors.Errorf("vet failed for \"%s\"\n", subdir)
			}
			b.Printf("%s   %s (exists=%v)\n", res, subdir, exists)

		}
	}

	b.Printf("go-vet overall pass: %v\n", failed_at_least_one)
	return nil
}
