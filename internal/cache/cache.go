package cache

import (
	"encoding/json"

	"github.com/aripalo/aws-mfa-credential-process/internal/cache/cachekey"
	"github.com/aripalo/aws-mfa-credential-process/internal/cache/securestorage"
	"github.com/aripalo/aws-mfa-credential-process/internal/data"
)

func Get(d data.Provider) (json.RawMessage, error) {
	c := d.GetConfig()
	p := d.GetProfile()

	cacheKey, err := cachekey.Get(c.Profile, *p)
	if err != nil {
		return nil, err
	}
	cached, err := securestorage.Get(cacheKey)
	return cached, err
}

func Save(d data.Provider, data json.RawMessage) error {
	c := d.GetConfig()
	p := d.GetProfile()

	cacheKey, err := cachekey.Get(c.Profile, *p)
	err = securestorage.Set(cacheKey, data)
	return err
}

// Remove a given configuration from cache
func Remove(d data.Provider) error {
	c := d.GetConfig()
	p := d.GetProfile()

	cacheKey, err := cachekey.Get(c.Profile, *p)
	err = securestorage.Remove(cacheKey)
	return err
}

// RemoveAll the hole cache or all items related to given profile
func RemoveAll(d data.Provider) error {
	c := d.GetConfig()
	/*
		if profileName != "" {
			utils.SafeLogLn(utils.FormatMessage(utils.COLOR_IMPORTANT, "ℹ️  ", "Cache", fmt.Sprintf("Deleting cache for profile \"%s\"", profileName)))
		} else {
			utils.SafeLogLn(utils.FormatMessage(utils.COLOR_IMPORTANT, "ℹ️  ", "Cache", "Deleting all items from cache"))
		}
	*/
	return securestorage.RemoveAll(c.Profile)
}
