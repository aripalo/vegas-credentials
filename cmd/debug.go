package cmd

import (
	"log"

	"github.com/aripalo/aws-mfa-credential-process/internal/application/debug"
	"github.com/aripalo/aws-mfa-credential-process/internal/flags"
	"github.com/spf13/cobra"
)

func init() {

	app, err := debug.New()
	if err != nil {
		log.Fatal(err)
	}
	cmd := buildDebugCommand(app)

	rootCmd.AddCommand(cmd)

}

func buildDebugCommand(app *debug.App) *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "debug",
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
			app.Debug()
		},
	}

	flags.DefineAssumeFlags(cmd, app.Config, "profile")

	return cmd

}
