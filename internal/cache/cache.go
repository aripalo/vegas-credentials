package cache

import (
	"encoding/json"

	"github.com/aripalo/aws-mfa-credential-process/internal/cachekey"
	"github.com/aripalo/aws-mfa-credential-process/internal/profile"
	"github.com/aripalo/aws-mfa-credential-process/internal/securestorage"
)

func Get(profileName string, config profile.Profile) (json.RawMessage, error) {
	cacheKey, err := cachekey.Get(profileName, config)
	if err != nil {
		return nil, err
	}
	cached, err := securestorage.Get(cacheKey)
	return cached, err
}

func Save(profileName string, config profile.Profile, data json.RawMessage) error {
	cacheKey, err := cachekey.Get(profileName, config)
	err = securestorage.Set(cacheKey, data)
	return err
}

func Remove(profileName string, config profile.Profile) error {
	cacheKey, err := cachekey.Get(profileName, config)
	err = securestorage.Remove(cacheKey)
	return err
}
