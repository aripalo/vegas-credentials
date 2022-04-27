package credentials

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/aripalo/vegas-credentials/internal/cache"
	"github.com/aripalo/vegas-credentials/internal/database"
	"github.com/aripalo/vegas-credentials/internal/locations"
	"github.com/aripalo/vegas-credentials/internal/utils"
)

var cacheLocation string = filepath.Join(locations.CacheDir, "session-cache")

func NewCredentialCache() *cache.Cache {
	db, err := database.Open(cacheLocation, database.DatabaseOptions{})
	if err != nil {
		utils.Bail(fmt.Sprintf("Configuration Error: %s", err))
	}
	return cache.New(db)
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

	key, err := resolveKey(c.opts.Name, c.opts.Checksum)
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
	key, err := resolveKey(c.opts.Name, c.opts.Checksum)
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
	key, err := resolveKey(c.opts.Name, c.opts.Checksum)
	if err != nil {
		return err
	}

	return c.cache.Remove(key)
}
