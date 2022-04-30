package cmd

import (
	"fmt"

	"github.com/aripalo/vegas-credentials/internal/application"
	"github.com/aripalo/vegas-credentials/internal/flag"
	"github.com/aripalo/vegas-credentials/internal/logger"

	"github.com/spf13/cobra"
)

var assumeCmd = &cobra.Command{
	Use:   "assume",
	Short: "Assume Temporary Session Credentials",
	PreRun: func(cmd *cobra.Command, args []string) {
		logger.Trace(fmt.Sprintf("%s cmd init", cmd.Name()))
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		logger.Trace(fmt.Sprintf("%s cmd done", cmd.Name()))
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		g, err := flag.Parse(application.GlobalFlags{}, cmd)
		if err != nil {
			return err
		}

		f, err := flag.Parse(application.AssumeFlags{}, cmd)
		if err != nil {
			return err
		}

		app := application.New(g)
		return app.Assume(f)
	},
}
