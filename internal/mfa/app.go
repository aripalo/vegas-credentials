package mfa

import (
	"context"
	"strings"

	"github.com/aripalo/aws-mfa-credential-process/internal/config"
	"github.com/aripalo/aws-mfa-credential-process/internal/profile"
	"github.com/aripalo/aws-mfa-credential-process/internal/prompt"
)

func getAppToken(ctx context.Context, flags config.Flags, profileConfig profile.Profile, out chan *Result, errors chan *error) {

	var err error
	var result Result

	defer getTokenErrorHandler(ctx, err, errors)

	if flags.DisableDialog {
		result, err = cliIpunt(ctx)
	} else {
		result, err = dialogInput(ctx)
	}

	errors <- &err
	out <- &result
}

func cliIpunt(ctx context.Context) (Result, error) {
	var result Result
	result.Provider = TOKEN_PROVIDER_CLI

	value, err := prompt.Cli(ctx, "")
	if err != nil {
		return result, err
	}

	result.Value = value

	return result, err

}

func dialogInput(ctx context.Context) (Result, error) {

	var result Result
	result.Provider = TOKEN_PROVIDER_DIALOG

	value, err := prompt.Dialog(
		ctx,
		"Multifactor Authentication",
		"Enter TOPT MFA Token Code:",
	)

	if err != nil {
		return result, err
	}

	result.Value = strings.TrimSpace(value)

	return result, err
}
