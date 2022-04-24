package yubikey

import (
	"context"
	"errors"
	"fmt"
	"time"
	"vegas3/internal/yubikey/passcache"
	"vegas3/internal/yubikey/password"
)

// Struct to hold internal configuration
type Yubikey struct {
	serial        string
	label         string
	enableGui     bool
	passwordCache password.PasswordCache
}

// Options passed in by the caller
type Options struct {
	YubikeySerial string
	YubikeyLabel  string
	EnableGui     bool
}

func New(options Options) (Yubikey, error) {
	y := Yubikey{
		serial:        options.YubikeySerial,
		label:         options.YubikeyLabel,
		enableGui:     options.EnableGui,
		passwordCache: passcache.New(options.YubikeySerial),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Hour))
	defer cancel()

	err := password.Resolve(ctx, y.serial, y.passwordCache, y.enableGui)
	if err != nil {
		return y, err
	}

	hasAccount, err := y.HasAccount()
	if err != nil {
		return y, err
	}

	if !hasAccount {
		return y, errors.New(fmt.Sprintf("Yubikey: OATH Account not configured: %s", y.label))
	}

	return y, nil
}
