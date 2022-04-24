package application

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/aripalo/vegas-credentials/internal/assumable"
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

var cacheLocation string = filepath.Join(locations.CacheDir, "session-cache")

func (app *App) Assume(flags AssumeFlags) error {

	a, err := assumable.New(locations.AwsConfig, flags.Profile)
	if err != nil {
		utils.Bail(fmt.Sprintf("Credentials: Error: %s", err))
	}

	msg.Message.Debugln("ℹ️", fmt.Sprintf("Credentials: Role: %s", a.RoleArn))

	credentialsCache := credentials.NewCredentialCache()

	creds := credentials.New(credentialsCache, a)

	if err = creds.FetchFromCache(); err != nil {
		msg.Message.Debugln("ℹ️", fmt.Sprintf("Credentials: Cached: %s", err))
		msg.Message.Debugln("ℹ️", "Credentials: STS: Fetching...")
		msg.Message.Debugln("ℹ️", fmt.Sprintf("MFA: TOTP: %s", a.MfaSerial))

		t := totp.New(totp.TotpOptions{
			YubikeySerial: a.YubikeySerial,
			YubikeyLabel:  a.YubikeyLabel,
			EnableGui:     !app.NoGui,
		})

		err = creds.FetchFromAWS(t.Get)
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

	creds.Output()

	return nil
}
