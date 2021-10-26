package provider

import (
	"context"

	"github.com/aripalo/vegas-credentials/internal/interfaces"
	"github.com/aripalo/vegas-credentials/internal/prompt"
)

func (t *TokenProvider) QueryCLI(ctx context.Context, a interfaces.AssumeCredentialProcess) {
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
