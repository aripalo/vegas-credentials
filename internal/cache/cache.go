package cache

import (
	"time"

	"github.com/aripalo/vegas-credentials/internal/cache/database"
	"github.com/aripalo/vegas-credentials/internal/cache/encryption"
)

// NewCache (and its methods) describes the caching mechanism
type NewCache struct {
	db *database.Database
}

// New instantiates a cache
func New(cacheName string) *NewCache {
	cachePath := CachePath(cacheName, "cachedb")
	db, err := database.Open(cachePath, database.DatabaseOptions{})
	if err != nil {
		panic(err)
	}
	return &NewCache{db}
}

// Set value to cache
func (n *NewCache) Set(key string, data []byte, ttl time.Duration) error {
	encrypted, err := encryption.Encrypt(data)
	if err != nil {
		return err
	}

	err = n.db.Write(key, encrypted, ttl)
	if err != nil {
		return err
	}

	return nil
}

// Get value from cache
func (n *NewCache) Get(key string) ([]byte, error) {
	cached, err := n.db.Read(key)
	if err != nil {
		return nil, err
	}

	decrypted, err := encryption.Decrypt(cached)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}

// Remove value from cache
func (n *NewCache) Remove(key string) error {
	err := n.db.Delete(key)
	if err != nil {
		return err
	}
	return nil
}

// RemoveByPrefix clears all values with key prefix from cache
func (n *NewCache) RemoveByPrefix(keyPrefix string) error {
	return n.db.DeleteByPrefix(keyPrefix)
}

// RemoveAll clears the whole cache
func (n *NewCache) RemoveAll() error {
	return n.db.DeleteAll()
}

// Disconnect closes cache database connections
func (n *NewCache) Disconnect() error {
	return n.db.Close()
}
