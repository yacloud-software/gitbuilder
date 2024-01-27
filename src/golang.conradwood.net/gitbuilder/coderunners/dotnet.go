package coderunners

import (
	"context"
	"fmt"
	"golang.conradwood.net/go-easyops/linux"
	"golang.conradwood.net/go-easyops/utils"
	"path/filepath"
	"time"
)

type dotnet struct {
	builder       brunner
	ctx           context.Context
	project_files []string // relative filenames to .csproj files
}

func (dn *dotnet) Run(ctx context.Context, builder brunner) error {
	dn.builder = builder
	dn.ctx = ctx
	topdir := dn.builder.GetRepoPath() + "/dotnet"
	topdir, err := filepath.Abs(topdir)
	if err != nil {
		return err
	}
	utils.DirWalk(topdir, func(root string, rel string) error {
		if filepath.Ext(rel) == ".csproj" {
			dn.builder.Printf("dotnet file: %s\n", rel)
			dn.project_files = append(dn.project_files, root+"/"+rel)
		}
		return nil
	})
	if len(dn.project_files) == 0 {
		return fmt.Errorf("no .csproj files found in %s", topdir)
	}

	for _, pf := range dn.project_files {
		// this is debatable. MS claims it is cross platform, and as such
		// it would deserve a path without platform/cpu
		// and yet MS reputation in that regard does not quite back this up
		// so we stay on the "safe" side and use the platform it was
		// compiled on in the path. This way, if necessary, we can
		// create different builds for different platforms later
		dir := filepath.Dir(pf)
		outdir := dn.builder.GetRepoPath() + "/dist/dotnet/linux/amd64/" + filepath.Base(dir)
		dn.builder.Printf("output to: %s\n", outdir)
		err := dn.compile_dot_net(pf, outdir)
		if err != nil {
			return err
		}

	}
	return nil
}
func (dn *dotnet) find_dotnet_version() string {
	return "6.0.418"
}
func (dn *dotnet) find_dotnet_install_dir() string {
	return "/opt/dotnet/" + dn.find_dotnet_version()
}
func (dn *dotnet) compile_dot_net(project_file string, outdir string) error {
	dir := filepath.Dir(project_file)
	relprojfile := filepath.Base(project_file)
	l := linux.New()
	l.SetMaxRuntime(time.Duration(1) * time.Minute)
	l.SetEnvironment(dn.create_env())
	compiler := dn.find_dotnet_install_dir() + "/dotnet/dotnet"
	com := []string{compiler, "build", "-v", "diagnostic", "--debug", "-o", outdir, relprojfile}
	dn.builder.Printf("Compiling %s in dir %s\n", relprojfile, dir)
	b, err := l.SafelyExecuteWithDir(com, dir, nil)
	if err != nil {
		dn.builder.Printf("failed\n%s\n%s\n", string(b), err)
		return fmt.Errorf("dotnet compile failed (%w)", err)
	}
	return nil
}
func (dn *dotnet) create_env() []string {
	env := map[string]string{
		"DOTNET_ROOT": dn.find_dotnet_install_dir(),
	}
	var res []string
	for k, v := range env {
		res = append(res, fmt.Sprintf("%s=%s", k, v))
	}
	return res
}
