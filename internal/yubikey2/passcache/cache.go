package passcache

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/aripalo/vegas-credentials/internal/cache"
	"github.com/aripalo/vegas-credentials/internal/config/locations"
	"github.com/aripalo/vegas-credentials/internal/interfaces"
	"github.com/aripalo/vegas-credentials/internal/msg"
)

type YubikeyPasswordCache struct {
	serial string
	cache  interfaces.Cache
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
	return cache.New(CacheLocation)
}

// Read password (if any) from Cache Database
func (ypc *YubikeyPasswordCache) Read() (string, error) {
	cached, err := ypc.cache.Get(ypc.serial)
	if err != nil {
		return "", err
	}
	return string(cached), err
}

// Save password to Cache Database
func (ypc *YubikeyPasswordCache) Write(password string) error {
	return ypc.cache.Set(ypc.serial, []byte(password), time.Duration(12*time.Hour))
}

// Removes a password from Cache Database
func (ypc *YubikeyPasswordCache) Delete() error {
	return ypc.cache.Remove(ypc.serial)
}
