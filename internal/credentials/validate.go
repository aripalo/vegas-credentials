package credentials

import (
	"errors"
)

// Validate ensures the Credentials is of correct format
func (c *Credentials) Validate() error {

	if c.Version != AWS_CREDENTIAL_PROCESS_VERSION {
		return errors.New("Incorrect Version")
	}

	if c.AccessKeyID == "" {
		return errors.New("Missing AccessKeyID")
	}

	if c.SecretAccessKey == "" {
		return errors.New("Missing SecretAccessKey")
	}

	if c.SessionToken == "" {
		return errors.New("Missing SessionToken")
	}

	return nil
}
