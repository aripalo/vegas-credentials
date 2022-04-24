package application

import (
	"io"
	"os"
)

type GlobalFlags struct {
	NoColor bool `mapstructure:"no-color"`
	NoEmoji bool `mapstructure:"no-emoji"`
	NoGui   bool `mapstructure:"no-gui"`
	Verbose bool `mapstructure:"verbose"`
}

type App struct {
	GlobalFlags
	dest io.Writer
}

// Instantiate a new instance of App with defaults
func New(globalFlags GlobalFlags) App {
	return App{
		GlobalFlags: globalFlags,
		dest:        os.Stdout,
	}
}
