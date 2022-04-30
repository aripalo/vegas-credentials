package application

import (
	_ "embed"
	"fmt"

	"github.com/aripalo/vegas-credentials/internal/assumable"
	"github.com/aripalo/vegas-credentials/internal/config/locations"
	"github.com/aripalo/vegas-credentials/internal/msg"
	"github.com/aripalo/vegas-credentials/internal/tmpl"
)

//go:embed data/config.tmpl
var ConfigTmpl string

//go:embed data/show-profile.tmpl
var ConfigShowProfileTmpl string

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

	return tmpl.Write(app.dest, "config-list", ConfigTmpl, v)
}

func (app *App) ConfigShowProfile(flags AssumeFlags) error {

	opts, err := assumable.New(locations.AwsConfig, flags.Profile)
	if err != nil {
		msg.Bail(fmt.Sprintf("Credentials: Error: %s", err))
	}

	fmt.Println(opts)

	return tmpl.Write(app.dest, "config-show-profile", ConfigShowProfileTmpl, opts)
}
