package app

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
