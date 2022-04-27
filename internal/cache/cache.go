package cache

import (
	"time"

	"github.com/aripalo/vegas-credentials/internal/encryption"
)

// Cache (and its methods) describes the caching mechanism
type Cache struct {
	db Database
}

type Database interface {
	Write(key string, value []byte, ttl time.Duration) error
	Read(key string) ([]byte, error)
	Delete(key string) error
	DeleteByPrefix(keyPrefix string) error
	DeleteAll() error
	Close() error
}

// New instantiates a cache
func New(db Database) *Cache {
	return &Cache{db}
}

// Set value to cache
func (c *Cache) Set(key string, data []byte, ttl time.Duration) error {
	encrypted, err := encryption.Encrypt(data)
	if err != nil {
		return err
	}

	err = c.db.Write(key, encrypted, ttl)
	if err != nil {
		return err
	}

	return nil
}

// Get value from cache
func (c *Cache) Get(key string) ([]byte, error) {
	cached, err := c.db.Read(key)
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
func (c *Cache) Remove(key string) error {
	err := c.db.Delete(key)
	if err != nil {
		return err
	}
	return nil
}

// RemoveByPrefix clears all values with key prefix from cache
func (c *Cache) RemoveByPrefix(keyPrefix string) error {
	return c.db.DeleteByPrefix(keyPrefix)
}

// RemoveAll clears the whole cache
func (c *Cache) RemoveAll() error {
	return c.db.DeleteAll()
}

// Disconnect closes cache database connections
func (c *Cache) Disconnect() error {
	return c.db.Close()
}
