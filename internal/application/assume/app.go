package assume

import (
	"io"

	"github.com/aripalo/aws-mfa-credential-process/internal/application/assume/awscreds"
	"github.com/aripalo/aws-mfa-credential-process/internal/cache"
	"github.com/aripalo/aws-mfa-credential-process/internal/cache/securestorage"
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
func (a *App) Assume(commandName string, version string) {

	logger.PrintBanner(a, commandName, version)

	logger.DebugJSON(a, "🔧 ", "Config", a.Config)

	err := a.Profile.Load(a.Config) // TODO could this use a?
	if err != nil {
		panic(err)
	}

	securestorage.Init(a.Config.DisableDialog)

	logger.DebugJSON(a, "🔧 ", "Profile", a.Profile)

	cached, cacheErr := cache.Get(a)
	if cacheErr == nil {
		result, err := utils.PrettyJSON(cached)
		if err == nil {
			awscreds.OutputToAwsCredentialProcess(result)
		}
	} else {
		logger.Infoln(a, "ℹ️ ", "Cache", cacheErr.Error())
	}

	if cached == nil {
		err = awscreds.CredentialProcess(a)
		if err != nil {
			panic(err)
		}
	}

}
