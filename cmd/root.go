package cmd

import (
	"github.com/aripalo/aws-mfa-credential-process/internal/config"
	"github.com/spf13/cobra"
)

var cfgFile string
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
	//cobra.OnInitialize(initConfig)

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

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.newApp.yaml)")
}

// initConfig reads in config file and ENV variables if set.
/*
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".newApp" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".newApp")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
*/
