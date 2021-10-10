package mfa

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aripalo/goawsmfa/internal/profile"
	"github.com/aripalo/goawsmfa/internal/utils"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
)

// The default timeout of Yubikey operations
const YUBIKEY_TIMEOUT_SECONDS = 15

func GetTokenResult(config profile.Profile, hideArns bool) (Result, error) {
	resultChan := make(chan *Result, 1)
	errorChan := make(chan *error, 1)

	ctx, cancel := context.WithTimeout(context.Background(), YUBIKEY_TIMEOUT_SECONDS*time.Second)
	defer cancel()

	if hideArns == false {
		utils.SafeLog(utils.TextBold(utils.TextGrayDark(fmt.Sprintf("👷 Role: %s", config.AssumeRoleArn))))
		utils.SafeLog(utils.TextBold(utils.TextGrayDark(fmt.Sprintf("🔒 MFA: %s", config.MfaSerial))))
	}

	hasYubikey := (config.YubikeySerial != "" && config.YubikeyLabel != "")

	if hasYubikey {
		go getYubikeyToken(ctx, config.YubikeySerial, config.YubikeyLabel, resultChan, errorChan)
	}
	go getCliToken(ctx, resultChan, errorChan)

	if hasYubikey {
		utils.SafeLog(utils.TextBold(utils.TextYellow("🔑 Touch Yubikey or enter TOPT MFA Token Code:")))
	} else {
		utils.SafeLog(utils.TextBold(utils.TextYellow("🔑 Enter TOPT MFA Token Code:")))
	}

	w := wow.New(utils.GetSafeWriter(), spin.Get(spin.Dots3), "  ")
	w.Start()

	select {
	case i := <-resultChan:
		result := *i
		err := validateToken(result.Value)
		if err != nil {
			w.Stop()
			return result, err
		}
		w.PersistWith(spin.Spinner{Frames: []string{"✅ "}}, utils.TextGreen(fmt.Sprintf("Token %s received via %s", printMaskedToken(result.Value), result.Provider)))
		return result, nil
	case <-ctx.Done():
		w.Stop()
		if ctx.Err() == context.DeadlineExceeded {
			return Result{}, errors.New("MFA Operation Timeout")
		}
		return Result{}, ctx.Err()
	}
}

func printMaskedToken(token string) string {
	tokenLength := len(token)
	masked := ""

	for i := 0; i < tokenLength-2; i++ {
		masked += "*"
	}

	return fmt.Sprintf("%s%s%s", token[:1], masked, token[tokenLength-1:])
}
