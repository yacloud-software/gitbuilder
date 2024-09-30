package coderunners

import "context"

type go_vet struct {
}

func (g go_vet) Run(ctx context.Context, b brunner) error {
	b.Printf("go-vet currently only has a stub\n")
	return nil
}
