package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Profile         string
	DurationSeconds int    `mapstructure:"duration_seconds"`
	YubikeySerial   string `mapstructure:"yubikey_serial"`
	YubikeyLabel    string `mapstructure:"yubikey_label"`
	Verbose         bool   `mapstructure:"verbose"`
	HideArns        bool   `mapstructure:"hide_arns"`
	DisableDialog   bool   `mapstructure:"disable_dialog"`
	DisableRefresh  bool   `mapstructure:"disable_refresh"`
}

func (c *Config) Load(cmd *cobra.Command) error {

	var err error

	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	configPath := filepath.Join(homedir, ".aws-mfa-credential-process")

	os.MkdirAll(configPath, os.ModePerm)

	// TODO how to support testing with temp file etc? (e.g. in profile_test.go)

	v := viper.New()

	v.SetConfigName("config")   // name of config file (without extension)
	v.AddConfigPath(configPath) // call multiple times to add many search paths

	v.SetDefault("DurationSeconds", 3600) // set the default duration seconds to match AWS defaults
	v.SetDefault("Verbose", false)        // by default, disable verbose output
	v.SetDefault("HideArns", false)       // by default, do not hide ARNs in verbose output
	v.SetDefault("DisableDialog", false)  // by default, use GUI dialog prompt
	v.SetDefault("DisableRefresh", false) // by default, refresh credentials like botocore does

	err = v.ReadInConfig()
	if err != nil && err != err.(viper.ConfigFileNotFoundError) {
		panic(err)
	}

	v.BindPFlags(cmd.Flags())

	err = v.Unmarshal(&c)
	if err != nil {
		return err
	}

	return err
}
