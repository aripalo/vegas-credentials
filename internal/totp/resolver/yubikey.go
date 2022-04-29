package resolver

import (
	"context"
	"fmt"

	"github.com/aripalo/vegas-credentials/internal/msg"
	"github.com/aripalo/vegas-credentials/internal/multinput"
	"github.com/aripalo/vegas-credentials/internal/yubikey2"
)

func ConfigureYubikey(options yubikey2.Options) multinput.InputResolver {
	y, err := yubikey2.New(options)

	fmt.Println(y)

	if err != nil {

		msg.Message.Warningln("⚠️", "YUBIERR: "+err.Error())

		// TODO fix this
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
