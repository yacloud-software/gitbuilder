package coderunners

import (
	"context"
	"io/ioutil"
	"strings"
)

// find files in a directory, return filenames relative to dir
func FindFiles(ctx context.Context, dir string, name string) ([]string, error) {
	return findFilesWithSuffix(ctx, dir, "", name)
}

func findFilesWithSuffix(ctx context.Context, dir, suffix string, name string) ([]string, error) {
	files, err := ioutil.ReadDir(dir + "/" + suffix)
	if err != nil {
		return nil, err
	}
	var res []string
	for _, f := range files {
		if f.IsDir() {
			tres, err := findFilesWithSuffix(ctx, dir, suffix+"/"+f.Name(), name)
			if err != nil {
				return nil, err
			}
			res = append(res, tres...)
		}
		if f.Name() == name {
			fname := suffix + "/" + f.Name()
			fname = strings.TrimPrefix(fname, "/")
			res = append(res, fname)
		}
	}
	return res, nil

}



