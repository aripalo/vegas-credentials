package mfa

import (
	"context"
	"time"

	"github.com/aripalo/aws-mfa-credential-process/internal/data"
	"github.com/aripalo/aws-mfa-credential-process/internal/logger"
	"github.com/aripalo/aws-mfa-credential-process/internal/mfa/provider"
)

func GetToken(d data.Provider) (provider.Token, error) {
	var t provider.TokenProvider
	p := d.GetProfile()

	enableYubikey := useYubikey(d)

	logger.Debugln(d, "🔒", "MFA", p.MfaSerial)
	logger.Promptf(d, "🔑", "MFA", "> ")

	token, err := t.Provide(d, enableYubikey)

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

// useYubikey decides if Yubikey should be used for MFA and also prints out debug messages for the user
func useYubikey(d data.Provider) bool {
	p := d.GetProfile()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	yubikeyErr := provider.VerifyYubikey(ctx, d)

	if yubikeyErr != nil {

		switch yubikeyErr.Error() {
		case provider.YubikeyErrorNotConfigured:
			logger.Debugln(d, "ℹ️ ", "MFA", "Yubikey not configured, ignoring...")
		case provider.YubikeyErrorFail:
			logger.Debugf(d, "ℹ️ ", "MFA", "Yubikey %s configured but can't access it, ignoring it...\n", p.YubikeySerial)
		case provider.YubikeyErrorNotConnected:
			logger.Debugf(d, "ℹ️ ", "MFA", "Yubikey %s configured but not connected, ignoring it...\n", p.YubikeySerial)
		default:
			logger.Debugf(d, "ℹ️ ", "MFA", "Yubikey %s configured but failed \"%s\", ignoring it...\n", p.YubikeySerial, yubikeyErr.Error())
		}

		return false
	} else {
		logger.Debugf(d, "ℹ️ ", "MFA", "Yubikey %s configured & connected!\n", p.YubikeySerial)
		return true
	}
}
