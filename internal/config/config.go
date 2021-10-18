package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config provides global/shared configuration passed downstream
type Config struct {
	Profile         string `mapstructure:"profile"`
	DurationSeconds int    `mapstructure:"duration_seconds"`
	YubikeySerial   string `mapstructure:"yubikey_serial"`
	YubikeyLabel    string `mapstructure:"yubikey_label"`
	Debug           bool   `mapstructure:"debug"`
	Verbose         bool   `mapstructure:"verbose"`
	HideArns        bool   `mapstructure:"hide_arns"`
	DisableDialog   bool   `mapstructure:"disable_dialog"`
	DisableRefresh  bool   `mapstructure:"disable_refresh"`
	NoColor         bool   `mapstructure:"no_color"`
}

// TODO how to support testing with temp file etc? (e.g. in profile_test.go)
func (c *Config) Load(cmd *cobra.Command) error {

	var err error

	// Initialize config Path
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	configPath := filepath.Join(homedir, ".aws-mfa-credential-process")
	os.MkdirAll(configPath, os.ModePerm) // TODO maybe remove this? XDG thing also...

	// New viper instance
	v := viper.New()

	// Set defaults
	v.SetDefault(Defaults.DurationSeconds.Name, Defaults.DurationSeconds.Value) // set the default duration seconds to match AWS defaults
	v.SetDefault(Defaults.Debug.Name, Defaults.Debug.Value)                     // by default, disable debug information
	v.SetDefault(Defaults.Verbose.Name, Defaults.Verbose.Value)                 // by default, disable verbose output
	v.SetDefault(Defaults.HideArns.Name, Defaults.HideArns.Value)               // by default, do not hide ARNs in verbose output
	v.SetDefault(Defaults.DisableDialog.Name, Defaults.DisableDialog.Value)     // by default, use GUI dialog prompt
	v.SetDefault(Defaults.DisableRefresh.Name, Defaults.DisableRefresh.Value)   // by default, refresh credentials like botocore does
	v.SetDefault(Defaults.NoColor.Name, Defaults.NoColor.Value)                 // by default, allow colored output (depending on environment)

	// Set Config Path
	v.SetConfigName("config")   // name of config file (without extension)
	v.AddConfigPath(configPath) // call multiple times to add many search paths

	// Read from Config
	err = v.ReadInConfig()
	if err != nil && err != err.(viper.ConfigFileNotFoundError) {
		return err
	} else {
		// Config file is optional, so ignore
		err = nil
	}

	// Read CLI flags
	// https://github.com/spf13/viper#working-with-flags
	v.BindPFlags(cmd.Flags())

	// Unmarshal viper configuration into config.Config
	err = v.Unmarshal(&c, decodeWithMixedCasing)

	return err
}

// decodeWithMixedCasing enables support for different kinds of casing in configuration (snake, param, etc)
// https://pkg.go.dev/github.com/mitchellh/mapstructure#DecoderConfig.MatchName
func decodeWithMixedCasing(config *mapstructure.DecoderConfig) {
	config.MatchName = func(mapKey string, fieldName string) bool {
		snakedMapKey := strcase.ToSnake(mapKey)
		return strings.EqualFold(snakedMapKey, fieldName)
	}
}
