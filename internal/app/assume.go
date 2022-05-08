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

	if err = creds.Load(); err != nil {
		msg.Debug("ℹ️", fmt.Sprintf("Credentials: Cached: %s", err))
		msg.Debug("ℹ️", fmt.Sprintf("MFA: Serial: %s", cfg.MfaSerial))

		code, err := totp.GetCode(context.Background(), totp.Options{
			EnableGui:     !a.NoGui,
			EnableYubikey: true, // TODO ??
			YubikeySerial: cfg.YubikeySerial,
			YubikeyLabel:  cfg.YubikeyLabel,
		})

		if err != nil {
			msg.Fatal(fmt.Sprintf("TOTP: %s", err))
		}

		err = creds.New(code)
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				msg.Fatal(fmt.Sprintf("Operation Timeout"))
			}
			msg.Fatal(fmt.Sprintf("Credentials: STS: %s", err))
		} else {
			msg.Success("✅", fmt.Sprintf("Credentials: STS: Received fresh credentials"))
			msg.Info("⏳", fmt.Sprintf("Credentials: STS: Expiration in %s", humanize.Time(creds.Expiration)))
		}
	} else {
		msg.Success("✅", "Credentials: Cached: Received")
		msg.Info("⏳", fmt.Sprintf("Credentials: Cached: Expiration in %s", humanize.Time(creds.Expiration)))
	}

	// TODO same for passwd cache
	err = creds.Teardown()
	if err != nil {
		return err
	}

	msg.HorizontalRuler()

	return creds.Output()
}
