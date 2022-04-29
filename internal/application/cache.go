package application

import (
	"fmt"

	"github.com/aripalo/vegas-credentials/internal/credentials"
	"github.com/aripalo/vegas-credentials/internal/msg"
	"github.com/aripalo/vegas-credentials/internal/yubikey2/passcache"
)

type CacheFlags struct {
	Password   bool `mapstructure:"password"`
	Credential bool `mapstructure:"credential"`
}

func (app *App) CacheClean(flags CacheFlags) error {

	if flags.Password {
		err := cleanPasswords()
		if err != nil {
			msg.Bail(fmt.Sprintf("error cleaning password cache:%v", err))
		}
		msg.Message.Successln("✅", "Password cache cleaned")
	}

	if flags.Credential {
		err := cleanCredentials()
		if err != nil {
			msg.Bail(fmt.Sprintf("error cleaning credential cache:%v", err))
		}
		msg.Message.Successln("✅", "Credential cache cleaned")
	}

	if !flags.Password && !flags.Credential {
		err := cleanPasswords()
		if err != nil {
			msg.Bail(fmt.Sprintf("error cleaning password cache:%v", err))
		}
		err = cleanCredentials()
		if err != nil {
			msg.Bail(fmt.Sprintf("error cleaning credential cache:%v", err))
		}
		msg.Message.Successln("✅", "Cache cleaned")
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
