package app

import (
	"context"
	"errors"
	"fmt"

	"github.com/aripalo/vegas-credentials/internal/assumecfg"
	"github.com/aripalo/vegas-credentials/internal/config/locations"
	"github.com/aripalo/vegas-credentials/internal/credentials"
	"github.com/aripalo/vegas-credentials/internal/msg"
	"github.com/aripalo/vegas-credentials/internal/totp"
	"github.com/dustin/go-humanize"
)

// CLI flags for this command.
type AssumeFlags struct {
	Profile string `mapstructure:"profile"`
}

// Implementation of "assume" CLI command
// without any knowledge of spf13/cobra internals.
func (a *App) Assume(flags AssumeFlags) error {

	cfg, err := assumecfg.New(locations.AwsConfig, flags.Profile)
	if err != nil {
		msg.Fatal(fmt.Sprintf("Credentials: Error: %s", err))
	}

	msg.Debug("ℹ️", fmt.Sprintf("Credentials: Role: %s", cfg.RoleArn))

	creds := credentials.New(cfg)

	// Loading from cache is preferred, so if Temporary Session Credentials
	// are found from cache then just return early.
	if err = creds.Load(); err == nil {
		msg.Success("⏳", fmt.Sprintf("Credentials: Loaded from cache, expiration in %s", humanize.Time(creds.Expiration)))
		return exit(creds)
	}

	msg.Debug("ℹ️", fmt.Sprintf("Credentials: Cache: %s", err))

	code, err := totp.GetCode(context.Background(), totp.Options{
		EnableGui:     !a.NoGui,
		EnableYubikey: true, // TODO ??
		YubikeySerial: cfg.YubikeySerial,
		YubikeyLabel:  cfg.YubikeyLabel,
	})

	if err != nil {
		msg.Fatal(fmt.Sprintf("MFA: TOTP: %s", err))
	}

	msg.Debug("ℹ️", fmt.Sprintf("MFA: Serial: %s", cfg.MfaSerial))

	err = creds.New(code)

	// Catch timeout error and return a cleaner error message.
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			msg.Fatal(fmt.Sprintf("Operation Timeout"))
		}
		msg.Fatal(fmt.Sprintf("Credentials: STS: %s", err))
	}

	msg.Success("⏳", fmt.Sprintf("Credentials: New from STS, expiration in %s", humanize.Time(creds.Expiration)))
	return exit(creds)

}

// Shared "exit" handler
func exit(creds *credentials.Credentials) error {
	err := creds.Teardown()
	if err != nil {
		return err
	}

	msg.HorizontalRuler()
	return creds.Output()
}
