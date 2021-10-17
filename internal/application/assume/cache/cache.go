package cache

import (
	"encoding/json"

	"github.com/aripalo/aws-mfa-credential-process/internal/application/assume/cache/cachekey"
	"github.com/aripalo/aws-mfa-credential-process/internal/application/assume/cache/securestorage"
	"github.com/aripalo/aws-mfa-credential-process/internal/profile"
)

func Get(profileName string, config profile.ProfileConfig) (json.RawMessage, error) {
	cacheKey, err := cachekey.Get(profileName, config)
	if err != nil {
		return nil, err
	}
	cached, err := securestorage.Get(cacheKey)
	return cached, err
}

func Save(profileName string, config profile.ProfileConfig, data json.RawMessage) error {
	cacheKey, err := cachekey.Get(profileName, config)
	err = securestorage.Set(cacheKey, data)
	return err
}

// Remove a given configuration from cache
func Remove(profileName string, config profile.ProfileConfig) error {
	cacheKey, err := cachekey.Get(profileName, config)
	err = securestorage.Remove(cacheKey)
	return err
}

// RemoveAll the hole cache or all items related to given profile
func RemoveAll(profileName string) error {
	/*
		if profileName != "" {
			utils.SafeLogLn(utils.FormatMessage(utils.COLOR_IMPORTANT, "ℹ️  ", "Cache", fmt.Sprintf("Deleting cache for profile \"%s\"", profileName)))
		} else {
			utils.SafeLogLn(utils.FormatMessage(utils.COLOR_IMPORTANT, "ℹ️  ", "Cache", "Deleting all items from cache"))
		}
	*/
	return securestorage.RemoveAll(profileName)
}
