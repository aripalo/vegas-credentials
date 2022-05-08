package app

import (
	_ "embed"
	"fmt"

	"github.com/aripalo/vegas-credentials/internal/assumecfg"
	"github.com/aripalo/vegas-credentials/internal/config/locations"
	"github.com/aripalo/vegas-credentials/internal/msg"
	"github.com/aripalo/vegas-credentials/internal/tmpl"
)

//go:embed data/config.tmpl
var ConfigTmpl string

//go:embed data/show-profile.tmpl
var ConfigShowProfileTmpl string

// Data used in templating.
type configData struct {
	AwsConfig      string
	YkmanPath      string
	CacheDir       string
	StateDir       string
	ExecDir        string
	YkmanInstalled bool
}

// Implementation of "config list" CLI command
// without any knowledge of spf13/cobra internals.
func (a *App) ConfigList() error {
	v := configData{
		AwsConfig: locations.AwsConfig,
		YkmanPath: locations.YkmanPath,
		CacheDir:  locations.CacheDir,
		StateDir:  locations.StateDir,
		ExecDir:   locations.ExecDir,
	}

	return tmpl.Write(a.dest, "config-list", ConfigTmpl, v)
}

// Implementation of "config show-profile" CLI command
// without any knowledge of spf13/cobra internals.
func (a *App) ConfigShowProfile(flags AssumeFlags) error {

	cfg, err := assumecfg.New(locations.AwsConfig, flags.Profile)
	if err != nil {
		msg.Fatal(fmt.Sprintf("Credentials: Error: %s", err))
	}

	fmt.Println(cfg)

	return tmpl.Write(a.dest, "config-show-profile", ConfigShowProfileTmpl, cfg)
}
