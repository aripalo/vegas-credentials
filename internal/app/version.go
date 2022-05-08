package app

import (
	"runtime"

	"github.com/aripalo/vegas-credentials/internal/config"
	"github.com/aripalo/vegas-credentials/internal/tmpl"
)

// Data used in templating.
type versionData struct {
	Version string
	Name    string
	Goos    string
	Goarch  string
}

// CLI flags for this command.
type VersionFlags struct {
	Full bool `mapstructure:"full"`
}

// Implementation of "version" CLI command
// without any knowledge of spf13/cobra internals.
func (a *App) Version(flags VersionFlags) error {
	v := versionData{
		Version: config.Version,
		Name:    config.AppName,
		Goos:    runtime.GOOS,
		Goarch:  runtime.GOARCH,
	}

	if flags.Full {
		return tmpl.Write(a.dest, "version-long", config.VersionLongTmpl, v)
	}
	return tmpl.Write(a.dest, "version-short", config.VersionShortTmpl, v)
}
