package coderunners

import (
	"context"
	"strings"

	"golang.conradwood.net/go-easyops/errors"
	"golang.conradwood.net/go-easyops/utils"
	"golang.yacloud.eu/apis/protomanager"
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
		return errors.Errorf("(dir \"[repo]/protos\" (%s) does not exist", d)
	}
	err := utils.DirWalk(d, g.submitFile)
	if err != nil {
		b.Printf("protobuilder encountered an error: %s\n", err)
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
	return errors.Errorf("%s", cres.ErrorMessage)

}
