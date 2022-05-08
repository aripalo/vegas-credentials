package cache

import (
	"fmt"
	"time"

	"github.com/aripalo/vegas-credentials/internal/cache/diskcache"
	"github.com/aripalo/vegas-credentials/internal/cache/encryption"
	"github.com/aripalo/vegas-credentials/internal/msg"
)

// Repository interface defines the methods that any cache implementation
// must implement.
type Repository interface {
	Write(key string, value []byte, ttl time.Duration) error
	Read(key string) ([]byte, error)
	Delete(key string) error
	DeleteByPrefix(keyPrefix string) error
	DeleteAll() error
	Close() error
}

type Cache struct {
	repo Repository
}

func New(cachePath string) Repository {
	repo, err := diskcache.New(cachePath, diskcache.Options{})
	if err != nil {
		msg.Fatal(fmt.Sprintf("Configuration Error: %s", err))
	}
	return &Cache{repo: repo}
}

// Write value to cache
func (c *Cache) Write(key string, data []byte, ttl time.Duration) error {
	encrypted, err := encryption.Encrypt(data)
	if err != nil {
		return err
	}

	err = c.repo.Write(key, encrypted, ttl)
	if err != nil {
		return err
	}

	return nil
}

// Read value from cache
func (c *Cache) Read(key string) ([]byte, error) {
	cached, err := c.repo.Read(key)
	if err != nil {
		return nil, err
	}

	decrypted, err := encryption.Decrypt(cached)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}

// Delete value from cache
func (c *Cache) Delete(key string) error {
	err := c.repo.Delete(key)
	if err != nil {
		return err
	}
	return nil
}

// DeleteByPrefix clears all values with key prefix from cache
func (c *Cache) DeleteByPrefix(keyPrefix string) error {
	return c.repo.DeleteByPrefix(keyPrefix)
}

// DeleteAll clears the whole cache
func (c *Cache) DeleteAll() error {
	return c.repo.DeleteAll()
}

// Close closes cache database connections
func (c *Cache) Close() error {
	return c.repo.Close()
}
