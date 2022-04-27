package passcache

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/aripalo/vegas-credentials/internal/cache"
	"github.com/aripalo/vegas-credentials/internal/config/locations"
	"github.com/aripalo/vegas-credentials/internal/database"
	"github.com/aripalo/vegas-credentials/internal/encryption"
	"github.com/aripalo/vegas-credentials/internal/msg"
)

type YubikeyPasswordCache struct {
	serial string
	cache  *cache.Cache
}

func New(serial string) *YubikeyPasswordCache {
	return &YubikeyPasswordCache{
		serial: serial,
		cache:  InitCache(),
	}
}

const cacheName string = "yubikey-oath-access"

var CacheLocation string = filepath.Join(locations.CacheDir, cacheName)

// Open new database where to store yubikey password
func InitCache() *cache.Cache {
	msg.Message.Debugln("🔧", fmt.Sprintf("Path: Yubikey OATH Cache: %s", CacheLocation))
	db, err := database.Open(CacheLocation, database.DatabaseOptions{})
	if err != nil {
		panic(err)
	}
	return cache.New(db)
}

// Read password (if any) from Cache Database
func (ypc *YubikeyPasswordCache) Read() (string, error) {
	cached, err := ypc.cache.Get(ypc.serial)
	if err != nil {
		return "", err
	}

	decrypted, err := encryption.Decrypt(cached)
	if err != nil {
		return "", err
	}

	return string(decrypted), err
}

// Save password to Cache Database
func (ypc *YubikeyPasswordCache) Write(password string) error {
	encrypted, err := encryption.Encrypt([]byte(password))
	if err != nil {
		return err
	}

	return ypc.cache.Set(ypc.serial, []byte(encrypted), time.Duration(12*time.Hour))
}

// Removes a password from Cache Database
func (ypc *YubikeyPasswordCache) Delete() error {
	return ypc.cache.Remove(ypc.serial)
}
