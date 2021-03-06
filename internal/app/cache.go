package app

import (
	"fmt"

	"github.com/aripalo/vegas-credentials/internal/credentials"
	"github.com/aripalo/vegas-credentials/internal/msg"
	"github.com/aripalo/vegas-credentials/internal/yubikey/passcache"
)

// CLI flags for this command.
type CacheFlags struct {
	Password   bool `mapstructure:"password"`
	Credential bool `mapstructure:"credential"`
}

// Implementation of "cache clean" CLI command
// without any knowledge of spf13/cobra internals.
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
	return cache.DeleteAll()
}

func cleanCredentials() error {
	cache := credentials.NewCredentialCache()
	return cache.DeleteAll()
}
