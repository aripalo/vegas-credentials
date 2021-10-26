package response

import (
	"github.com/aripalo/vegas-credentials/internal/interfaces"
)

// GetCachedCredentials handles fetching cached Temporary Session Credentials
func (r *Response) GetCachedCredentials(a interfaces.AssumeCredentialProcess) error {

	var err error

	err = r.ReadFromCache(a)
	if err != nil {
		return err
	}

	err = r.Validate(a)
	if err != nil {
		return err
	}

	err = r.ValidateForMandatoryRefresh(a)
	if err != nil {
		return err
	}

	return nil
}

// GetNewCredentials handles fetching new Temporary Session Credentials from STS
func (r *Response) GetNewCredentials(a interfaces.AssumeCredentialProcess) error {

	var err error

	err = r.AssumeRole(a)
	if err != nil {
		return err
	}

	err = r.Validate(a)
	if err != nil {
		return err
	}

	err = r.SaveToCache(a)
	if err != nil {
		return err
	}

	return nil

}
