package application

import (
	_ "embed"

	"github.com/aripalo/vegas-credentials/internal/locations"
	"github.com/aripalo/vegas-credentials/internal/utils"
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

	return utils.PrintTemplate(app.dest, "config", ConfigTmpl, v)
}
