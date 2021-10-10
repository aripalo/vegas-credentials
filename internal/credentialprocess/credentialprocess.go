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

		parsed, err := parseCredentials(fresh)
		if err != nil {
			// TODO better error
			return nil, errors.New("Fresh data could not be converted to valid credential_process response")
		}

		utils.SafeLog(utils.FormatMessage(utils.COLOR_SUCCESS, "✅ ", "Session Credentials", "Received from STS"))
		utils.SafeLog(utils.FormatMessage(utils.COLOR_DEBUG, "ℹ️  ", "Session Credentials", utils.FormatExpirationMessage(parsed.Expiration)))

		validationErr := validate(parsed)
		if validationErr != nil {
			return nil, validationErr
		}
		err = cache.Save(profileName, config, fresh)
		utils.SafeLog(utils.FormatMessage(utils.COLOR_DEBUG, "ℹ️  ", "Session Credentials", "Saved to cache"))

		utils.SafeLog(utils.TextGrayDark(utils.CreateRuler("=")))

		return fresh, nil
	}

	return nil, err

}
