package app

import (
	"fmt"

	"github.com/aripalo/vegas-credentials/internal/credentials"
	"github.com/aripalo/vegas-credentials/internal/msg"
	"github.com/aripalo/vegas-credentials/internal/yubikey/passcache"
)

type CacheFlags struct {
	Password   bool `mapstructure:"password"`
	Credential bool `mapstructure:"credential"`
}

func (a *App) CacheClean(flags CacheFlags) error {

	if flags.Password {
		err := cleanPasswords()
		if err != nil {
			msg.Fatal(fmt.Sprintf("error cleaning password cache:%v", err))
		}
		msg.Success("✅", "Password cache cleaned")
	}

	if flags.Credential {
		err := cleanCredentials()
		if err != nil {
			msg.Fatal(fmt.Sprintf("error cleaning credential cache:%v", err))
		}
		msg.Success("✅", "Credential cache cleaned")
	}

	if !flags.Password && !flags.Credential {
		err := cleanPasswords()
		if err != nil {
			msg.Fatal(fmt.Sprintf("error cleaning password cache:%v", err))
		}
		err = cleanCredentials()
		if err != nil {
			msg.Fatal(fmt.Sprintf("error cleaning credential cache:%v", err))
		}
		msg.Success("✅", "Cache cleaned")
	}

	return nil
}

func cleanPasswords() error {
	cache := passcache.InitCache()
	return cache.RemoveAll()
}

func cleanCredentials() error {
	cache := credentials.NewCredentialCache()
	return cache.RemoveAll()
}
