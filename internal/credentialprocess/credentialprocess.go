package credentialprocess

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/aripalo/goawsmfa/internal/cache"
	"github.com/aripalo/goawsmfa/internal/profile"
	"github.com/aripalo/goawsmfa/internal/utils"
)

func GetOutput(profileName string, config profile.Profile) (json.RawMessage, error) {
	var err error

	cached, cacheErr := getCachedTemporaryCredentials(profileName, config)

	if cacheErr == nil {
		return cached, nil
	}

	fresh, err := getFreshTemporaryCredentials(config)
	if err == nil {
		utils.SafeLog("Fetched new session credentials")
		err = cache.Save(profileName, config, fresh)
		validationErr := validate(fresh)
		if validationErr != nil {
			return nil, validationErr
		}
		return fresh, nil
	}

	return nil, err

}

func validate(data json.RawMessage) error {

	var r CredentialProcessResponse

	invalidErr := errors.New("Invalid session credentials")

	err := json.Unmarshal(data, &r)
	if err != nil {
		return invalidErr
	}

	if r.AccessKeyID == "" {
		return invalidErr
	}

	if r.SecretAccessKey == "" {
		return invalidErr
	}

	if r.SessionToken == "" {
		return invalidErr
	}

	now := time.Now()

	if r.Expiration.Before(now) {
		return invalidErr
	}

	return nil
}
