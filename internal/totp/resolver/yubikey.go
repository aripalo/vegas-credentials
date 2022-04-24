package resolver

import (
	"context"
	"vegas3/internal/multinput"
	"vegas3/internal/yubikey"
)

func ConfigureYubikey(options yubikey.Options) multinput.InputResolver {
	y, err := yubikey.New(options)
	if err != nil {
		// To avoid nil pointer reference, return just a resolver that resolves
		// into an emtpy value with an error
		return func(ctx context.Context) (*multinput.Result, error) {
			return &multinput.Result{Value: "", ResolverID: ResolverIdYubikey}, err
		}
	}
	return func(ctx context.Context) (*multinput.Result, error) {
		result, err := y.Code(ctx)
		return &multinput.Result{Value: result, ResolverID: ResolverIdYubikey}, err
	}
}
