package builder

import (
	"context"
	"golang.conradwood.net/gitbuilder/coderunners"
	"golang.conradwood.net/go-easyops/utils"
)

// this will take the build rules and build all of it
func (b *Builder) BuildAll(ctx context.Context) error {
	buildrules := b.buildrules
	return b.BuildWithRules(ctx, buildrules)
}
func (b *Builder) BuildWithRules(ctx context.Context, buildrules *BuildRules) error {
	b.Printf("Building (%d rules)...\n", len(buildrules.Builds))
	for _, bds := range buildrules.Builds {
		b.Printf("Build: %s\n", bds)
	}
	target_arch := "amd64"
	target_os := "linux"

	for _, rulename := range buildrules.Builds {
		b.Printf("rule: \"%s\"\n", rulename)
		tagname := buildrules.CheckBuildType(rulename)
		if tagname == "" {
			b.Printf("rule \"%s\" is not valid\n", rulename)
			continue
		}
		// a script is EITHER a coderunner OR a script (coderunner has precedence)
		// coderunners are preferred, scripts will be migrated to coderunners once scripts work well
		for _, scriptname := range BUILD_SCRIPTS[tagname] {
			b.Printf("rule \"%s\" triggers script \"%s\"\n", rulename, scriptname)
			ran, err := coderunners.Run(ctx, b, scriptname)
			if err != nil {
				b.Printf("Coderunner failed: %s\n", utils.ErrorString(err))
				return err
			}
			if ran {
				continue
			}

			bscript := b.findscript(scriptname)
			err = b.buildscript(ctx, bscript, target_arch, target_os)
			if err != nil {
				b.Printf("Script \"%s\" (as \"%s\") failed: %s\n", bscript, scriptname, utils.ErrorString(err))
				return err
			}
		}
	}

	return nil
}
