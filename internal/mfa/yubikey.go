package mfa

import (
	"context"
	"os/exec"
)

func getYubikeyToken(ctx context.Context, yubikeySerial string, yubikeyLabel string, out chan *Result, errors chan *error) {

	var err error
	var result Result
	result.Provider = TOKEN_PROVIDER_YUBIKEY

	defer getTokenErrorHandler(ctx, err, errors)

	cmd := exec.CommandContext(ctx, "ykman", "--device", yubikeySerial, "oath", "accounts", "code", yubikeyLabel)
	stdout, err := cmd.Output()

	token := tokenPattern.FindString(string(stdout))

	if token != "" {
		result.Value = token
		out <- &result
	}
}
