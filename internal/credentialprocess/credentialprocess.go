package credentialprocess

import (
	"encoding/json"

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
		utils.SafeLogger.Println("Fetched new session credentials")
		err = cache.Save(profileName, config, fresh)
		return fresh, nil
	}

	return nil, err

}
