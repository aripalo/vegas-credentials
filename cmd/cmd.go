// Package cmd defines all the spf13/cobra commands and their configuration.
// It does not implement the actual application logic; Instead application logic are
// is implemented in app package which contains a method per command.
package cmd

import (
	_ "embed"
	"fmt"

	"github.com/aripalo/vegas-credentials/internal/config"
	"github.com/aripalo/vegas-credentials/internal/msg"

	"github.com/spf13/cobra"
)

//go:embed data/root-short.txt
var rootDescShort string

//go:embed data/root-long.txt
var rootDescLong string

var rootCmd = &cobra.Command{
	Use:     "vegas-credentials",
	Version: config.Version,
	Short:   rootDescShort,
	Long:    rootDescLong,
	PreRun: func(cmd *cobra.Command, args []string) {
		msg.Trace("", fmt.Sprintf("%s cmd init", cmd.Name()))
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		msg.Trace("", fmt.Sprintf("%s cmd done", cmd.Name()))
	},
}

// init is automatically called by spf13/cobra to setup the commands.
func init() {

	rootCmd.PersistentFlags().Bool("no-color", false, "disable both colors and emoji from visible output")
	rootCmd.PersistentFlags().Bool("no-emoji", false, "disable emoji from visible output (but keep colors)")
	rootCmd.PersistentFlags().Bool("no-gui", false, "disable GUI Diaglog Prompt")
	rootCmd.PersistentFlags().Bool("verbose", false, "enable verbose output")

	rootCmd.SetVersionTemplate(config.VersionShortTmpl)
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().Bool("full", false, "display full version information")

	profileFlag := "profile"

	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configListCmd)
	configCmd.AddCommand(configShowProfileCmd)
	configShowProfileCmd.Flags().StringP(profileFlag, "p", "", "aws profile to use from config")
	err := configShowProfileCmd.MarkFlagRequired(profileFlag)
	if err != nil {
		msg.Fatal(err.Error())
	}

	rootCmd.AddCommand(cacheCmd)
	cacheCmd.AddCommand(cacheCleanCmd)
	cacheCleanCmd.Flags().Bool("password", false, "delete yubikey oath application password cache")
	cacheCleanCmd.Flags().Bool("credential", false, "delete temporary session credential cache")

	rootCmd.AddCommand(assumeCmd)

	assumeCmd.Flags().StringP(profileFlag, "p", "", "aws profile to use from config")
	err = assumeCmd.MarkFlagRequired(profileFlag)
	if err != nil {
		msg.Fatal(err.Error())
	}
}

// Execute the root cmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		msg.Fatal(err.Error())
	}
}
