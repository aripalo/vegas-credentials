package credentialprocess

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/aripalo/aws-mfa-credential-process/internal/cache"
	"github.com/aripalo/aws-mfa-credential-process/internal/config"
	"github.com/aripalo/aws-mfa-credential-process/internal/profile"
	"github.com/aripalo/aws-mfa-credential-process/internal/utils"
	"github.com/dustin/go-humanize"
)

func getCachedTemporaryCredentials(flags config.Flags, profileConfig profile.Profile) (json.RawMessage, error) {
	cached, cacheErr := cache.Get(flags.ProfileName, profileConfig)
	if cacheErr != nil {

		msg := utils.FormatMessage(utils.COLOR_DEBUG, "ℹ️  ", "Session Credentials", "NOT found from cache")
		utils.SafeLogLn(msg)

		cache.Remove(flags.ProfileName, profileConfig)
		return nil, cacheErr
	}

	parsed, err := parseCredentials(cached)
	if err != nil {
		return nil, errors.New("Cached data could not be converted to valid credential_process response")
	}

	expirationErr := ensureNotExpired(parsed)
	if expirationErr != nil {

		msg := utils.FormatMessage(utils.COLOR_DEBUG, "ℹ️  ", "Session Credentials", fmt.Sprintf("Found from cache, but expired at %s", humanize.Time(parsed.Expiration)))
		utils.SafeLogLn(msg)

		cache.Remove(flags.ProfileName, profileConfig)
		return nil, expirationErr
	}

	if flags.DisableRefresh != true {

		mandatoryRefreshErr := ensureMandatoryRefreshNotNeeded(parsed)
		if mandatoryRefreshErr != nil {

			msg := utils.FormatMessage(utils.COLOR_DEBUG, "ℹ️  ", "Session Credentials", fmt.Sprintf("Found from cache, but expiring in %s so mandatory refresh required", humanize.Time(parsed.Expiration)))
			utils.SafeLogLn(msg)

			cache.Remove(flags.ProfileName, profileConfig)
			return nil, mandatoryRefreshErr
		}
	}

	validationErr := validate(parsed)
	if validationErr != nil {
		msg := utils.FormatMessage(utils.COLOR_DEBUG, "ℹ️  ", "Session Credentials", "Found from cache, but invalid")
		utils.SafeLogLn(msg)

		cache.Remove(flags.ProfileName, profileConfig)
		return nil, validationErr
	}

	if flags.Verbose {
		utils.SafeLogLn(utils.FormatMessage(utils.COLOR_SUCCESS, "✅ ", "Session Credentials", "FOUND from cache"))
		utils.SafeLogLn(utils.FormatMessage(utils.COLOR_DEBUG, "ℹ️  ", "Session Credentials", utils.FormatExpirationInMessage(parsed.Expiration)))
		utils.SafeLogLn(utils.FormatMessage(utils.COLOR_DEBUG, "ℹ️  ", "Session Credentials", utils.FormatExpirationAtMessage(parsed.Expiration)))
		utils.SafeLogLn(utils.TextGrayDark(utils.CreateRuler("=")))
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

func ensureMandatoryRefreshNotNeeded(cached CredentialProcessResponse) error {
	// Match botocore mandatory timeout
	// https://github.com/boto/botocore/blob/221ffa67a567df99ee78d7ae308c0e12d7eeeea7/botocore/credentials.py#L350-L355
	now := time.Now()
	count := 10 * 60
	limit := now.Add(time.Duration(-count) * time.Second)

	if cached.Expiration.Before(limit) {
		return errors.New("Mandatory cached credentials expiration")
	}

	return nil
}
