package coderunners

import (
	"context"
	"flag"
	"fmt"
	"golang.conradwood.net/apis/protorenderer"
	"golang.conradwood.net/go-easyops/utils"
	"golang.yacloud.eu/apis/protomanager"
	"strings"
)

var (
	use_pm = flag.Bool("use_protomanager", true, "if true use protomanager instead of protorenderer")
)

type protobuilder struct {
	ctx context.Context
	b   brunner
}

func (g protobuilder) Run(ctx context.Context, b brunner) error {
	g.b = b
	g.ctx = ctx
	d := b.GetRepoPath() + "/protos/"
	if !utils.FileExists(d) {
		return fmt.Errorf("dir %s does not exist", d)
	}
	err := utils.DirWalk(d, g.submitFile)
	if err != nil {
		return err
	}
	b.Printf("protobuilder completed with no errors\n")
	return nil
}
func (g protobuilder) submitFile(root, rel_file string) error {
	if strings.HasPrefix(rel_file, "#") {
		return nil
	}
	if !strings.HasSuffix(rel_file, ".proto") {
		return nil
	}
	abs := root + "/" + rel_file
	pfilename := "protos/" + rel_file
	content, err := utils.ReadFile(abs)
	if err != nil {
		return err
	}
	if *use_pm {
		csr := &protomanager.CheckSubmitRequest{
			Filename:     pfilename,
			Content:      content,
			RepositoryID: g.b.BuildInfo().RepositoryID(),
		}
		cres, err := protomanager.GetProtoManagerClient().CheckAndSubmit(g.ctx, csr)
		if err != nil {
			return err
		}
		if cres.IsValid {
			return nil
		}
		return fmt.Errorf("%s", cres.ErrorMessage)
	}

	// use oldstyle
	apr := &protorenderer.AddProtoRequest{
		Name:         pfilename,
		Content:      string(content),
		RepositoryID: g.b.BuildInfo().RepositoryID(),
	}
	cr := &protorenderer.CompileRequest{
		Compilers:       []protorenderer.CompilerType{protorenderer.CompilerType_GOLANG}, // fastest, good for testing if compile works
		AddProtoRequest: apr,
	}
	cres, err := protorenderer.GetProtoRendererClient().CompileFile(g.ctx, cr)
	if err != nil {
		return err
	}
	if cres.CompileError != "" {
		return fmt.Errorf("Whilst compiling proto %s: %s", rel_file, cres.CompileError)
	}
	ur, err := protorenderer.GetProtoRendererClient().UpdateProto(g.ctx, apr)
	if err != nil {
		return err
	}
	fmt.Printf("Proto %s will be in version %d\n", pfilename, ur.Version)
	//	fmt.Printf("Root: %s, Rel: %s\n", root, rel_file)
	return nil
}
