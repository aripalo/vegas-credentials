package response

import "github.com/aripalo/aws-mfa-credential-process/internal/data"

// GetCachedCredentials handles fetching cached Temporary Session Credentials from secure keyring
func (r *Response) GetCachedCredentials(d data.Provider) error {

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

// GetNewCredentials handles fetching new Temporary Session Credentials from STS
func (r *Response) GetNewCredentials(d data.Provider) error {

	var err error

	err = r.Get(d)
	if err != nil {
		return err
	}

	err = r.Validate(d)
	if err != nil {
		return err
	}

	err = r.SaveToCache(d)
	if err != nil {
		return err
	}

	return nil

}
