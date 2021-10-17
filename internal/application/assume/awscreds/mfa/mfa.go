package mfa

import (
	"context"

	"github.com/aripalo/aws-mfa-credential-process/internal/application/assume/awscreds/mfa/provider"
	"github.com/aripalo/aws-mfa-credential-process/internal/data"
	"github.com/aripalo/aws-mfa-credential-process/internal/logger"
)

func GetToken(d data.Provider) (provider.Token, error) {
	var t provider.TokenProvider
	p := d.GetProfile()

	logger.Debugln(d, "🔒", "MFA", p.MfaSerial)
	logger.Promptf(d, "🔑", "MFA", "> ")

	token, err := t.Provide(d)

	// No need to print newline if user entered token via CLI stdin as it contained "enter-press"
	if token.Provider != provider.TOKEN_PROVIDER_CLI_INPUT {
		logger.Newline(d)
	}

	if err != nil {
		if err == context.DeadlineExceeded {
			logger.Errorln(d, "❌", "MFA", "Error: Timeout exceeded")
		} else {
			logger.Errorf(d, "❌", "MFA", "Error: %s\n   Input Method: %s\n   Received Value: \"%s\"\n\n   Valid token value should contain at least 6 digits.\n   Check your configuration/input and try again.\n\n%s", err.Error(), token.Provider, token.Value, logger.GetSupportString("   "))
		}
		return token, err
	}

	logger.Successf(d, "🔓", "MFA", "OATH TOPT MFA Token %s received from %s\n", token.Value, token.Provider)
	return token, err
}
