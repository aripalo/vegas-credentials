package passcache

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/aripalo/vegas-credentials/internal/cache"
	"github.com/aripalo/vegas-credentials/internal/config/locations"
	"github.com/aripalo/vegas-credentials/internal/msg"
)

type YubikeyPasswordCache struct {
	serial string
	repo   cache.Repository
}

func New(serial string) *YubikeyPasswordCache {
	return &YubikeyPasswordCache{
		serial: serial,
		repo:   InitCache(),
	}
}

const cacheName string = "yubikey-oath-access"

var CacheLocation string = filepath.Join(locations.CacheDir, cacheName)

// Open new database where to store yubikey password
func InitCache() cache.Repository {
	msg.Debug("🔧", fmt.Sprintf("Path: Yubikey OATH Cache: %s", CacheLocation))
	return cache.New(CacheLocation)
}

// Read password (if any) from Cache Database
func (ypc *YubikeyPasswordCache) Read() (string, error) {
	cached, err := ypc.repo.Read(ypc.serial)
	if err != nil {
		return "", err
	}
	return string(cached), err
}

// Save password to Cache Database
func (ypc *YubikeyPasswordCache) Write(password string) error {
	return ypc.repo.Write(ypc.serial, []byte(password), time.Duration(12*time.Hour))
}

// Removes a password from Cache Database
func (ypc *YubikeyPasswordCache) Delete() error {
	return ypc.repo.Delete(ypc.serial)
}
