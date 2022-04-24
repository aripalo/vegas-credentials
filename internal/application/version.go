package application

import (
	"runtime"

	"github.com/aripalo/vegas-credentials/internal/config"
	"github.com/aripalo/vegas-credentials/internal/utils"
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
		return utils.PrintTemplate(a.dest, "version-long", config.VersionLongTmpl, v)
	}
	return utils.PrintTemplate(a.dest, "version-short", config.VersionShortTmpl, v)
}
