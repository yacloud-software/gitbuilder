package builder

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

type BuildRules struct {
	Prebuild   int
	PostCommit int
	Builds     []string
	Targets    []string
}

func (b *Builder) readBuildrules() error {
	rules := &BuildRules{}
	b.buildrules = rules
	br := b.GetRepoPath() + "/BUILD_RULES"
	if !utils.FileExists(br) {
		b.Printf("rules file (%s) does not exist. no builds.\n", br)
		return nil
	}

	fc, err := ioutil.ReadFile(br)
	if err != nil {
		return err
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
			return fmt.Errorf("buildrules: Line %d invalid (only %d parts) [%s]", ln+1, len(sp), line)
		}
		if sp[0] == "PREBUILD" {
			rules.Prebuild, err = parseAction(sp[1])
		} else if sp[0] == "POSTCOMMIT" {
			rules.PostCommit, err = parseAction(sp[1])
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
			return fmt.Errorf("buildrules: Line %d invalid (invalid tag \"%s\") [%s]", ln+1, sp[0], line)
		}
		if err != nil {
			return err
		}
	}

	// set default to autobuild.sh if it exists
	if !gotBuilds {
		if utils.FileExists(b.GetRepoPath() + "/autobuild.sh") {
			rules.Builds = []string{"AUTOBUILD_SH"}
		}
	}

	// add defaults:
	b.buildrules.Builds = append([]string{"CLEAN"}, b.buildrules.Builds...)
	b.buildrules.Builds = append([]string{"STATICCHECK"}, b.buildrules.Builds...)
	b.buildrules.Builds = append(b.buildrules.Builds, "DIST")
	return nil
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

// return tag or ""
func (b *BuildRules) CheckBuildType(buildtype string) string {
	_, valid := BUILD_SCRIPTS[buildtype]
	if valid {
		return buildtype
	}
	return ""
}
