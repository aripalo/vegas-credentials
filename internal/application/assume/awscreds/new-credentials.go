package awscreds

import (
	"github.com/aripalo/aws-mfa-credential-process/internal/application/assume/awscreds/response"
	"github.com/aripalo/aws-mfa-credential-process/internal/data"
)

// getNewCredentials handles fetching new Temporary Session Credentials from STS
func getNewCredentials(d data.Provider) (*response.Response, error) {

	var err error
	r := response.New()

	err = r.Get(d)
	if err != nil {
		return r, err
	}

	err = r.Validate(d)
	if err != nil {
		return r, err
	}

	err = r.SaveToCache(d)
	if err != nil {
		return r, err
	}

	return r, nil

}
