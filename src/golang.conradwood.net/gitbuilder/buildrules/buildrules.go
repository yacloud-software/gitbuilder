package buildrules

import (
	"fmt"
	pb "golang.conradwood.net/apis/gitbuilder"
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
		"STANDARD_GO":     []string{"coderunner-go-version", "go-build.sh", "go-vet.sh"},
		"GO_VET":          []string{"go-vet.sh"},
		"GO":              []string{"coderunner-go-version", "go-build.sh"},
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
	br      *pb.BuildRules
	Targets []string
	//ExcludeGoDirs []string
}
type Printer interface {
	Printf(format string, args ...interface{})
}

func defaultBuildRules() *BuildRules {
	res := &BuildRules{
		br: &pb.BuildRules{
			PreBuild:   "reject",
			PostCommit: "do",
		},
	}
	res.addBuildTypeFromString("CLEAN")
	res.addBuildTypeFromString("STATICCHECK")
	return res
}

// parses oldstyle
func Read(p Printer, filename string) (*BuildRules, error) {
	b, err := readYaml(p, filename)
	if err == nil {
		return b, nil
	} else {
		fmt.Printf("failed to parse BUILD_RULES as yaml: %s\n", err)
		if must_be_yaml(filename) {
			return nil, err
		}
	}
	b, err = readOldStyle(p, filename)
	return b, fmt.Errorf("invalid 'old-style' BUILD_RULES: %s", err)
}

func must_be_yaml(filename string) bool {
	b, err := utils.ReadFile(filename)
	if err != nil {
		return false
	}
	for _, line := range strings.Split(string(b), "\n") {
		line = strings.Trim(line, " ")
		if len(line) < 3 {
			continue
		}
		if line[0] == '-' {
			return true
		}
		if line == "rules:" {
			return true
		}
	}
	return false
}

// read the new yaml style
func readYaml(p Printer, filename string) (*BuildRules, error) {
	rules := defaultBuildRules()
	r := &pb.BuildRules{}
	err := utils.ReadYaml(filename, r)
	if err != nil {
		return nil, err
	}
	rules.br = r
	return rules, nil
}

// read the old BUILD_RULES text style
func readOldStyle(p Printer, filename string) (*BuildRules, error) {
	rules := defaultBuildRules()
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
			rules.br.PreBuild, err = actionString(sp[1])
		} else if sp[0] == "POSTCOMMIT" {
			rules.br.PostCommit, err = actionString(sp[1])
		} else if sp[0] == "GO_CGO_ENABLED" {
			b, err := parseBoolean(sp[1])
			if err != nil {
				return nil, fmt.Errorf("%s invalid boolean: %s", sp[0], err)
			}
			rules.getGo().CGOEnabled = b
		} else if sp[0] == "GO_EXCLUDE_DIRS" {
			rules.getGo().ExcludeDirs = strings.Split(sp[1], ",")
		} else if sp[0] == "BUILDS" {
			gotBuilds = true
			for _, bs := range strings.Split(sp[1], ",") {
				b := strings.Trim(bs, " ")
				err = rules.addBuildTypeFromString(b)
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
	rules.addBuildTypeFromString("DIST")
	return rules, nil
}
func defaultBuildOS() string {
	return "debian-12.4"
}
func (br *BuildRules) addBuildTypeFromString(name string) error {
	if br.findType(name) != nil {
		return nil
	}
	bdr := &pb.BuildRuleDef{BuildType: name, BuildOS: defaultBuildOS()}
	if name == "STANDARD_GO" {
		bdr.Go = &pb.BuildRuleDef_Go{}
	}
	br.br.Rules = append(br.br.Rules, bdr)
	return nil
}

func (br *BuildRules) findType(name string) *pb.BuildRuleDef {
	for _, bd := range br.br.Rules {
		if bd.BuildType == name {
			return bd
		}
	}
	return nil
}

func actionString(s string) (string, error) {
	_, e := parseAction(s)
	if e != nil {
		return "", e
	}
	return s, nil
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

// get the first go build rule, create if necessary
func (br *BuildRules) getGo() *pb.BuildRuleDef_Go {
	for _, bd := range br.br.Rules {
		if bd.Go != nil {
			return bd.Go
		}
	}
	res := &pb.BuildRuleDef_Go{}
	rule := &pb.BuildRuleDef{
		BuildType: "STANDARD_GO",
		Go:        res,
	}
	br.br.Rules = append(br.br.Rules, rule)
	return res
}

func (br *BuildRules) Go_CGO_EnabledAsEnv() string {
	if br.getGo().CGOEnabled {
		return "1"
	}
	return "0"
}
func (br *BuildRules) Go_ExcludeDirsAsEnv() string {
	if br.br == nil {
		return ""
	}
	var excludes []string
	for _, rule := range br.br.Rules {
		if rule.Go == nil {
			continue
		}
		excludes = append(excludes, rule.Go.ExcludeDirs...)
	}
	if len(excludes) == 0 {
		return ""
	}
	s := strings.Join(excludes, " ")
	return s
}
func (br *BuildRules) Proto() *pb.BuildRules {
	return br.br
}
func (br *BuildRules) BuildDefs() []*pb.BuildRuleDef {
	return br.br.Rules
}
