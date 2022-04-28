package cmd

import (
	"fmt"
	"os"

	"github.com/aripalo/vegas-credentials/internal/application"
	"github.com/aripalo/vegas-credentials/internal/config"
	"github.com/aripalo/vegas-credentials/internal/msg"
	"github.com/aripalo/vegas-credentials/internal/utils"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "vegas-credentials",
	Version: config.Version,
	Short:   "TODO1",
	Long:    `TODO2`,
}

func init() {
	rootCmd.PersistentFlags().Bool("no-color", false, "disable both colors and emoji from visible output")
	rootCmd.PersistentFlags().Bool("no-emoji", false, "disable emoji from visible output (but keep colors)")
	rootCmd.PersistentFlags().Bool("no-gui", false, "disable GUI Diaglog Prompt")
	rootCmd.PersistentFlags().Bool("verbose", false, "enable verbose output")

	rootCmd.SetVersionTemplate(config.VersionShortTmpl)
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().Bool("full", false, "display full version information")

	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configListCmd)

	rootCmd.AddCommand(cacheCmd)
	cacheCmd.AddCommand(cacheCleanCmd)
	cacheCleanCmd.Flags().Bool("password", false, "delete yubikey oath application password cache")
	cacheCleanCmd.Flags().Bool("credential", false, "delete temporary session credential cache")

	rootCmd.AddCommand(assumeCmd)

	profileFlag := "profile"
	assumeCmd.Flags().StringP(profileFlag, "p", "", "aws profile to use from config")
	err := assumeCmd.MarkFlagRequired(profileFlag)
	if err != nil {
		panic(err)
	}

	g, err := utils.ParseFlags(application.GlobalFlags{}, rootCmd)
	if err != nil {
		panic(err)
	}

	msg.Init(msg.Options{
		SilentMode:  true,
		VerboseMode: g.Verbose,
		ColorMode:   !g.NoColor,
		EmojiMode:   !g.NoEmoji,
	})
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
