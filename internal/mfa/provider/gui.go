package provider

import (
	"context"

	"github.com/aripalo/vegas-credentials/internal/interfaces"
	"github.com/aripalo/vegas-credentials/internal/prompt"
)

var guiPrompt = prompt.Dialog

func (t *TokenProvider) QueryGUI(ctx context.Context, a interfaces.AssumeCredentialProcess) {
	var token Token
	var err error

	token.Provider = TOKEN_PROVIDER_GUI_DIALOG_PROMPT

	value, err := guiPrompt(ctx, "Multifactor Authentication", "Enter TOTP MFA Token Code:")
	if err != nil {
		t.errorChan <- &err
	} else {
		token.Value = value
		t.tokenChan <- &token
	}
}
