package resolver

import (
	"context"
	"strings"

	"github.com/aripalo/vegas-credentials/internal/multinput"
	"github.com/aripalo/vegas-credentials/internal/prompt"
)

var cliPrompt = prompt.Cli

func CLI(ctx context.Context) (*multinput.Result, error) {
	result, err := cliPrompt(ctx, "")
	code := strings.TrimSpace(result)
	return &multinput.Result{Value: code, ResolverID: ResolverIdCliStdin}, err
}
