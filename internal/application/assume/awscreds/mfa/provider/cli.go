package provider

import (
	"context"

	"github.com/aripalo/aws-mfa-credential-process/internal/data"
	"github.com/aripalo/aws-mfa-credential-process/internal/prompt"
)

func (t *TokenProvider) QueryCLI(ctx context.Context, d data.Provider) {
	var token Token
	var err error

	token.Provider = TOKEN_PROVIDER_CLI_INPUT

	value, err := prompt.Cli(ctx, "")
	if err != nil {
		t.errorChan <- &err
	} else {
		token.Value = value
		t.tokenChan <- &token
	}

}
