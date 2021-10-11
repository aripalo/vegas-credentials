package mfa

import (
	"bufio"
	"context"
	"os"
	"strings"

	"github.com/aripalo/aws-mfa-credential-process/internal/config"
	"github.com/aripalo/aws-mfa-credential-process/internal/profile"
	"github.com/ncruces/zenity"
)

func getCliToken(ctx context.Context, flags config.Flags, profileConfig profile.Profile, out chan *Result, errors chan *error) {

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

	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')
	if err != nil {
		return result, err
	}

	result.Value = strings.TrimSpace(value)

	return result, err

}

func dialogInput(ctx context.Context) (Result, error) {

	var result Result
	result.Provider = TOKEN_PROVIDER_DIALOG

	value, err := zenity.Entry(
		"Enter TOPT MFA Token Code:",
		zenity.Title("Multifactor Authentication"),
		zenity.Context(ctx),
	)

	if err != nil {
		return result, err
	}

	result.Value = strings.TrimSpace(value)

	return result, err
}
