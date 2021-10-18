package cmd

import (
	"log"

	"github.com/aripalo/aws-mfa-credential-process/internal/application/assume"
	"github.com/aripalo/aws-mfa-credential-process/internal/config"
	"github.com/spf13/cobra"
)

func init() {

	app, err := assume.New()
	if err != nil {
		log.Fatal(err)
	}
	cmd := buildAssumeCommand(app)

	rootCmd.AddCommand(cmd)

}

func buildAssumeCommand(app *assume.App) *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "assume",
		Short: config.ASSUME_DESCRIPTION_SHORT,
		Long:  config.ASSUME_DESCRIPTION_LONG,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return app.Config.Load(cmd)
		},
		Run: func(cmd *cobra.Command, args []string) {
			app.Assume(cmd.CalledAs(), cmd.Parent().Version)
		},
	}

	cmd.Flags().String(
		config.Defaults.Profile.Name,
		config.Defaults.Profile.Value,
		config.Defaults.Profile.Usage,
	)
	cmd.MarkFlagRequired(config.Defaults.Profile.Name)

	cmd.Flags().Bool(
		config.Defaults.HideArns.Name,
		config.Defaults.HideArns.Value,
		config.Defaults.HideArns.Usage,
	)
	cmd.Flags().Bool(
		config.Defaults.DisableDialog.Name,
		config.Defaults.DisableDialog.Value,
		config.Defaults.DisableDialog.Usage,
	)
	cmd.Flags().Bool(
		config.Defaults.DisableRefresh.Name,
		config.Defaults.DisableRefresh.Value,
		config.Defaults.DisableRefresh.Usage,
	)

	return cmd

}
