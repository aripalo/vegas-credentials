package application

import (
	"context"
	"errors"
	"fmt"

	"github.com/aripalo/vegas-credentials/internal/assumeopts"
	"github.com/aripalo/vegas-credentials/internal/credentials"
	"github.com/aripalo/vegas-credentials/internal/locations"
	"github.com/aripalo/vegas-credentials/internal/msg"
	"github.com/aripalo/vegas-credentials/internal/totp"
	"github.com/aripalo/vegas-credentials/internal/utils"
	"github.com/dustin/go-humanize"
)

type AssumeFlags struct {
	Profile string `mapstructure:"profile"`
}

func (app *App) Assume(flags AssumeFlags) error {

	opts, err := assumeopts.New(locations.AwsConfig, flags.Profile)
	if err != nil {
		utils.Bail(fmt.Sprintf("Credentials: Error: %s", err))
	}

	msg.Message.Debugln("ℹ️", fmt.Sprintf("Credentials: Role: %s", opts.RoleArn))

	checksum, err := opts.Checksum()
	if err != nil {
		utils.Bail(fmt.Sprintf("Credentials: Error: %s", err))
	}

	// TODO refactor this
	t := totp.New(totp.TotpOptions{
		YubikeySerial: opts.YubikeySerial,
		YubikeyLabel:  opts.YubikeyLabel,
		EnableGui:     !app.NoGui,
	})

	assumeRoleProvider := opts.BuildAssumeRoleProvider(t.Get)

	creds := credentials.New(credentials.Options{
		Name:               opts.ProfileName,
		SourceProfile:      opts.SourceProfile,
		Region:             opts.Region,
		RoleArn:            opts.RoleArn,
		Checksum:           checksum,
		AssumeRoleProvider: assumeRoleProvider,
	})

	if err = creds.FetchFromCache(); err != nil {
		msg.Message.Debugln("ℹ️", fmt.Sprintf("Credentials: Cached: %s", err))
		msg.Message.Debugln("ℹ️", "Credentials: STS: Fetching...")
		msg.Message.Debugln("ℹ️", fmt.Sprintf("MFA: TOTP: %s", opts.MfaSerial))

		err = creds.FetchFromAWS()
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				utils.Bail(fmt.Sprintf("Operation Timeout"))
			}
			utils.Bail(fmt.Sprintf("Credentials: STS: %s", err))
		} else {
			msg.Message.Successln("✅", fmt.Sprintf("Credentials: STS: Received fresh credentials"))
			msg.Message.Infoln("⏳", fmt.Sprintf("Credentials: STS: Expiration in %s", humanize.Time(creds.Expiration)))
		}
	} else {
		msg.Message.Successln("✅", "Credentials: Cached: Received")
		msg.Message.Infoln("⏳", fmt.Sprintf("Credentials: Cached: Expiration in %s", humanize.Time(creds.Expiration)))
	}

	msg.Message.HorizontalRuler()

	return creds.Output()
}
