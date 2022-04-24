package passcache

import (
	"fmt"
	"time"
	"vegas3/internal/cache"
	"vegas3/internal/config"
	"vegas3/internal/database"
	"vegas3/internal/encryption"
	"vegas3/internal/msg"
)

type YubikeyPasswordCache struct {
	serial string
	cache  *cache.Cache
}

func New(serial string) *YubikeyPasswordCache {
	return &YubikeyPasswordCache{
		serial: serial,
		cache:  initCache(),
	}
}

// Open new database where to store yubikey password
func initCache() *cache.Cache {
	yubikeyCache := config.YubikeyOathPasswordCacheFile
	msg.Message.Debugln("🔧", fmt.Sprintf("Path: Yubikey OATH Cache: %s", yubikeyCache))
	db, err := database.Open(yubikeyCache, database.DatabaseOptions{})
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

	return ypc.cache.Set(ypc.serial, []byte(encrypted), time.Duration(24*time.Hour))
}

// Removes a password from Cache Database
func (ypc *YubikeyPasswordCache) Delete() error {
	return ypc.cache.Remove(ypc.serial)
}
