package cmd

import (
	"github.com/aripalo/vegas-credentials/internal/config"
	"github.com/spf13/cobra"
)

var version string = "development"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     config.APP_NAME,
	Short:   config.APP_DESCRIPTION_SHORT,
	Long:    config.APP_DESCRIPTION_LONG,
	Version: version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {

	rootCmd.PersistentFlags().Bool(
		config.Defaults.NoColor.Name,
		config.Defaults.NoColor.Value,
		config.Defaults.NoColor.Usage,
	)

	rootCmd.PersistentFlags().Bool(
		config.Defaults.Verbose.Name,
		config.Defaults.Verbose.Value,
		config.Defaults.Verbose.Usage,
	)

	rootCmd.PersistentFlags().Bool(
		config.Defaults.Debug.Name,
		config.Defaults.Debug.Value,
		config.Defaults.Debug.Usage,
	)

}
