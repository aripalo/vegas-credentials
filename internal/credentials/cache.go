package credentials

import (
	"fmt"
	"time"

	"github.com/aripalo/vegas-credentials/internal/cache"
	"github.com/aripalo/vegas-credentials/internal/config/locations"
	"github.com/aripalo/vegas-credentials/internal/msg"
)

var cacheLocation string = locations.EnsureWithinDir(locations.CacheDir, "session-cache")

func NewCredentialCache() cache.Repository {
	msg.Debug("🔧", fmt.Sprintf("Credentials: Cache: %s", cacheLocation))
	return cache.New(cacheLocation)
}

func resolveKey(profileName string, checksum string) (string, error) {
	key := cache.Key(profileName, checksum)
	return key, nil
}

// saveToCache saves response to cache in cache database
func (c *Credentials) saveToCache() error {
	data, err := c.Serialize()
	if err != nil {
		return err
	}

	key, err := resolveKey(c.cfg.ProfileName, c.cfg.Checksum)
	if err != nil {
		return err
	}

	now := time.Now()
	ttl := c.Expiration.Sub(now)

	err = c.repo.Write(key, data, ttl)
	if err != nil {
		return err
	}

	//logger.Debugln(a, "ℹ️ ", "Credentials", "Saved to cache")

	return nil
}

// ReadFromCache gets the cached response from cache database
func (c *Credentials) readFromCache() error {
	key, err := resolveKey(c.cfg.ProfileName, c.cfg.Checksum)
	if err != nil {
		return err
	}

	data, err := c.repo.Read(key)
	if err != nil {
		return err
	}

	err = c.Deserialize(data)
	if err != nil {
		return err
	}

	return nil
}
