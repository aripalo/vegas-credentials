package provider

import (
	"context"

	"github.com/aripalo/aws-mfa-credential-process/internal/data"
)

func (t *TokenProvider) FromCli(ctx context.Context, d data.Provider) {
	var token Token
	var err error

	token.Provider = TOKEN_PROVIDER_CLI_INPUT
	token.Value = "111111"

	t.tokenChan <- &token
	t.errorChan <- &err
}
