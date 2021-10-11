package mfa

import (
	"context"
	"os/exec"

	"github.com/aripalo/aws-mfa-credential-process/internal/config"
	"github.com/aripalo/aws-mfa-credential-process/internal/profile"
)

func getYubikeyToken(ctx context.Context, flags config.Flags, profileConfig profile.Profile, out chan *Result, errors chan *error) {

	var err error
	var result Result
	result.Provider = TOKEN_PROVIDER_YUBIKEY

	defer getTokenErrorHandler(ctx, err, errors)

	cmd := exec.CommandContext(ctx, "ykman", "--device", profileConfig.YubikeySerial, "oath", "accounts", "code", profileConfig.YubikeyLabel)
	stdout, err := cmd.Output()

	token := tokenPattern.FindString(string(stdout))

	if token != "" {
		result.Value = token
		out <- &result
	}
}
