package application

import (
	"demo/internal/locations"
	"demo/internal/utils"
	_ "embed"
	"os/exec"
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
