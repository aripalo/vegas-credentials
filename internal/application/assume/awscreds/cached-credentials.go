package awscreds

import (
	"github.com/aripalo/aws-mfa-credential-process/internal/application/assume/awscreds/response"
	"github.com/aripalo/aws-mfa-credential-process/internal/data"
)

// getCachedCredentials handles fetching cached Temporary Session Credentials from secure keyring
func getCachedCredentials(d data.Provider, r *response.Response) error {

	var err error
	//r := response.New()

	err = r.ReadFromCache(d)
	if err != nil {
		return err
	}

	err = r.Validate(d)
	if err != nil {
		return err
	}

	err = r.ValidateForMandatoryRefresh(d)
	if err != nil {
		return err
	}

	return nil
}
