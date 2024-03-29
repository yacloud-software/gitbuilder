package builder

import (
	"context"
	"fmt"
	"golang.conradwood.net/gitbuilder/buildrules"
	"golang.conradwood.net/gitbuilder/coderunners"
	"golang.conradwood.net/go-easyops/utils"
)

// this will take the build rules and build all of it
func (b *Builder) BuildAll(ctx context.Context) error {
	buildrules := b.buildrules
	return b.BuildWithRules(ctx, buildrules)
}
func (b *Builder) BuildWithRules(ctx context.Context, br *buildrules.BuildRules) error {
	b.Printf("Building #%d, RepositoryID %d, Artefact %d\n", b.buildinfo.BuildNumber(), b.buildinfo.RepositoryID(), b.buildinfo.ArtefactID())
	b.Printf("Building (%d rules)...\n", len(br.BuildDefs()))
	for _, bds := range br.BuildDefs() {
		b.Printf("Build: %s\n", bds.BuildType)
	}
	target_arch := "amd64"
	target_os := "linux"

	for _, bdef := range br.BuildDefs() {
		rulename := bdef.BuildType
		b.Printf("rule: \"%s\"\n", rulename)
		if !b.BuildInfo().IsScriptIncluded(rulename) {
			b.Printf("Rule skipped (it is explicitly excluded in buildrequest)\n")
			continue
		}

		ran, err := coderunners.Run(ctx, b, rulename)
		if err != nil {
			b.Printf("Coderunner %s failed: %s\n", rulename, utils.ErrorString(err))
			return err
		}
		if ran {
			continue
		}

		tagname := br.CheckBuildType(rulename)
		if tagname == "" {
			b.Printf("rule \"%s\" is not valid\n", rulename)
			continue
		}
		// a script is EITHER a coderunner OR a script (coderunner has precedence)
		// coderunners are preferred, scripts will be migrated to coderunners once scripts work well
		for _, scriptname := range buildrules.BUILD_SCRIPTS[tagname] {
			ran, err = coderunners.Run(ctx, b, scriptname)
			if err != nil {
				b.Printf("Coderunner failed: %s\n", utils.ErrorString(err))
				return err
			}
			if ran {
				continue
			}
			b.Printf("rule \"%s\" triggers script \"%s\"\n", rulename, scriptname)

			bscript := b.findscript(scriptname)
			fmt.Printf("Found script \"%s\" here: %s\n", scriptname, bscript)
			err = b.buildscript(ctx, bscript, target_arch, target_os)
			if err != nil {
				b.Printf("Script \"%s\" (as \"%s\") failed: %s\n", bscript, scriptname, utils.ErrorString(err))
				return err
			}
		}
	}

	return nil
}
