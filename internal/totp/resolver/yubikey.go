package resolver

import (
	"context"

	"github.com/aripalo/vegas-credentials/internal/multinput"
	"github.com/aripalo/vegas-credentials/internal/yubikey"
)

func Yubikey(y yubikey.Yubikey) multinput.InputResolver {
	return func(ctx context.Context) (*multinput.Result, error) {
		result, err := y.Code(ctx)
		return &multinput.Result{Value: result, ResolverID: ResolverIdYubikey}, err
	}
}
