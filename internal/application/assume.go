package application

import (
	"context"
	"errors"
	"fmt"

	"github.com/aripalo/vegas-credentials/internal/assumable"
	"github.com/aripalo/vegas-credentials/internal/config/locations"
	"github.com/aripalo/vegas-credentials/internal/credentials"
	"github.com/aripalo/vegas-credentials/internal/msg"
	"github.com/aripalo/vegas-credentials/internal/totp"
	"github.com/dustin/go-humanize"
)

type AssumeFlags struct {
	Profile string `mapstructure:"profile"`
}

func (app *App) Assume(flags AssumeFlags) error {

	opts, err := assumable.New(locations.AwsConfig, flags.Profile)
	if err != nil {
		msg.Bail(fmt.Sprintf("Credentials: Error: %s", err))
	}

	msg.Message.Debugln("ℹ️", fmt.Sprintf("Credentials: Role: %s", opts.RoleArn))

	creds := credentials.New(opts)

	if err = creds.FetchFromCache(); err != nil {
		msg.Message.Debugln("ℹ️", fmt.Sprintf("Credentials: Cached: %s", err))
		msg.Message.Debugln("ℹ️", "Credentials: STS: Fetching...")
		msg.Message.Debugln("ℹ️", fmt.Sprintf("MFA: TOTP: %s", opts.MfaSerial))

		// TODO refactor this
		t := totp.New(totp.TotpOptions{
			YubikeySerial: opts.YubikeySerial,
			YubikeyLabel:  opts.YubikeyLabel,
			EnableGui:     !app.NoGui,
		})

		err = creds.FetchFromAWS(creds.BuildProvider(t.Get))
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				msg.Bail(fmt.Sprintf("Operation Timeout"))
			}
			msg.Bail(fmt.Sprintf("Credentials: STS: %s", err))
		} else {
			msg.Message.Successln("✅", fmt.Sprintf("Credentials: STS: Received fresh credentials"))
			msg.Message.Infoln("⏳", fmt.Sprintf("Credentials: STS: Expiration in %s", humanize.Time(creds.Expiration)))
		}
	} else {
		msg.Message.Successln("✅", "Credentials: Cached: Received")
		msg.Message.Infoln("⏳", fmt.Sprintf("Credentials: Cached: Expiration in %s", humanize.Time(creds.Expiration)))
	}

	// TODO same for passwd cache
	err = creds.Teardown()
	if err != nil {
		return err
	}

	msg.Message.HorizontalRuler()

	return creds.Output()
}
