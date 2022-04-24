package application

import "fmt"

type CacheFlags struct {
	Password   bool `mapstructure:"password"`
	Credential bool `mapstructure:"credential"`
}

func (app *App) CacheClean(flags CacheFlags) error {

	if flags.Password {
		fmt.Println("TODO: DELETE PASSWORDS")
	}

	if flags.Credential {
		fmt.Println("TODO: DELETE CREDENTIALS")
	}

	if !flags.Password && !flags.Credential {
		fmt.Println("TODO: DELETE ALL")
	}

	return nil
}
