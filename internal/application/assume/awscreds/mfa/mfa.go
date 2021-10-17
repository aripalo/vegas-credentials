package mfa

import (
	"fmt"

	"github.com/aripalo/aws-mfa-credential-process/internal/application/assume/awscreds/mfa/provider"
	"github.com/aripalo/aws-mfa-credential-process/internal/data"
)

func GetToken(d data.Provider) (provider.Token, error) {
	var t provider.TokenProvider

	fmt.Println("MFA:")
	token, err := t.Provide(d)
	if err != nil {
		fmt.Println("MFA error!", err)
	}

	return token, err
}
