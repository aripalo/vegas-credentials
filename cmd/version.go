package cmd

import (
	"github.com/aripalo/vegas-credentials/internal/application"
	"github.com/aripalo/vegas-credentials/internal/utils"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version information",
	RunE: func(cmd *cobra.Command, args []string) error {

		g, err := utils.ParseFlags(application.GlobalFlags{}, cmd)
		if err != nil {
			return err
		}

		f, err := utils.ParseFlags(application.VersionFlags{}, cmd)
		if err != nil {
			return err
		}

		app := application.New(g)
		return app.Version(f)
	},
}
