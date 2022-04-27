package yubikey

import (
	"context"

	"github.com/aripalo/ykmangoath"
)

func (y *Yubikey) IsAvailable() (bool, error) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	oathAccounts, err := ykmangoath.New(ctx, y.serial)
	if err != nil {
		return false, err
	}

	available := oathAccounts.IsAvailable()

	return available, nil
}
