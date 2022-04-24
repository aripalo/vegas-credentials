package password

import (
	"context"

	"github.com/aripalo/ykmangoath"
)

func authenticate(ctx context.Context, device string, password string) (bool, error) {

	oathAccounts, err := ykmangoath.New(ctx, device) // set up with a new ctx
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

	_, err = oathAccounts.List()
	if err != nil {
		if err == ykmangoath.ErrOathAccountPasswordIncorrect {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
