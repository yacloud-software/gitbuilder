package builder

import (
	"context"
	"fmt"
	am "golang.conradwood.net/apis/auth"
	"golang.conradwood.net/go-easyops/auth"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/cmdline"
	"golang.conradwood.net/go-easyops/utils"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func (b *Builder) findscript(scriptname string) string {
	cwd, err := os.Getwd()
	if err != nil {
		b.Printf("Unable to find current working directory: %s\n", err)
		panic(fmt.Sprintf("Unable to find current working directory: %s\n", err))
	}
	pwd, err := filepath.Abs(cwd)
	if err != nil {
		b.Printf("Unable to find absolute current working directory: %s\n", err)
		panic(fmt.Sprintf("Unable to find absolute current working directory: %s\n", err))
	}
	if utils.FileExists("/tmp/build_scripts/" + scriptname) {
		return "/tmp/build_scripts/" + scriptname
	}
	for {
		if utils.FileExists(pwd + "/scripts/" + scriptname) {
			return pwd + "/scripts/" + scriptname
		}
		pwd = filepath.Dir(pwd)
		if (pwd == "/") || (pwd == "") {
			break
		}
	}
	b.Printf("Could not find script \"%s\"\n", scriptname)
	return ""

}
func (b *Builder) buildscript(ctx context.Context, fscript, target_arch, target_os string) error {
	cmd := exec.Command(fscript)
	cmd.Dir = b.GetRepoPath()
	fmt.Printf("Executing script %s in cwd \"%s\"\n", fscript, cmd.Dir)
	if cmd.Dir == "" || !utils.FileExists(cmd.Dir) {
		return fmt.Errorf("Directory \"%s\" does not exist\n", cmd.Dir)
	}
	ep, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	op, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	fn := filepath.Base(fscript)
	go b.pipeOutput(fn, ep)
	go b.pipeOutput(fn, op)
	cmd.Env = b.env()
	b.addContextEnv(ctx, cmd)
	cmd.Env = append(cmd.Env, fmt.Sprintf("TARGET_ARCH=%s", target_arch))
	cmd.Env = append(cmd.Env, fmt.Sprintf("TARGET_OS=%s", target_os))
	err = cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Wait()
	if err != nil {
		return err
	}
	return nil
}
func (b *Builder) pipeOutput(scriptname string, rc io.ReadCloser) {
	buf := make([]byte, 1024)
	for {
		n, err := rc.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Failed to read %s:\n", err)
			break
		}
		b.Printf("%s: %s", scriptname, string(buf[:n]))
	}

}

func (b *Builder) env() []string {
	// standard environment variables...
	std := `
JAVA_HOME=/etc/java-home
GRADLE_HOME=/srv/java/gradle/latest
TERM=xterm
SHELL=/bin/bash
ANT_HOME=/srv/java/ant/current/
PATH=/opt/yacloud/ctools/dev/bin:/opt/yacloud/ctools/dev/go/current/go/bin/:/etc/java-home/bin:/srv/singingcat/binutils/bin/:~/bin:/sbin:/usr/sbin:/usr/local/bin:/usr/bin:/bin:/srv/java/ant/current/bin:/srv/singingcat/esp8266/sdk/xtensa-lx106-elf/bin/:/srv/java/ant/bin:/srv/java/gradle/latest/bin
PWD=/tmp
GOROOT=/opt/yacloud/ctools/dev/go/current/go
LANG=en_GB.UTF-8
LANGUAGE=en_GB:en
LC_CTYPE=en_GB.UTF-8
`
	var res []string
	for _, s := range strings.Split(std, "\n") {
		if len(s) < 2 {
			continue
		}
		res = append(res, s)
	}
	dir, err := os.Getwd()
	bindir := "./"
	if err != nil {
		fmt.Printf("Unable to get current directory. (%s)\n", err)
	} else {
		bindir = dir
	}
	absdir := b.GetRepoPath()

	fmt.Printf("Bindir: \"%s\"\n", bindir)
	os.MkdirAll(bindir+"/gobin", 0777)
	os.MkdirAll(bindir+"/gocache", 0777)
	os.MkdirAll(bindir+"/gotmp", 0777)
	res = append(res, fmt.Sprintf("BUILD_NUMBER=%d", b.buildid))
	res = append(res, fmt.Sprintf("GOPATH=%s", absdir))
	res = append(res, fmt.Sprintf("HOME=%s", absdir))
	res = append(res, fmt.Sprintf("BUILD_DIR=%s", absdir))
	res = append(res, fmt.Sprintf("COMMIT_ID=%s", b.buildinfo.CommitID()))
	res = append(res, fmt.Sprintf("REPOSITORY_ID=%d", b.buildinfo.RepositoryID()))
	res = append(res, fmt.Sprintf("PROJECT_NAME=%s", b.buildinfo.RepositoryName()))
	res = append(res, fmt.Sprintf("BUILD_REPOSITORY=%s", b.buildinfo.RepositoryName()))
	res = append(res, fmt.Sprintf("BUILD_ARTEFACT=%s", b.buildinfo.RepositoryArtefactName()))
	res = append(res, fmt.Sprintf("BUILD_TIMESTAMP=%d", b.timestamp.Unix()))
	res = append(res, fmt.Sprintf("GIT_BRANCH=%s", "master"))
	res = append(res, fmt.Sprintf("GOBIN=%s/gobin", bindir))
	res = append(res, fmt.Sprintf("GOCACHE=%s/gocache", bindir))
	res = append(res, fmt.Sprintf("GOTMPDIR=%s/gotmp", bindir))
	res = append(res, fmt.Sprintf("REGISTRY=%s", cmdline.GetClientRegistryAddress()))
	//	res = append(res, fmt.Sprintf("SCRIPTDIR=%s", scriptsdir))

	return res
}
func (b *Builder) addContextEnv(ctx context.Context, cmd *exec.Cmd) error {
	u := auth.GetUser(ctx)
	if u == nil {
		fmt.Printf("WARNING: no user in context\n")
	} else {
		fmt.Printf("Executing scripts as user %s (%s)\n", u.ID, auth.Description(u))
	}
	ncb, err := auth.SerialiseContextToString(ctx)
	if err != nil {
		fmt.Printf("Failed to encode context to string: %s\n", err)
		return err
	}

	ncs := fmt.Sprintf("GE_CTX=%s", ncb)
	for i, e := range cmd.Env {
		if strings.HasPrefix(e, "GE_CTX=") {
			cmd.Env[i] = ncs
			return nil
		}
	}
	cmd.Env = append(cmd.Env, ncs)
	if u != nil {
		cmd.Env = append(cmd.Env, fmt.Sprintf("GE_USER_EMAIL=%s", u.Email))
		cmd.Env = append(cmd.Env, fmt.Sprintf("GE_USER_ID=%s", u.ID))
		/*
			tr, err := GetAuthManagerClient().GetTokenForMe(ctx, &am.GetTokenRequest{DurationSecs: 300})
			if err != nil {
				fmt.Printf("unable to get authentication token for external script(s): %s\n", utils.ErrorString(err))
				return err
			}
		*/

	}
	return nil
}
func GetAuthManagerClient() am.AuthManagerServiceClient {
	return authremote.GetAuthManagerClient()
}
