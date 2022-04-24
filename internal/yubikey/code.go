package yubikey

import (
	"context"

	"github.com/aripalo/ykmangoath"
)

// Code is responsible for querying the TOTP code from Yubikey device.
func (y *Yubikey) Code(ctx context.Context) (string, error) {
	oathAccounts, err := ykmangoath.New(ctx, y.serial)
	if err != nil {
		return "", err
	}

	password, err := y.passwordCache.Read()
	if err != nil {
		return "", err
	}

	if password != "" {
		// set the password we already know (after yubikey init)
		err := oathAccounts.SetPassword(password)
		if err != nil {
			return "", err
		}
	}

	return oathAccounts.Code(y.label)
}
