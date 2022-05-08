package yubikey

import (
	"context"
	"fmt"
	"time"

	"github.com/aripalo/vegas-credentials/internal/cache"
	"github.com/aripalo/vegas-credentials/internal/config/locations"
	"github.com/aripalo/vegas-credentials/internal/msg"
	"github.com/aripalo/vegas-credentials/internal/yubikey/setup"
	"github.com/aripalo/ykmangoath"
)

var cacheLocation string = locations.EnsureWithinDir(locations.CacheDir, "yubikey-oath-access")

type Yubikey struct {
	repo      cache.Repository
	device    string
	account   string
	enableGui bool
	password  string
	GetCode   func(ctx context.Context) (string, error)
}

// Options passed in by the caller
type Options struct {
	Device    string
	Account   string
	EnableGui bool
}

func NewCache() cache.Repository {
	msg.Debug("🔧", fmt.Sprintf("Path: Credentials Cache: %s", cacheLocation))
	return cache.New(cacheLocation)
}

func New(options Options) (Yubikey, error) {
	var err error

	y := Yubikey{
		repo:      NewCache(),
		device:    options.Device,
		account:   options.Account,
		enableGui: options.EnableGui,
	}

	err = setup.Setup(setup.Options(options), &y)

	return y, err
}

// Save password to cache and assign it into the instance.
func (y *Yubikey) SetPassword(password string) error {
	err := y.repo.Write(y.device, []byte(password), time.Duration(12*time.Hour))
	if err != nil {
		return err
	}
	y.password = password
	return nil
}

// Get password from cache, assign it into the instance and finally return it.
func (y *Yubikey) GetPassword() (string, error) {
	result, err := y.repo.Read(y.device)
	if err != nil {
		return "", err
	}
	y.password = string(result)
	return y.password, nil
}

// Remove password from cache and fomr the instance.
func (y *Yubikey) RemovePassword() error {
	err := y.repo.Delete(y.device)
	if err != nil {
		return err
	}
	y.password = ""
	return nil
}

// Code is responsible for querying the TOTP code from Yubikey device.
func (y *Yubikey) Code(ctx context.Context) (string, error) {

	oathAccounts, err := ykmangoath.New(ctx, y.device)
	if err != nil {
		return "", err
	}

	password, err := y.GetPassword()
	if err != nil {
		return "", err
	}

	if password != "" {
		// set the password we already know (after yubikey init)
		err := oathAccounts.SetPassword(password)
		if err != nil {
			return "", err
		}
	}

	return oathAccounts.Code(y.account)
}
