package yubikey

import (
	"context"

	"github.com/aripalo/ykmangoath"
)

func (y *Yubikey) HasAccount() (bool, error) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	oathAccounts, err := ykmangoath.New(ctx, y.serial)
	if err != nil {
		return false, err
	}

	password, err := y.passwordCache.Read()
	if err != nil {
		return false, err
	}

	if password != "" {
		// set the password we already know (after yubikey init)
		err := oathAccounts.SetPassword(password)
		if err != nil {
			return false, err
		}
	}

	ok, err := oathAccounts.HasAccount(y.label)
	if err != nil {
		return false, err
	}

	return ok, nil
}
