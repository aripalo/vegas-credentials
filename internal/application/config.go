package application

import (
	_ "embed"

	"github.com/aripalo/vegas-credentials/internal/config/locations"
	"github.com/aripalo/vegas-credentials/internal/tmpl"
)

//go:embed data/config.tmpl
var ConfigTmpl string

type configData struct {
	AwsConfig      string
	YkmanPath      string
	CacheDir       string
	StateDir       string
	ExecDir        string
	YkmanInstalled bool
}

func (app *App) ConfigList() error {
	v := configData{
		AwsConfig: locations.AwsConfig,
		YkmanPath: locations.YkmanPath,
		CacheDir:  locations.CacheDir,
		StateDir:  locations.StateDir,
		ExecDir:   locations.ExecDir,
	}

	return tmpl.Write(app.dest, "config", ConfigTmpl, v)
}
