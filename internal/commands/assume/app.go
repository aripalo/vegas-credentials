package assume

import (
	"io"
	"time"

	"github.com/aripalo/vegas-credentials/internal/cache"
	"github.com/aripalo/vegas-credentials/internal/config"
	"github.com/aripalo/vegas-credentials/internal/logger"
	"github.com/aripalo/vegas-credentials/internal/profile"
	"github.com/spf13/cobra"
)

// App declaration
type App struct {
	WriteStream io.Writer
	Config      *config.Flags
	Profile     *profile.Profile
	command     string
	version     string
	startedAt   time.Time
}

// GetDestination implements interfaces.AssumeCredentialProcess method
func (a *App) GetDestination() io.Writer {
	return a.WriteStream
}

// GetFlags implements interfaces.AssumeCredentialProcess method
func (a *App) GetFlags() *config.Flags {
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
		Config:      &config.Flags{},
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
	p, err := profile.New(app.Config.Profile)
	if err != nil {
		return err
	}

	app.Profile = p

	app.command = cmd.CalledAs()
	app.version = cmd.Parent().Version

	logger.PrintBanner(app, app.command, app.version)
	logger.DebugJSON(app, "🔧 ", "Config", app.Config)
	logger.DebugJSON(app, "🔧 ", "Profile", app.Profile)

	return nil
}

// Run executes the cobra command (but does not directly depend on cobra)
func (app *App) Run() {

	unlock := cache.Lock()

	err := getCredentials(app)

	unlockErr := unlock()
	if unlockErr != nil {
		panic(unlockErr) // TODO handle better
	}

	if err != nil {
		panic(err) // TODO handle better
	}

}

// PostRunE executes after everything
func (app *App) PostRunE() error {
	if app.Config.Debug {
		logger.Debugf(app, "⏱ ", "Duration", "%s\n", time.Since(app.startedAt))
		logger.PrintRuler(app, "-")
	}
	return nil
}
