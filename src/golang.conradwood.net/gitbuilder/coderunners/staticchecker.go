package coderunners

import (
	"context"
)

type staticcheck struct {
}

func (g *staticcheck) Run(ctx context.Context, b brunner) error {
	b.Printf("static check invoked but not implemented")
	return nil
}
