package builder

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	am "golang.conradwood.net/apis/auth"
	"golang.conradwood.net/gitbuilder/common"
	"golang.conradwood.net/go-easyops/auth"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/cmdline"
	"golang.conradwood.net/go-easyops/utils"
)

var (
	locpath string

	PATH = []string{
		"/opt/yacloud/current/ctools/dev/bin", "/opt/yacloud/current/ctools/dev/go/current/go/bin/",
		"/opt/yacloud/ctools/dev/bin", "/opt/yacloud/ctools/dev/go/current/go/bin/",
		"/etc/java-home/bin", "/srv/singingcat/binutils/bin/", "~/bin", "/sbin", "/usr/sbin", "/usr/local/bin", "/usr/bin", "/bin", "/srv/java/ant/current/bin", "/srv/singingcat/esp8266/sdk/xtensa-lx106-elf/bin/", "/srv/java/ant/bin", "/srv/java/gradle/latest/bin"}
)

// given a scriptname, e.g. "autobuild.sh" or "go-build.sh" tries to find the script.
// autobuild.sh is the only which will be searched for in the working directory (legacy requirement)
func (b *Builder) findscript(scriptname string) string {
	if scriptname == "autobuild.sh" {
		res := b.GetRepoPath() + "/autobuild.sh"
		if !utils.FileExists(res) {
			b.Printf("autobuild.sh configured but not found\n")
			return ""
		}
		return res
	}
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
	f, err := utils.FindFile("scripts/" + scriptname)
	if err == nil {
		pwd, err := filepath.Abs(f)
		if err != nil {
			fmt.Printf("WARNING: can't turn \"%s\" into absolute path: %s\n", f, err)
			return f
		}
		return pwd
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
	if !utils.FileExists(fscript) {
		return fmt.Errorf("file %s does not exist", fscript)
	}
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
	cmd.Env = b.env(ctx)
	b.addContextEnv(ctx, cmd)
	cmd.Env = append(cmd.Env, fmt.Sprintf("TARGET_ARCH=%s", target_arch))
	cmd.Env = append(cmd.Env, fmt.Sprintf("TARGET_OS=%s", target_os))
	err = cmd.Start()
	if err != nil {
		return err
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("Error executing script \"%s\": %s\n", fscript, err)
		return err
	}
	b.Printf("Script \"%s\" completed successfully\n", fscript)
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

func (b *Builder) env(ctx context.Context) []string {
	u := auth.GetUser(ctx)
	if u == nil {
		fmt.Printf("executing without a user account!\n")
		return nil
	}
	// standard environment variables...
	std := `
JAVA_HOME=/etc/java-home
GRADLE_HOME=/srv/java/gradle/latest
TERM=xterm
SHELL=/bin/bash
ANT_HOME=/srv/java/ant/current/
PWD=/tmp
GOROOT=` + cmdline.GetYACloudDir() + `/ctools/dev/go/current/go
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

	sp := strings.Join(PATH, ":")

	fmt.Printf("Bindir: \"%s\"\n", bindir)
	os.MkdirAll(bindir+"/gobin", 0777)
	res = append(res, fmt.Sprintf("CGO_ENABLED=%s", b.buildrules.Go_CGO_EnabledAsEnv()))
	res = append(res, fmt.Sprintf("EXCLUDE_GO_DIRS=%s", b.buildrules.Go_ExcludeDirsAsEnv()))
	res = append(res, fmt.Sprintf("PATH=%s", sp))
	res = append(res, fmt.Sprintf("GIT_URL=%s", b.buildinfo.GitURL()))
	res = append(res, fmt.Sprintf("BUILD_NUMBER=%d", b.buildid))
	res = append(res, fmt.Sprintf("GOPATH=%s", absdir))
	res = append(res, fmt.Sprintf("HOME=%s", absdir))
	res = append(res, fmt.Sprintf("BUILD_DIR=%s", absdir))
	res = append(res, fmt.Sprintf("COMMIT_ID=%s", b.buildinfo.CommitID()))
	res = append(res, fmt.Sprintf("REPOSITORY_ID=%d", b.buildinfo.RepositoryID()))
	res = append(res, fmt.Sprintf("PROJECT_NAME=%s", b.buildinfo.RepositoryName()))
	res = append(res, fmt.Sprintf("BUILD_REPOSITORY=%s", b.buildinfo.RepositoryName()))
	res = append(res, fmt.Sprintf("BUILD_ARTEFACT=%s", b.buildinfo.RepositoryArtefactName()))
	res = append(res, fmt.Sprintf("BUILD_ARTEFACTID=%d", b.buildinfo.ArtefactID()))
	res = append(res, fmt.Sprintf("BUILD_TIMESTAMP=%d", b.timestamp.Unix()))
	res = append(res, fmt.Sprintf("GIT_BRANCH=%s", "master"))
	res = append(res, fmt.Sprintf("GOBIN=%s/gobin", bindir))
	res = append(res, fmt.Sprintf("GOMAXPROCS=%d", common.GetMaxProcs()))
	res = append(res, "GOSUMDB=off")
	res = append(res, fmt.Sprintf("REGISTRY=%s", cmdline.GetClientRegistryAddress()))
	//	res = append(res, fmt.Sprintf("SCRIPTDIR=%s", scriptsdir))
	if common.GetGoCache() == "" {
		res = append(res, fmt.Sprintf("GOTMPDIR=%s/gotmp", bindir))
		res = append(res, fmt.Sprintf("GOCACHE=%s/gocache", bindir))
		os.MkdirAll(bindir+"/gocache", 0777)
		os.MkdirAll(bindir+"/gotmp", 0777)
	} else {
		gc, err := filepath.Abs(common.GetGoCache())
		gc = gc + fmt.Sprintf("/%s/", u.ID) // must have seperate caches per user, so we force download of go modules per user
		utils.Bail("failed to absolute gocache", err)
		res = append(res, fmt.Sprintf("GOCACHE=%s/gocache", gc))
		res = append(res, fmt.Sprintf("GOMODCACHE=%s/gomodcache", gc))
		res = append(res, fmt.Sprintf("GOTMP=%s/gotmp", gc))
		res = append(res, fmt.Sprintf("GOTMPDIR=%s/gotmp", gc))
		os.MkdirAll(fmt.Sprintf("%s/gocache", gc), 0777)
		os.MkdirAll(fmt.Sprintf("%s/gomodcache", gc), 0777)
		os.MkdirAll(fmt.Sprintf("%s/gotmp", gc), 0777)
	}

	// make LDFLAGS="-ldflags '-X golang.conradwood.net/go-easyops/appinfo.LD_Number=56'"
	// make LDFLAGS="-ldflags '-X golang.conradwood.net/go-easyops/appinfo.LD_Number=56 -X golang.conradwood.net/go-easyops/appinfo.LD_Timestamp=89'"
	ldflags := `-ldflags '-X golang.conradwood.net/go-easyops/appinfo.LD_Number=%d -X golang.conradwood.net/go-easyops/appinfo.LD_Description=%s -X golang.conradwood.net/go-easyops/appinfo.LD_Timestamp=%d -X golang.conradwood.net/go-easyops/appinfo.LD_ArtefactID=%d -X golang.conradwood.net/go-easyops/appinfo.LD_RepositoryID=%d -X golang.conradwood.net/go-easyops/appinfo.LD_RepositoryName=%s -X golang.conradwood.net/go-easyops/appinfo.LD_CommitID=%s -X golang.conradwood.net/go-easyops/appinfo.LD_GitURL=%s' `
	ldflags = fmt.Sprintf(ldflags, b.buildid, "gitbuilder", b.timestamp.Unix(), b.buildinfo.ArtefactID(), b.buildinfo.RepositoryID(), b.buildinfo.RepositoryName(), b.buildinfo.CommitID(), b.buildinfo.GitURL())
	res = append(res, fmt.Sprintf("GO_LDFLAGS=%s", ldflags))
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

	cmd.Env = append(cmd.Env, fmt.Sprintf("GE_USER_EMAIL=%s", u.Email))
	cmd.Env = append(cmd.Env, fmt.Sprintf("GE_USER_ID=%s", u.ID))
	tr, err := GetAuthManagerClient().GetTokenForMe(ctx, &am.GetTokenRequest{DurationSecs: 300})
	if err != nil {
		fmt.Printf("unable to get authentication token for external script(s): %s\n", utils.ErrorString(err))
		fmt.Printf("Context:%#v\n", ctx)
		return err
	}
	cmd.Env = append(cmd.Env, fmt.Sprintf("PROXY_USER=%s@token.yacloud.eu", u.ID))
	cmd.Env = append(cmd.Env, fmt.Sprintf("PROXY_PASSWORD=%s", tr.Token))
	if common.GetGoProxyHost() != "" {
		s := ""
		if common.GetGoProxyDirect() {
			s = ",direct"
		}
		cmd.Env = append(cmd.Env, fmt.Sprintf("GOPROXY=https://%s@token.yacloud.eu:%s@%s%s", u.ID, tr.Token, common.GetGoProxyHost(), s))
	}
	return nil
}
func GetAuthManagerClient() am.AuthManagerServiceClient {
	return authremote.GetAuthManagerClient()
}
