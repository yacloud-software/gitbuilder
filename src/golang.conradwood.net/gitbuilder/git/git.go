package git

import (
	"context"
	"flag"
	"fmt"
	pb "golang.conradwood.net/apis/gitbuilder"
	"golang.conradwood.net/go-easyops/linux"
	"golang.conradwood.net/go-easyops/utils"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	repos          []*LocalRepo
	with_recursive = flag.Bool("git_with_recursive", false, "if true use git-clone --recursive")
	wd_lock        sync.Mutex
	workdir_ctr    = 0
	f_workdir      = flag.String("workdir", "/tmp/gitbuilder", "workdir for repos")
	recreated      = false
)

func init() {
	go cleaner()
}
func workdir() string {
	s, err := filepath.Abs(*f_workdir)
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
	url       string
	fetchurls []string
	inuse     bool
	workdir   string // the directory containing "repo"
	stdout    io.Writer
	created   time.Time
	released  time.Time
}

// clone a repo, check it out to current head in master and fetch optional urls too
func GetLocalRepo(ctx context.Context, url string, fetchurls []string, stdout io.Writer) (*LocalRepo, error) {
	recreate()
	lr := &LocalRepo{
		url:       url,
		fetchurls: fetchurls,
		inuse:     true,
		stdout:    stdout,
		created:   time.Now(),
	}
	if len(lr.fetchurls) != 0 {
		panic("cannot do fetch yet")
	}
	repos = append(repos, lr)
	lr.workdir = fmt.Sprintf("%s/%d", workdir(), getworkdirctr())
	err := lr.Clone(ctx)
	if err != nil {
		return nil, err
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
	l.SetRuntime(300)

	dir := lr.workdir
	os.MkdirAll(dir, 0777)
	lr.Printf("Cloning git repo %s into %s...\n", lr.url, dir)
	var err error
	var out string
	if *with_recursive {
		out, err = l.SafelyExecuteWithDir([]string{"git", "clone", "--depth", "3", "--recurse-submodules", lr.url, "repo"}, dir, nil)
	} else {
		out, err = l.SafelyExecuteWithDir([]string{"git", "clone", "--depth", "3", lr.url, "repo"}, dir, nil)
	}
	if err != nil {
		lr.Printf("Error (%s). Git-clone %s said: %s\n", err, lr.url, out)
		return err
	}
	lr.Printf("Cloned.\n")
	return nil
}

func (lr *LocalRepo) Checkout(ctx context.Context, commitid string) error {
	l := linux.NewWithContext(ctx)
	l.SetRuntime(300)

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
	l.SetRuntime(300)
	gitlog, err := l.SafelyExecuteWithDir([]string{"git", "log", "-1"}, lr.GitRepoPath(), nil)
	if err != nil {
		fmt.Printf("Git said: %s\n", gitlog)
		return "", fmt.Errorf("failed to get git log (%s)", err)
	}

	logmessage := ""
	sep := strings.Split(gitlog, "\n")
	for i, l := range sep {
		if l == "" {
			logmessage = strings.Join(sep[i+1:], "\n")
			logmessage = strings.TrimSuffix(logmessage, "\n")
			logmessage = strings.TrimSpace(logmessage)

		}
		//fmt.Printf("%d. \"%s\"\n", i, l)
	}
	return logmessage, nil
}
