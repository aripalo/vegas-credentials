package resolver

import (
	"context"
	"strings"

	"github.com/aripalo/vegas-credentials/internal/multinput"
	"github.com/aripalo/vegas-credentials/internal/yubikey"
)

func Yubikey(y yubikey.Yubikey) multinput.InputResolver {
	return func(ctx context.Context) (*multinput.Result, error) {
		result, err := y.Code(ctx)
		code := strings.TrimSpace(result)
		return &multinput.Result{Value: code, ResolverID: ResolverIdYubikey}, err
	}
}
