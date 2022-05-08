// Package app implements a common application struct with global configuration
// and a method per each CLI command.
package app

import (
	"fmt"
	"io"
	"os"

	"github.com/aripalo/vegas-credentials/internal/msg"
)

// GlobalFlags describes all the CLI flags applied to all commands.
type GlobalFlags struct {
	NoColor bool `mapstructure:"no-color"`
	NoEmoji bool `mapstructure:"no-emoji"`
	NoGui   bool `mapstructure:"no-gui"`
	Verbose bool `mapstructure:"verbose"`
}

// App describes the global application configuration.
type App struct {
	GlobalFlags
	dest io.Writer
}

// Instantiate a new instance of App with defaults.
func New(globalFlags GlobalFlags) App {

	msg.Init(msg.Options{
		VerboseMode: globalFlags.Verbose,
		ColorMode:   !globalFlags.NoColor,
		EmojiMode:   !globalFlags.NoEmoji,
	})

	msg.Trace("", fmt.Sprintf(
		"app configuration verbose=%v noColor=%v noEmoji=%v noGui=%v",
		globalFlags.Verbose,
		globalFlags.NoColor,
		globalFlags.NoEmoji,
		globalFlags.NoGui,
	))

	return App{
		GlobalFlags: globalFlags,
		dest:        os.Stdout,
	}
}
