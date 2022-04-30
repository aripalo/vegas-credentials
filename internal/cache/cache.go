package cache

import (
	"fmt"
	"time"

	"github.com/aripalo/vegas-credentials/internal/cache/database"
	"github.com/aripalo/vegas-credentials/internal/cache/encryption"
	"github.com/aripalo/vegas-credentials/internal/msg"
)

// Cache (and its methods) describes the caching mechanism
type Cache struct {
	db databaseConnection
}

// Internal interface to describe methods found from BadgerDB.
// Allows switching the internal implementation if required later
// and also useful for testing.
type databaseConnection interface {
	Write(key string, value []byte, ttl time.Duration) error
	Read(key string) ([]byte, error)
	Delete(key string) error
	DeleteByPrefix(keyPrefix string) error
	DeleteAll() error
	Close() error
}

func New(databasePath string) *Cache {
	db, err := database.Open(databasePath, database.DatabaseOptions{})
	if err != nil {
		msg.Fatal(fmt.Sprintf("Configuration Error: %s", err))
	}
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
