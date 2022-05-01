package application

import (
	"runtime"

	"github.com/aripalo/vegas-credentials/internal/config"
	"github.com/aripalo/vegas-credentials/internal/tmpl"
)

type versionData struct {
	Version string
	Name    string
	Goos    string
	Goarch  string
}

type VersionFlags struct {
	Full bool `mapstructure:"full"`
}

func (app *App) Version(flags VersionFlags) error {
	v := versionData{
		Version: config.Version,
		Name:    config.AppName,
		Goos:    runtime.GOOS,
		Goarch:  runtime.GOARCH,
	}

	if flags.Full {
		return tmpl.Write(app.dest, "version-long", config.VersionLongTmpl, v)
	}
	return tmpl.Write(app.dest, "version-short", config.VersionShortTmpl, v)
}
