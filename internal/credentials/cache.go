package credentials

import (
	"time"

	"github.com/aripalo/vegas-credentials/internal/assumable"
	"github.com/aripalo/vegas-credentials/internal/cache"
	"github.com/aripalo/vegas-credentials/internal/utils"
)

func resolveKey(options assumable.Assumable) (string, error) {
	checksum, err := utils.CalculateChecksum(options)
	if err != nil {
		return "", err
	}
	key := cache.Key(options.ProfileName, checksum)
	return key, nil
}

// saveToCache saves response to cache in cache database
func (c *Credentials) saveToCache() error {
	data, err := c.Serialize()
	if err != nil {
		return err
	}

	key, err := resolveKey(c.options)
	if err != nil {
		return err
	}

	now := time.Now()
	ttl := c.Expiration.Sub(now)

	err = c.cache.Set(key, data, ttl)
	if err != nil {
		return err
	}

	//logger.Debugln(a, "ℹ️ ", "Credentials", "Saved to cache")

	return nil
}

// ReadFromCache gets the cached response from cache database
func (c *Credentials) readFromCache() error {
	key, err := resolveKey(c.options)
	if err != nil {
		return err
	}

	data, err := c.cache.Get(key)
	if err != nil {
		return err
	}

	//logger.DebugJSON(a, "🔧 ", "Cached Credentials", data)

	err = c.Deserialize(data)
	if err != nil {
		return err
	}

	return nil
}

// DeleteFromCache deletes the cached response cache database
func (c *Credentials) deleteFromCache() error {
	key, err := resolveKey(c.options)
	if err != nil {
		return err
	}

	return c.cache.Remove(key)
}
