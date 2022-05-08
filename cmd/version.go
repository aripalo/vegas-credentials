package cmd

import (
	"fmt"

	"github.com/aripalo/vegas-credentials/internal/app"
	"github.com/aripalo/vegas-credentials/internal/app/flagparser"
	"github.com/aripalo/vegas-credentials/internal/msg"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version information",
	PreRun: func(cmd *cobra.Command, args []string) {
		msg.Trace("", fmt.Sprintf("%s cmd init", cmd.Name()))
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		msg.Trace("", fmt.Sprintf("%s cmd done", cmd.Name()))
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		g, err := flagparser.Parse(app.GlobalFlags{}, cmd)
		if err != nil {
			return err
		}

		f, err := flagparser.Parse(app.VersionFlags{}, cmd)
		if err != nil {
			return err
		}

		a := app.New(g)
		return a.Version(f)
	},
}
