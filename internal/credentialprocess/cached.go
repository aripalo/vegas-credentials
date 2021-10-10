package credentialprocess

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/aripalo/goawsmfa/internal/cache"
	"github.com/aripalo/goawsmfa/internal/profile"
	"github.com/aripalo/goawsmfa/internal/utils"
)

func getCachedTemporaryCredentials(verboseOutput bool, profileName string, config profile.Profile) (json.RawMessage, error) {
	cached, cacheErr := cache.Get(profileName, config)
	if cacheErr != nil {
		if verboseOutput {
			utils.SafeLog(utils.TextGray("ℹ️  [Session Credential] NOT found from cache"))
		}
		cache.Remove(profileName, config)
		//utils.SafeLog("Cached does not contain valid Temporary Credentials")
		return nil, cacheErr
	}

	parsed, err := parseCredentials(cached)
	if err != nil {
		return nil, errors.New("Cached data could not be converted to valid credential_process response")
	}

	expirationErr := ensureNotExpired(parsed)
	if expirationErr != nil {
		if verboseOutput {
			utils.SafeLog(utils.TextGray("ℹ️  [Session Credential] Found from cache, but expired, ignoring..."))
		}
		cache.Remove(profileName, config)
		return nil, expirationErr
	}

	advisoryRefreshErr := ensureAdvisoryRefreshNotNeeded(parsed)
	if advisoryRefreshErr != nil {
		if verboseOutput {
			utils.SafeLog(utils.TextGray("ℹ️  [Session Credential] Found from cache, but advisory refresh required, ignoring..."))
		}
		cache.Remove(profileName, config)
		return nil, advisoryRefreshErr
	}

	validationErr := validate(parsed)
	if validationErr != nil {
		if verboseOutput {
			utils.SafeLog(utils.TextGray("ℹ️  [Session Credential] Found from cache, but invalid, ignoring..."))
		}
		cache.Remove(profileName, config)
		return nil, validationErr
	}

	if verboseOutput {
		utils.SafeLog(utils.TextGreen("✅  [Session Credential] FOUND from cache!"))
	}

	return cached, nil
}

func parseCredentials(data json.RawMessage) (CredentialProcessResponse, error) {
	var response CredentialProcessResponse
	err := json.Unmarshal(data, &response)
	return response, err
}

func ensureNotExpired(cached CredentialProcessResponse) error {
	now := time.Now()

	if cached.Expiration.Before(now) {
		return errors.New("Cached credentials expired")
	}

	return nil
}

func ensureAdvisoryRefreshNotNeeded(cached CredentialProcessResponse) error {
	// Match botocore advisory timeout
	// https://github.com/boto/botocore/blob/221ffa67a567df99ee78d7ae308c0e12d7eeeea7/botocore/credentials.py#L350-L355
	now := time.Now()
	count := 15 * 60
	limit := now.Add(time.Duration(-count) * time.Second)

	if cached.Expiration.Before(limit) {
		return errors.New("Cached credentials expired")
	}

	return nil
}
