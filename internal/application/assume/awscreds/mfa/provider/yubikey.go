package provider

import (
	"context"

	"github.com/aripalo/aws-mfa-credential-process/internal/data"
)

func (t *TokenProvider) FromYubikey(ctx context.Context, d data.Provider) {
	var token Token
	var err error

	token.Provider = TOKEN_PROVIDER_YUBIKEY_TOUCH
	token.Value = "333333"

	t.tokenChan <- &token
	t.errorChan <- &err
}
