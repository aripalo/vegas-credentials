package database

import (
	"time"

	"github.com/dgraph-io/badger/v3"
)

// Database definition
type Database struct {
	db *badger.DB
}

// DatabaseOptions allow providing configuration
type DatabaseOptions struct {
	Logger badger.Logger
}

// Open a database connection
func Open(path string, opts DatabaseOptions) (*Database, error) {
	// Open the Badger database located in the path
	// It will be created if it doesn't exist.
	db, err := badger.Open(
		badger.DefaultOptions(path).WithLogger(opts.Logger).WithLoggingLevel(badger.ERROR),
	)
	if err != nil {
		return nil, err
	}
	d := &Database{db}
	return d, nil
}

// Close database connection
func (d *Database) Close() error {
	return d.db.Close()
}

// Read item from database
func (d *Database) Read(key string) ([]byte, error) {
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

// Write item to database (with TTL)
func (d *Database) Write(key string, value []byte, ttl time.Duration) error {
	return d.db.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry([]byte(key), value).WithTTL(ttl)
		err := txn.SetEntry(e)
		return err
	})
}

// Delete item by key from database
func (d *Database) Delete(key string) error {
	return d.db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})
}

// DeleteByPrefix deletes all keys matching the prefix from database
func (d *Database) DeleteByPrefix(keyPrefix string) error {
	return d.db.DropPrefix([]byte(keyPrefix))
}

// DeleteAll deletes everything from the database
func (d *Database) DeleteAll() error {
	return d.db.DropAll()
}
