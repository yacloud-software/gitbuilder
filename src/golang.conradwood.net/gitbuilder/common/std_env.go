package common

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.conradwood.net/gitbuilder/buildinfo"

	"golang.conradwood.net/go-easyops/auth"
	"golang.conradwood.net/go-easyops/cmdline"
	"golang.conradwood.net/go-easyops/utils"
)

var (
	PATH = []string{
		"/opt/yacloud/current/ctools/dev/bin", "/opt/yacloud/current/ctools/dev/go/current/go/bin/",
		"/opt/yacloud/ctools/dev/bin", "/opt/yacloud/ctools/dev/go/current/go/bin/",
		"/etc/java-home/bin", "/srv/singingcat/binutils/bin/", "~/bin", "/sbin", "/usr/sbin", "/usr/local/bin", "/usr/bin", "/bin", "/srv/java/ant/current/bin", "/srv/singingcat/esp8266/sdk/xtensa-lx106-elf/bin/", "/srv/java/ant/bin", "/srv/java/gradle/latest/bin"}
)

type EnvDev interface {
	GetRepoPath() string
	BuildInfo() buildinfo.BuildInfo
}

func StdEnv(ctx context.Context, def EnvDev) []string {
	bi := def.BuildInfo()
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
	absdir := def.GetRepoPath()

	sp := strings.Join(PATH, ":")

	fmt.Printf("Bindir: \"%s\"\n", bindir)
	os.MkdirAll(bindir+"/gobin", 0777)
	res = append(res, fmt.Sprintf("PATH=%s", sp))
	res = append(res, fmt.Sprintf("GIT_URL=%s", bi.GitURL()))
	res = append(res, fmt.Sprintf("BUILD_NUMBER=%d", bi.BuildNumber()))
	res = append(res, fmt.Sprintf("GOPATH=%s", absdir))
	res = append(res, fmt.Sprintf("HOME=%s", absdir))
	res = append(res, fmt.Sprintf("BUILD_DIR=%s", absdir))
	res = append(res, fmt.Sprintf("COMMIT_ID=%s", bi.CommitID()))
	res = append(res, fmt.Sprintf("REPOSITORY_ID=%d", bi.RepositoryID()))
	res = append(res, fmt.Sprintf("PROJECT_NAME=%s", bi.RepositoryName()))
	res = append(res, fmt.Sprintf("BUILD_REPOSITORY=%s", bi.RepositoryName()))
	res = append(res, fmt.Sprintf("BUILD_ARTEFACT=%s", bi.RepositoryArtefactName()))
	res = append(res, fmt.Sprintf("BUILD_ARTEFACTID=%d", bi.ArtefactID()))
	//	res = append(res, fmt.Sprintf("BUILD_TIMESTAMP=%d", b.timestamp.Unix()))
	res = append(res, fmt.Sprintf("GIT_BRANCH=%s", "master"))
	res = append(res, fmt.Sprintf("GOBIN=%s/gobin", bindir))
	res = append(res, fmt.Sprintf("GOMAXPROCS=%d", GetMaxProcs()))
	res = append(res, "GOSUMDB=off")
	res = append(res, fmt.Sprintf("REGISTRY=%s", cmdline.GetClientRegistryAddress()))
	//	res = append(res, fmt.Sprintf("SCRIPTDIR=%s", scriptsdir))
	if GetGoCache() == "" {
		res = append(res, fmt.Sprintf("GOTMPDIR=%s/gotmp", bindir))
		res = append(res, fmt.Sprintf("GOCACHE=%s/gocache", bindir))
		os.MkdirAll(bindir+"/gocache", 0777)
		os.MkdirAll(bindir+"/gotmp", 0777)
	} else {
		gc, err := filepath.Abs(GetGoCache())
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
	bts := time.Now()
	ldflags = fmt.Sprintf(ldflags, bi.BuildNumber(), "gitbuilder", bts.Unix(), bi.ArtefactID(), bi.RepositoryID(), bi.RepositoryName(), bi.CommitID(), bi.GitURL())
	res = append(res, fmt.Sprintf("GO_LDFLAGS=%s", ldflags))
	return res

}
