package resolver

import (
	"context"

	"github.com/aripalo/vegas-credentials/internal/multinput"
	"github.com/aripalo/vegas-credentials/internal/prompt"
)

var cliPrompt = prompt.Cli

func CLI(ctx context.Context) (*multinput.Result, error) {
	value, err := cliPrompt(ctx, "")
	return &multinput.Result{Value: value, ResolverID: ResolverIdCliStdin}, err
}
