package provider

import (
	"context"

	"github.com/aripalo/aws-mfa-credential-process/internal/data"
)

func (t *TokenProvider) FromGui(ctx context.Context, d data.Provider) {
	var token Token
	var err error

	token.Provider = TOKEN_PROVIDER_GUI_DIALOG_PROMPT
	token.Value = "222222"

	t.tokenChan <- &token
	t.errorChan <- &err
}
