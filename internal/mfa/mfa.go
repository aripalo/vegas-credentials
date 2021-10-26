package mfa

import (
	"context"
	"time"

	"github.com/aripalo/vegas-credentials/internal/interfaces"
	"github.com/aripalo/vegas-credentials/internal/logger"
	"github.com/aripalo/vegas-credentials/internal/mfa/provider"
)

func GetToken(a interfaces.AssumeCredentialProcess) (provider.Token, error) {
	p := a.GetProfile()

	enableYubikey := useYubikey(a)

	logger.Debugln(a, "🔒", "MFA", p.Source.MfaSerial)
	logger.Promptf(a, "🔑", "MFA", "> ")

	t := provider.New(a, enableYubikey)
	token, err := t.Provide(a, enableYubikey)

	// No need to print newline if user entered token via CLI stdin as it contained "enter-press"
	if token.Provider != provider.TOKEN_PROVIDER_CLI_INPUT {
		logger.Newline(a)
	}

	if err != nil {
		if err == context.DeadlineExceeded {
			logger.Errorln(a, "❌", "MFA", "Error: Timeout exceeded")
		} else {
			logger.Errorf(a, "❌", "MFA", "Error: %s\n   Input Method: %s\n   Received Value: \"%s\"\n\n   Valid token value should contain at least 6 digits.\n   Check your configuration/input and try again.\n\n%s", err.Error(), token.Provider, token.Value, logger.GetSupportString("   "))
		}
		return token, err
	}

	logger.Successf(a, "🔓", "MFA", "OATH TOPT MFA Token %s received from %s\n", token.Value, token.Provider)
	return token, err
}

// useYubikey decides if Yubikey should be used for MFA and also prints out debug messages for the user
func useYubikey(a interfaces.AssumeCredentialProcess) bool {
	p := a.GetProfile()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	yubikeyErr := provider.VerifyYubikey(ctx, a)

	if yubikeyErr != nil {

		switch yubikeyErr.Error() {
		case provider.YubikeyErrorNotConfigured:
			logger.Debugln(a, "ℹ️ ", "MFA", "Yubikey not configured, ignoring...")
		case provider.YubikeyErrorFail:
			logger.Debugf(a, "ℹ️ ", "MFA", "Yubikey %s configured but can't access it, ignoring it...\n", p.Source.YubikeySerial)
		case provider.YubikeyErrorNotConnected:
			logger.Debugf(a, "ℹ️ ", "MFA", "Yubikey %s configured but not connected, ignoring it...\n", p.Source.YubikeySerial)
		default:
			logger.Debugf(a, "ℹ️ ", "MFA", "Yubikey %s configured but failed \"%s\", ignoring it...\n", p.Source.YubikeySerial, yubikeyErr.Error())
		}

		return false
	} else {
		logger.Debugf(a, "ℹ️ ", "MFA", "Yubikey %s configured & connected!\n", p.Source.YubikeySerial)
		return true
	}
}
