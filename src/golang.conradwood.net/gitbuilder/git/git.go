package git

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	pb "golang.conradwood.net/apis/gitbuilder"
	"golang.conradwood.net/gitbuilder/common"
	"golang.conradwood.net/go-easyops/linux"
	"golang.conradwood.net/go-easyops/utils"
)

var (
	repos          []*LocalRepo
	with_recursive = flag.Bool("git_with_recursive", true, "if true use git-clone --recursive")
	wd_lock        sync.Mutex
	workdir_ctr    = 0
	recreated      = false
)

func init() {
	go cleaner()
}
func workdir() string {
	s, err := filepath.Abs(common.WorkDir() + "/git")
	utils.Bail("failed to get absolute path", err)
	return s
}

func getworkdirctr() int {
	wd_lock.Lock()
	workdir_ctr++
	w := workdir_ctr
	wd_lock.Unlock()
	return w
}
func recreate() {
	wd_lock.Lock()
	defer wd_lock.Unlock()
	if recreated {
		return
	}
	err := utils.RecreateSafely(workdir())
	utils.Bail("failed to recreate workdir", err)
	recreated = true
}

type LocalRepo struct {
	shallow   bool
	url       string
	fetchurls []*pb.FetchURL
	inuse     bool
	workdir   string // the directory containing "repo"
	stdout    io.Writer
	created   time.Time
	released  time.Time
}

// clone a repo, check it out to current head in master and fetch optional urls too
func GetLocalRepo(ctx context.Context, url string, fetchurls []*pb.FetchURL, shallow bool, stdout io.Writer) (*LocalRepo, error) {
	recreate()
	lr := &LocalRepo{
		url:       url,
		fetchurls: fetchurls,
		inuse:     true,
		stdout:    stdout,
		created:   time.Now(),
		shallow:   shallow,
	}
	repos = append(repos, lr)
	lr.workdir = fmt.Sprintf("%s/%d", workdir(), getworkdirctr())
	err := lr.Clone(ctx)
	if err != nil {
		return nil, fmt.Errorf("git clone failed: %w", err)
	}
	for _, fu := range lr.fetchurls {
		err = lr.Fetch(ctx, fu)
		if err != nil {
			return nil, fmt.Errorf("git fetch failed: %w", err)
		}
	}
	return lr, nil
}
func GetLocalRepos() *pb.LocalRepoList {
	wd_lock.Lock()
	defer wd_lock.Unlock()
	res := &pb.LocalRepoList{}
	for _, r := range repos {
		lr := &pb.LocalRepo{
			URL:       r.url,
			FetchURLs: r.fetchurls,
			InUse:     r.inuse,
			WorkDir:   r.workdir,
			Created:   uint32(r.created.Unix()),
			Released:  uint32(r.released.Unix()),
		}
		res.Repos = append(res.Repos, lr)
	}
	return res
}
func (lr *LocalRepo) Release() {
	lr.inuse = false
	lr.released = time.Now()
}
func (lr *LocalRepo) Printf(txt string, args ...interface{}) {
	s := fmt.Sprintf("[git] ")
	s = fmt.Sprintf(s+txt, args...)
	fmt.Print(s)
	if lr.stdout != nil {
		lr.stdout.Write([]byte(s))
	}
}

// todo: recursive submodules?
func (lr *LocalRepo) Clone(ctx context.Context) error {
	l := linux.NewWithContext(ctx)
	l.SetMaxRuntime(time.Duration(300) * time.Second)
	l.SetEnvironment(GetGitEnv())

	dir := lr.workdir
	os.MkdirAll(dir, 0777)
	lr.Printf("Cloning git repo %s into %s...\n", lr.url, dir)
	var err error
	var out string
	com := []string{"git", "clone"}
	if lr.shallow {
		com = append(com, []string{"--depth", "3"}...)
	}
	if *with_recursive {
		com = append(com, []string{"--recurse-submodules", lr.url, "repo"}...)
	} else {
		com = append(com, []string{lr.url, "repo"}...)
	}
	out, err = l.SafelyExecuteWithDir(com, dir, nil)
	if err != nil {
		lr.Printf("Error (%s). Git-clone %s said: %s\n", err, lr.url, out)
		return err
	}
	lr.Printf("Cloned.\n")
	return nil
}

// todo: recursive submodules?
func (lr *LocalRepo) Fetch(ctx context.Context, fu *pb.FetchURL) error {
	l := linux.NewWithContext(ctx)
	l.SetMaxRuntime(time.Duration(300) * time.Second)
	l.SetEnvironment(GetGitEnv())

	dir := lr.GitRepoPath()
	os.MkdirAll(dir, 0777)
	lr.Printf("Fetching from %s into %s...\n", fu.URL, dir)
	var err error
	var out string
	com := []string{"git", "fetch", fu.URL}
	if fu.RefSpec != "" {
		com = append(com, fu.RefSpec)
	}
	out, err = l.SafelyExecuteWithDir(com, dir, nil)
	if err != nil {
		lr.Printf("Error (%s). Git-fetch %s said: %s\n", err, lr.url, out)
		return err
	}
	lr.Printf("Cloned.\n")
	return nil
}

func (lr *LocalRepo) Checkout(ctx context.Context, commitid string) error {
	l := linux.NewWithContext(ctx)
	l.SetMaxRuntime(time.Duration(300) * time.Second)
	l.SetEnvironment(GetGitEnv())

	lr.Printf("Checking out commit %s\n", commitid)
	dir := lr.GitRepoPath()
	out, err := l.SafelyExecuteWithDir([]string{"git", "checkout", commitid}, dir, nil)
	if err != nil {
		lr.Printf("Error (%s): Git-checkout %s said: %s\n", err, commitid, out)
		return err
	}
	lr.Printf("Checkout completed.\n")
	return nil
}

// returns directory containing ".git"
func (lr *LocalRepo) GitRepoPath() string {
	return lr.workdir + "/repo"
}

// gets the logmessage of the currently checked out commit
func (lr *LocalRepo) GetLogMessage(ctx context.Context) (string, error) {
	l := linux.NewWithContext(ctx)
	l.SetMaxRuntime(time.Duration(300) * time.Second)
	l.SetEnvironment(GetGitEnv())

	gitlog, err := l.SafelyExecuteWithDir([]string{"git", "log", "-1"}, lr.GitRepoPath(), nil)
	if err != nil {
		fmt.Printf("Git said: %s\n", gitlog)
		return "", fmt.Errorf("failed to get git log (%s)", err)
	}

	logmessage := tidyLogMessage(gitlog)
	return logmessage, nil
}
func tidyLogMessage(msg string) string {
	logmessage := ""
	sep := strings.Split(msg, "\n")
	empty_lines := 0
	last_was_empty := false
	for _, l := range sep {
		l = strings.Trim(l, " ")
		l = strings.Trim(l, "\t")
		fmt.Printf("[%d] Line: \"%s\"\n", empty_lines, l)
		if l == "" {
			if last_was_empty {
				continue
			}
			last_was_empty = true
			empty_lines++
			continue
		}
		last_was_empty = false
		if empty_lines != 1 {
			continue
		}
		logmessage = logmessage + l
	}
	if logmessage == "" {
		logmessage = msg
	}
	return logmessage
}
