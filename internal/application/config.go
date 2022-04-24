package application

import (
	_ "embed"
	"os/exec"

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

func (a *App) ConfigList() error {
	v := configData{
		AwsConfig: locations.AwsConfig,
		YkmanPath: getYkmanPath(),
		CacheDir:  locations.CacheDir,
		StateDir:  locations.StateDir,
		ExecDir:   locations.ExecDir,
	}

	return utils.PrintTemplate(a.dest, "config", ConfigTmpl, v)
}

func getYkmanPath() string {
	if ykmanPath, err := exec.LookPath("ykman"); err != nil {
		return ""
	} else {
		return ykmanPath
	}
}
