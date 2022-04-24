package resolver

import (
	"context"
	"vegas3/internal/multinput"
	"vegas3/internal/prompt"
)

var cliPrompt = prompt.Cli

func CLI(ctx context.Context) (*multinput.Result, error) {
	value, err := cliPrompt(ctx, "")
	return &multinput.Result{Value: value, ResolverID: ResolverIdCliStdin}, err
}

func ConfigureCLI() multinput.InputResolver {
	return CLI
}
