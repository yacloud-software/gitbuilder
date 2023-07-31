package coderunners

import (
	"archive/zip"
	"context"
	"fmt"
	"golang.conradwood.net/go-easyops/linux"
	"golang.conradwood.net/go-easyops/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type gomodule struct {
}

// find go.mod files, zip it and stuff it into "dist"
func (g *gomodule) Run(ctx context.Context, b brunner) error {
	version := fmt.Sprintf("v0.1.%d", b.BuildInfo().BuildNumber())

	// find go.mod files:
	d := b.GetRepoPath() + "/src/"
	b.Printf("Searching for go.mod files in \"%s\"\n", d)
	files, err := FindFiles(ctx, d, "go.mod")
	if err != nil {
		return err
	}
	b.Printf("found %d go.mod files\n", len(files))
	for _, f := range files {
		pkg := filepath.Dir(f)
		if strings.HasSuffix(pkg, "/tests") || strings.HasSuffix(pkg, "/test") {
			continue
		}
		if strings.Contains(f, "/vendor/") {
			continue
		}
		td := b.GetRepoPath() + "/dist/gomod/" + pkg + "/"
		os.MkdirAll(td, 0777)

		err = linux.Copy(d+"/"+f, td+"go.mod")
		if err != nil {
			return err
		}
		b.Printf("FILE: %s, package: \"%s\"\n", f, pkg)
		err = createZip(td+"mod.zip", b, pkg, version)
		if err != nil {
			return err
		}
	}

	return nil
}

// dir is typically "src/", pkg is something like "golang.conradwood.net/go-easyops"
func createZip(zipfile string, b brunner, pkg string, version string) error {
	zf, err := os.Create(zipfile)
	if err != nil {
		return err
	}
	b.Printf("zip file \"%s\" created...\n", zipfile)
	w := zip.NewWriter(zf)
	err = addzipfiles(w, b, pkg, "", version)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return nil
}

func addzipfiles(zf *zip.Writer, b brunner, pkg string, dir, version string) error {
	d := b.GetRepoPath() + "/src/" + pkg + "/" + dir
	b.Printf("Adding files in \"%s\"\n", d)
	files, err := ioutil.ReadDir(d)
	if err != nil {
		return err
	}
	for _, f := range files {
		if f.IsDir() {
			err = addzipfiles(zf, b, pkg, dir+"/"+f.Name(), version)
			if err != nil {
				return err
			}
			continue
		}
		fname := dir + "/" + f.Name()
		fname = strings.TrimPrefix(fname, "/")
		zname := pkg + "@" + version + "/" + fname
		b.Printf("  adding \"%s\" as \"%s\"\n", fname, zname)
		b, err := utils.ReadFile(d + "/" + f.Name())
		if err != nil {
			return err
		}
		zff, err := zf.Create(zname)
		if err != nil {
			return err
		}
		_, err = zff.Write(b)
		if err != nil {
			return err
		}
	}
	return nil
}
