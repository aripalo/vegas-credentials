package cmd

import (
	"github.com/aripalo/vegas-credentials/internal/application"
	"github.com/aripalo/vegas-credentials/internal/flag"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration information",
}

var configListCmd = &cobra.Command{
	Use:   "list",
	Short: "List configuration information",
	RunE: func(cmd *cobra.Command, args []string) error {

		g, err := flag.Parse(application.GlobalFlags{}, cmd)
		if err != nil {
			return err
		}

		app := application.New(g)
		return app.ConfigList()
	},
}

var configShowProfileCmd = &cobra.Command{
	Use:   "show-profile",
	Short: "Show resolved profile configuration",
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
		return app.ConfigShowProfile(f)
	},
}
