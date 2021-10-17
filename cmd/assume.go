package cmd

import (
	"log"

	"github.com/aripalo/aws-mfa-credential-process/internal/application/assume"
	"github.com/aripalo/aws-mfa-credential-process/internal/flags"
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
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return app.Config.Load(cmd)
		},
		Run: func(cmd *cobra.Command, args []string) {
			app.Assume()
		},
	}

	flags.DefineAssumeFlags(cmd, app.Config, "profile")

	return cmd

}
