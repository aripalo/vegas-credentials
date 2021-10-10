package credentialprocess

import (
	"encoding/json"
	"errors"

	"github.com/aripalo/goawsmfa/internal/cache"
	"github.com/aripalo/goawsmfa/internal/profile"
	"github.com/aripalo/goawsmfa/internal/utils"
)

func GetOutput(verboseOutput bool, profileName string, hideArns bool, config profile.Profile) (json.RawMessage, error) {
	var err error

	cached, cacheErr := getCachedTemporaryCredentials(verboseOutput, profileName, config)

	if cacheErr == nil {
		return cached, nil
	}

	fresh, err := getFreshTemporaryCredentials(config, hideArns)
	if err == nil {
		if verboseOutput {
			utils.SafeLog(utils.TextGreen("✅ [Session Credential] Fetched new session credentials"))
		}

		parsed, err := parseCredentials(fresh)
		if err != nil {
			// TODO better error
			return nil, errors.New("Fresh data could not be converted to valid credential_process response")
		}

		validationErr := validate(parsed)
		if validationErr != nil {
			return nil, validationErr
		}
		err = cache.Save(profileName, config, fresh)
		return fresh, nil
	}

	return nil, err

}
