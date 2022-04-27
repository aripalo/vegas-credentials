package password

import (
	"context"

	"github.com/aripalo/ykmangoath"
)

func check(ctx context.Context, device string) (bool, error) {

	oathAccounts, err := ykmangoath.New(ctx, device) // set up with a new ctx
	if err != nil {
		return false, err
	}

	isPasswordProtected := oathAccounts.IsPasswordProtected()

	return isPasswordProtected, nil
}
