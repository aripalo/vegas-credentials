package cache

import (
	"encoding/json"

	"github.com/aripalo/aws-mfa-credential-process/internal/cache/encryption"
	"github.com/aripalo/aws-mfa-credential-process/internal/cache/filecache"
	"github.com/aripalo/aws-mfa-credential-process/internal/profile"
)

func Get(profileName string, config profile.Profile) (json.RawMessage, error) {
	cached, err := filecache.Get(profileName, config)
	if err != nil {
		return nil, err
	}
	decrypted, err := encryption.Decrypt(cached)
	return decrypted, err
}

func Save(profileName string, config profile.Profile, data json.RawMessage) error {
	encrypted, err := encryption.Encrypt(data)
	if err != nil {
		return err
	}
	err = filecache.Save(profileName, config, encrypted)
	return err
}

func Remove(profileName string, config profile.Profile) error {
	return filecache.Remove(profileName, config)
}
