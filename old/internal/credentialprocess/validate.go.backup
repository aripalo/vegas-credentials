package credentialprocess

import (
	"errors"
	"time"
)

func validate(data CredentialProcessResponse) error {

	invalidErr := errors.New("Invalid session credentials")

	if data.AccessKeyID == "" {
		return invalidErr
	}

	if data.SecretAccessKey == "" {
		return invalidErr
	}

	if data.SessionToken == "" {
		return invalidErr
	}

	now := time.Now()

	if data.Expiration.Before(now) {
		return invalidErr
	}

	return nil
}
