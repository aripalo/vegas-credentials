package assume

import (
	"io"
	"time"

	"github.com/aripalo/aws-mfa-credential-process/internal/config"
	"github.com/aripalo/aws-mfa-credential-process/internal/logger"
	"github.com/aripalo/aws-mfa-credential-process/internal/profile"
	"github.com/spf13/cobra"
)

// App declaration
type App struct {
	WriteStream io.Writer
	Config      *config.Config
	Profile     *profile.Profile
	command     string
	version     string
	startedAt   time.Time
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
		startedAt:   time.Now(),
	}
	return a, nil
}

// PreRunE is responsible for loading in configurations & init code etc and is the only method that directly depends on Cobra
func (app *App) PreRunE(cmd *cobra.Command) error {
	var err error
	err = app.Config.Load(cmd)
	if err != nil {
		return err
	}
	err = app.Profile.Load(app.Config)
	if err != nil {
		return err
	}
	app.command = cmd.CalledAs()
	app.version = cmd.Parent().Version

	logger.PrintBanner(app, app.command, app.version)
	logger.DebugJSON(app, "🔧 ", "Config", app.Config)
	logger.DebugJSON(app, "🔧 ", "Profile", app.Profile)

	return nil
}

// Run executes the cobra command (but does not directly depend on cobra)
func (app *App) Run() {

	var err error

	err = getCredentials(app)

	if err != nil {
		panic(err)
	}

}

// PostRunE executes after everything
func (app *App) PostRunE() error {
	if app.Config.Debug {
		logger.Debugf(app, "⏱ ", "Duration", "%s\n", time.Now().Sub(app.startedAt))
		logger.PrintRuler(app, "-")
	}
	return nil
}
