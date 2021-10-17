package provider

import (
	"context"
	"os/exec"

	"github.com/aripalo/aws-mfa-credential-process/internal/data"
)

func (t *TokenProvider) QueryYubikey(ctx context.Context, d data.Provider) {
	var token Token
	var err error

	p := d.GetProfile()
	token.Provider = TOKEN_PROVIDER_YUBIKEY_TOUCH

	cmd := exec.CommandContext(ctx, "ykman", "--device", p.YubikeySerial, "oath", "accounts", "code", p.YubikeyLabel)
	stdout, err := cmd.Output()
	if err != nil {
		t.errorChan <- &err
	} else {
		value := tokenPattern.FindString(string(stdout))
		token.Value = value
		t.tokenChan <- &token
	}
}
