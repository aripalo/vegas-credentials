package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	// PRODUCT_NAME defines product name used in various outputs
	PRODUCT_NAME string = "aws-mfa-credential-process"

	PRODUCT_CONFIG_LOCATION = "." + PRODUCT_NAME
)

// Configuration holds all the global configuration user may provide via config file
type Configuration struct {
	DurationSeconds int    `mapstructure:"duration_seconds"`
	YubikeySerial   string `mapstructure:"yubikey_serial"`
	YubikeyLabel    string `mapstructure:"yubikey_label"`
	Verbose         bool   `mapstructure:"verbose"`
	HideArns        bool   `mapstructure:"hide_arns"`
	DisableDialog   bool   `mapstructure:"disable_dialog"`
	DisableRefresh  bool   `mapstructure:"disable_refresh"`
}

// Holds the global configuration
var Config Configuration

// Init creates the main config folder with keyring folder
func Init() {

	var err error

	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	configPath := filepath.Join(homedir, PRODUCT_CONFIG_LOCATION)

	os.MkdirAll(configPath, os.ModePerm)

	// TODO how to support testing with temp file etc? (e.g. in profile_test.go)

	viper.SetConfigName("config")             // name of config file (without extension)
	viper.AddConfigPath(configPath)           // call multiple times to add many search paths
	viper.SetDefault("DurationSeconds", 3600) // set the default duration seconds to match AWS defaults
	viper.SetDefault("Verbose", false)        // by default, disable verbose output
	viper.SetDefault("HideArns", false)       // by default, do not hide ARNs in verbose output
	viper.SetDefault("DisableDialog", false)  // by default, use GUI dialog prompt
	viper.SetDefault("DisableRefresh", false) // by default, refresh credentials like botocore does

	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error
		} else {
			// Config file was found but another error was produced
			panic(fmt.Errorf("Fatal error config file: %w \n", err))
		}
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(fmt.Sprintf("unable to decode into struct, %s", err.Error()))
	}
}
