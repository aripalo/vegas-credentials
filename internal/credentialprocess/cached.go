package credentialprocess

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/aripalo/goawsmfa/internal/cache"
	"github.com/aripalo/goawsmfa/internal/profile"
	"github.com/aripalo/goawsmfa/internal/utils"
)

func getCachedTemporaryCredentials(profileName string, config profile.Profile) (json.RawMessage, error) {
	cached, cacheErr := cache.Get(profileName, config)
	if cacheErr != nil {
		utils.SafeLog("NOT found from cache")
		cache.Remove(profileName, config)
		//utils.SafeLog("Cached does not contain valid Temporary Credentials")
		return nil, cacheErr
	}

	expirationErr := ensureNotExpired(cached)
	if expirationErr != nil {
		utils.SafeLog("Found from cache, but expired, ignoring...")
		cache.Remove(profileName, config)
		return nil, expirationErr
	}

	utils.SafeLog("FOUND from cache")

	return cached, nil
}

func ensureNotExpired(cached json.RawMessage) error {
	var response CredentialProcessResponse

	err := json.Unmarshal(cached, &response)
	if err != nil {
		return errors.New("Cached data could not be converted to valid credential_process response")
	}

	now := time.Now()

	if response.Expiration.Before(now) {
		return errors.New("Cached credentials expired")
	}

	return nil
}
