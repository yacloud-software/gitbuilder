package git

import (
	"flag"
	"fmt"
	"golang.conradwood.net/go-easyops/cmdline"
	"golang.conradwood.net/go-easyops/utils"
	"os"
	"path/filepath"
)

const (
	GITCONFIG = `
[credential]
        helper="%s"
        useHttpPath = true
[pull]
        rebase = false
[safe]
        directory = *
`
)

var (
	locpath         string
	local_gitconfig = flag.Bool("use_local_gitconfig", false, "if true use a local gitconfig, contained within gitbuilder")
)

func init() {
	l, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("os.Getwd: %s", err))
	}
	locpath = l
}
func getGitEnv() []string {
	if !*local_gitconfig {
		return nil
	}

	bin_name, err := utils.FindFile("extra/gitcredentials-client")
	utils.Bail("Unable to find gitcredentials-client: %s\n", err)
	bin_name, err = filepath.Abs(bin_name)
	utils.Bail("unable to make path absolute", err)
	x := fmt.Sprintf("%s -registry=%s", bin_name, cmdline.GetRegistryAddress())
	gitconf := fmt.Sprintf(GITCONFIG, x)
	fname := locpath + "/git.config"
	err = utils.WriteFile(fname, []byte(gitconf))
	utils.Bail("failed to write gitconf", err)
	return []string{"GIT_CONFIG=" + fname}
}
