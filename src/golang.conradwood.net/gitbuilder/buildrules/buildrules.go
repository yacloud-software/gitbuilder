package buildrules

import (
	"fmt"
	"golang.conradwood.net/go-easyops/utils"
	"io/ioutil"
	"strings"
)

const (
	RULES_REJECT = 1 // reject if broken build
	RULES_DO     = 2 // warn only
)

var (
	// either name of scripts or coderunners. order of the array matters
	BUILD_SCRIPTS = map[string][]string{
		"STANDARD_PROTOS": []string{"protos-build.sh"},
		"STANDARD_GO":     []string{"coderunner-go-version", "go-build.sh"},
		"KICAD":           []string{"kicad-build.sh"},
		"STANDARD_JAVA":   []string{"java-build.sh"},
		"AUTOBUILD_SH":    []string{"autobuild.sh"},
		"CLEAN":           []string{"clean-build.sh"},
		"DIST":            []string{"dist.sh"},
		"GO_VERSION":      []string{"coderunner-go-version"},
		"GO_MODULES":      []string{"coderunner-gomodule"},
	}
)

type BuildRules struct {
	Prebuild      int
	PostCommit    int
	Builds        []string
	Targets       []string
	ExcludeGoDirs []string
	GoCGOEnabled  bool
}
type Printer interface {
	Printf(format string, args ...interface{})
}

func defaultBuildRules() *BuildRules {
	return &BuildRules{
		Prebuild:   RULES_REJECT,
		PostCommit: RULES_DO,
	}

}
func Read(p Printer, filename string) (*BuildRules, error) {
	rules := &BuildRules{}
	br := filename //b.GetRepoPath() + "/BUILD_RULES"
	if !utils.FileExists(br) {
		p.Printf("rules file (%s) does not exist. no builds.\n", br)
		return defaultBuildRules(), nil
	}
	fc, err := ioutil.ReadFile(br)
	if err != nil {
		return nil, err
	}
	gotBuilds := false
	for ln, line := range strings.Split(string(fc), "\n") {
		if len(line) < 2 {
			continue
		}
		if strings.HasPrefix(line, "#") {
			continue
		}
		sp := strings.SplitN(line, "=", 2)
		if len(sp) < 1 {
			return nil, fmt.Errorf("buildrules: Line %d invalid (only %d parts) [%s]", ln+1, len(sp), line)
		}
		if sp[0] == "PREBUILD" {
			rules.Prebuild, err = parseAction(sp[1])
		} else if sp[0] == "POSTCOMMIT" {
			rules.PostCommit, err = parseAction(sp[1])
		} else if sp[0] == "GO_CGO_ENABLED" {
			b, err := parseBoolean(sp[1])
			if err != nil {
				return nil, fmt.Errorf("%s invalid boolean: %s", sp[0], err)
			}
			rules.GoCGOEnabled = b
		} else if sp[0] == "GO_EXCLUDE_DIRS" {
			rules.ExcludeGoDirs = strings.Split(sp[1], ",")
		} else if sp[0] == "BUILDS" {
			gotBuilds = true
			for _, bs := range strings.Split(sp[1], ",") {
				rules.Builds = append(rules.Builds, strings.Trim(bs, " "))
			}
		} else if sp[0] == "TARGETS" {
			for _, bs := range strings.Split(sp[1], ",") {
				rules.Targets = append(rules.Targets, bs)
			}
		} else {
			return nil, fmt.Errorf("buildrules: Line %d invalid (invalid tag \"%s\") [%s]", ln+1, sp[0], line)
		}
		if err != nil {
			return nil, err
		}
	}

	// set default to autobuild.sh if it exists
	if !gotBuilds {
		p.Printf("no builds defined!\n")
		/*
			if utils.FileExists(b.GetRepoPath() + "/autobuild.sh") {
				rules.Builds = []string{"AUTOBUILD_SH"}
			}
		*/
	}
	// add defaults:
	rules.Builds = append([]string{"CLEAN"}, rules.Builds...)
	rules.Builds = append([]string{"STATICCHECK"}, rules.Builds...)
	rules.Builds = append(rules.Builds, "DIST")
	return rules, nil
}

func parseAction(s string) (int, error) {
	if s == "reject" {
		return RULES_REJECT, nil
	} else if s == "do" {
		return RULES_DO, nil
	} else {
		return -1, fmt.Errorf("buildrules: Tag %s unknown", s)
	}
}
func parseBoolean(s string) (bool, error) {
	s = strings.ToLower(s)
	if s == "true" || s == "yes" || s == "on" {
		return true, nil
	}
	if s == "false" || s == "no" || s == "off" {
		return false, nil
	}
	return false, fmt.Errorf("\"%s\" is not valid for booleans", s)
}

// return tag or ""
func (b *BuildRules) CheckBuildType(buildtype string) string {
	_, valid := BUILD_SCRIPTS[buildtype]
	if valid {
		return buildtype
	}
	return ""
}

func (br *BuildRules) Go_CGO_EnabledAsEnv() string {
	if br.GoCGOEnabled {
		return "1"
	}
	return "0"
}
func (br *BuildRules) Go_ExcludeDirsAsEnv() string {
	if len(br.ExcludeGoDirs) == 0 {
		return ""
	}
	s := strings.Join(br.ExcludeGoDirs, " ")
	return s
}
