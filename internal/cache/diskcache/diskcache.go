package diskcache

import (
	"time"

	"github.com/dgraph-io/badger/v3"
)

// DiskCache definition
type DiskCache struct {
	db *badger.DB
}

// Options allow providing configuration
type Options struct {
	Logger badger.Logger
}

// New a disk cache connection
func New(path string, opts Options) (*DiskCache, error) {
	// Open the Badger database located in the path
	// It will be created if it doesn't exist.
	db, err := badger.Open(
		badger.DefaultOptions(path).WithLogger(opts.Logger).WithLoggingLevel(badger.ERROR),
	)
	if err != nil {
		return nil, err
	}
	d := &DiskCache{db}
	return d, nil
}

// Read item from disk cache
func (d *DiskCache) Read(key string) ([]byte, error) {
	var value []byte
	err := d.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		value, err = item.ValueCopy(nil)
		return err
	})
	return value, err
}

// Write item to disk cache (with TTL)
func (d *DiskCache) Write(key string, value []byte, ttl time.Duration) error {
	return d.db.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry([]byte(key), value).WithTTL(ttl)
		err := txn.SetEntry(e)
		return err
	})
}

// Delete item by key from disk cache
func (d *DiskCache) Delete(key string) error {
	return d.db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})
}

// DeleteByPrefix deletes all keys matching the prefix from disk cache
func (d *DiskCache) DeleteByPrefix(keyPrefix string) error {
	return d.db.DropPrefix([]byte(keyPrefix))
}

// DeleteAll deletes everything from the disk cache
func (d *DiskCache) DeleteAll() error {
	return d.db.DropAll()
}

// Close disk cache connection
func (d *DiskCache) Close() error {
	return d.db.Close()
}
