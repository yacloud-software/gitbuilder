package builder

import (
	"context"
)

func (b *Builder) BuildAll(ctx context.Context) error {
	b.Printf("Building (%d rules)...\n", len(b.buildrules.Builds))
	for _, bds := range b.buildrules.Builds {
		b.Printf("Build: %s\n", bds)
	}
	target_arch := "amd64"
	target_os := "linux"
	err := b.buildscript(ctx, b.findscript("clean-build.sh"), target_arch, target_os)
	if err != nil {
		return err
	}
	for rulename, _ := range BUILD_SCRIPTS {
		scriptname := b.buildrules.CheckBuildType(rulename)
		if scriptname == "" {
			continue
		}
		err = b.buildscript(ctx, b.findscript(scriptname), target_arch, target_os)
		if err != nil {
			return err
		}
	}

	// ** now create the dist and upload it
	err = b.buildscript(ctx, b.findscript("dist.sh"), target_arch, target_os)
	if err != nil {
		return err
	}
	return nil
}
