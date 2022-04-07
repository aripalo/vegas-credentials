package cache

import (
	"log"
	"time"

	"github.com/alexflint/go-filemutex"
	"github.com/aripalo/vegas-credentials/internal/cache/database"
	"github.com/aripalo/vegas-credentials/internal/cache/encryption"
	"github.com/aripalo/vegas-credentials/internal/config"
)

// NewCache (and its methods) describes the caching mechanism
type NewCache struct {
	db *database.Database
}

// Lock helps with parallel executions (e.g. with Terraform parallelism=n)
// ensuring only a single process at a time can interact with AWS and
// the internal cache – as the BadgerDB requires a filelock:
// https://github.com/dgraph-io/badger/blob/69926151f6532f2fe97a9b11ee9281519c8ec5e6/dir_unix.go#L45
//
// An "unlock function" is returned which removes the directory lock, allowing
// other processes to continue.
func Lock() func() error {
	// as we can't lock the BadgerDB directory (since BadgerDB has its own lock),
	// we use this "lock-control" directory to achieve the desired result
	cachePath := CachePath(config.APP_NAME, "lock-control")
	m, err := filemutex.New(cachePath)
	if err != nil {
		log.Fatalln("Directory did not exist or file could not created")
	}

	m.Lock() // Will block until lock can be acquired

	// return the unlock function which user can call
	return m.Unlock
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
