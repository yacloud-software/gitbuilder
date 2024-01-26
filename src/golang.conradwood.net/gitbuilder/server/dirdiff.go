package main

import (
	"crypto/md5"
	"fmt"
	"golang.conradwood.net/go-easyops/utils"
	"io"
	"os"
	"time"
)

// "remember" files in a dir and compare it later to the state again (finding changed files essentially)
type Dirdiff struct {
	dir string
}
type Hint struct {
	files map[string]*fileinfo
}

type fileinfo struct {
	md5           string
	last_modified time.Time
	exists        bool
}
type ChangedFile struct {
	removed  bool
	filename string
}

func (cf *ChangedFile) RelativeFilename() string {
	return cf.filename
}
func (cf *ChangedFile) String() string {
	if cf.removed {
		return fmt.Sprintf("%s (was removed)", cf.filename)
	}
	return fmt.Sprintf("%s", cf.filename)
}

// this is an I/O and CPU intensive calculation. every file is read fully and its MD5 (or SHA256) sum calculated
func (dd *Dirdiff) Remember() (*Hint, error) {
	h := &Hint{files: make(map[string]*fileinfo)}
	err := utils.DirWalk(dd.dir, func(root, rel string) error {
		fi, err := get_file_info(root + "/" + rel)
		if err != nil {
			return err
		}
		h.files[rel] = fi
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h, nil
}
func (dd *Dirdiff) ChangedFiles(h *Hint) ([]*ChangedFile, error) {
	now, err := dd.Remember()
	if err != nil {
		return nil, err
	}

	// now compare hint with "now"
	var res []*ChangedFile
	for filename, fi := range h.files {
		cf := now.get_file_change(filename, fi)
		if cf != nil {
			res = append(res, cf)
		}
	}

	// now find files that are "new"
	for filename, _ := range now.files {
		_, old := h.files[filename]
		if !old {
			// new file:
			cf := &ChangedFile{filename: filename}
			res = append(res, cf)
		}
	}

	return res, nil
}

// compare the fileinfo with this one and return a changedfile object if there are differences
func (h *Hint) get_file_change(filename string, fi *fileinfo) *ChangedFile {
	res := &ChangedFile{filename: filename}
	mfi, exists := h.files[filename]
	if !exists {
		// used to exist, but not any more
		res.removed = true
		return res
	}
	if !mfi.last_modified.Equal(fi.last_modified) {
		// timestamp changed
		return res
	}
	if mfi.md5 != fi.md5 {
		// contents changed
		return res
	}
	return nil
}

func get_file_info(filename string) (*fileinfo, error) {
	fi := &fileinfo{}
	if !utils.FileExists(filename) {
		fi.exists = false
		return fi, nil
	}

	fileInfo, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}
	fi.last_modified = fileInfo.ModTime()

	fi.md5, err = get_md5(filename)
	if err != nil {
		return nil, err
	}
	return fi, nil
}

func get_md5(filename string) (string, error) {

	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	x := fmt.Sprintf("%x", h.Sum(nil))
	return x, nil
}
