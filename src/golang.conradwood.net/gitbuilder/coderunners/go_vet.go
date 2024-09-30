package coderunners

import (
	"context"
	"strings"

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

	for _, gomod := range gomods {
		subdirs := make(map[string]bool)
		root := b.GetRepoPath() + "/" + gomod
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
			}
			b.Printf("   %s (exists=%v)\n", subdir, exists)

		}
	}

	b.Printf("go-vet currently only has a stub\n")
	return nil
}
