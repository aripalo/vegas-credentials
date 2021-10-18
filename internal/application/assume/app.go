package assume

import (
	"fmt"
	"io"
	"os"

	"github.com/aripalo/aws-mfa-credential-process/internal/application/assume/awscreds"
	"github.com/aripalo/aws-mfa-credential-process/internal/config"
	"github.com/aripalo/aws-mfa-credential-process/internal/logger"
	"github.com/aripalo/aws-mfa-credential-process/internal/profile"
	"github.com/aripalo/aws-mfa-credential-process/internal/utils"
)

// App declaration
type App struct {
	WriteStream io.Writer
	Config      *config.Config
	Profile     *profile.Profile
}

// GetWriteStream implements data.Provider method
func (a *App) GetWriteStream() io.Writer {
	return a.WriteStream
}

// GetConfig implements data.Provider method
func (a *App) GetConfig() *config.Config {
	return a.Config
}

// GetProfile implements data.Provider method
func (a *App) GetProfile() *profile.Profile {
	return a.Profile
}

// New instantiates the App
func New() (*App, error) {
	a := &App{
		WriteStream: logger.GetSafeWriter(),
		Config:      &config.Config{},
		Profile:     &profile.Profile{},
	}
	return a, nil
}

// Assume defines the command attached to cobra
func (a *App) Assume() {

	if a.Config.Debug {
		fmt.Fprintf(os.Stderr, "CONFIG:\n%s\n", utils.PrettyJSON(a.Config))
	}

	err := a.Profile.Load(a.Config) // TODO could this use a?
	if err != nil {
		panic(err)
	}

	if a.Config.Debug {
		fmt.Fprintf(os.Stderr, "PROFILE:\n%s\n", utils.PrettyJSON(a.Profile))
	}

	var credentialprocess *awscreds.CredentialProcess
	credentialprocess, err = credentialprocess.New(a)
	if err != nil {
		// TODO log
		panic(err)
	}
	err = credentialprocess.Get()
	if err != nil {
		// TODO log
		panic(err)
	}

	err = credentialprocess.Print()
	if err != nil {
		// TODO log
		panic(err)
	}

}
