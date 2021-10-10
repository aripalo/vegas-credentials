package mfa

import (
	"context"
	"errors"
	"time"

	"github.com/aripalo/goawsmfa/internal/utils"
)

func GetTokenResult(yubikeySerial string, yubikeyLabel string) (Result, error) {
	resultChan := make(chan *Result, 1)
	errorChan := make(chan *error, 1)

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	hasYubikey := (yubikeySerial != "" && yubikeyLabel != "")

	if hasYubikey {
		go getYubikeyToken(ctx, yubikeySerial, yubikeyLabel, resultChan, errorChan)
	}
	go getCliToken(ctx, resultChan, errorChan)

	if hasYubikey {
		utils.SafeLog("Touch Yubikey or enter TOPT MFA Token Code:")
	} else {
		utils.SafeLog("Enter TOPT MFA Token Code:")
	}

	select {
	case i := <-resultChan:
		result := *i
		err := validateToken(result.Value)
		if err != nil {
			return result, err
		}
		return result, nil
	case <-ctx.Done():
		if ctx.Err() == context.DeadlineExceeded {
			return Result{}, errors.New("Operation Timeout")
		}
		return Result{}, ctx.Err()
	}
}
