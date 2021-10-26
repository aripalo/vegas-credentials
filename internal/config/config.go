package config

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Flags provides global/shared configuration passed downstream
type Flags struct {
	Profile                 string `mapstructure:"profile"`
	DurationSeconds         int    `mapstructure:"duration_seconds"` // TODO remove
	YubikeySerial           string `mapstructure:"yubikey_serial"`   // TODO remove
	YubikeyLabel            string `mapstructure:"yubikey_label"`    // TODO remove
	Debug                   bool   `mapstructure:"debug"`
	Verbose                 bool   `mapstructure:"verbose"`
	HideArns                bool   `mapstructure:"hide_arns"`
	DisableDialog           bool   `mapstructure:"disable_dialog"`
	DisableMandatoryRefresh bool   `mapstructure:"disable_refresh"`
	NoColor                 bool   `mapstructure:"no_color"`
}

// TODO how to support testing with temp file etc? (e.g. in profile_test.go)
func (c *Flags) Load(cmd *cobra.Command) error {

	var err error

	// New viper instance
	v := viper.New()

	// Set defaults
	v.SetDefault(Defaults.DurationSeconds.Name, Defaults.DurationSeconds.Value)
	v.SetDefault(Defaults.Debug.Name, Defaults.Debug.Value)
	v.SetDefault(Defaults.Verbose.Name, Defaults.Verbose.Value)
	v.SetDefault(Defaults.HideArns.Name, Defaults.HideArns.Value)
	v.SetDefault(Defaults.DisableDialog.Name, Defaults.DisableDialog.Value)
	v.SetDefault(Defaults.DisableMandatoryRefresh.Name, Defaults.DisableMandatoryRefresh.Value)
	v.SetDefault(Defaults.NoColor.Name, Defaults.NoColor.Value)

	// Set Config file name (without extension)
	v.SetConfigName("config")

	// Config file search pahts
	v.AddConfigPath(fmt.Sprintf("$XDG_CONFIG_HOME/%s", APP_NAME))
	v.AddConfigPath(fmt.Sprintf("$HOME/.config/%s", APP_NAME))
	v.AddConfigPath(fmt.Sprintf("$HOME/.%s", APP_NAME))

	// Read from Config
	err = v.ReadInConfig()
	// Config file is optional, so ignore
	if err != nil && err != err.(viper.ConfigFileNotFoundError) {
		return err
	}

	// Read CLI flags
	// https://github.com/spf13/viper#working-with-flags
	err = v.BindPFlags(cmd.Flags())
	if err != nil {
		return err
	}

	// Unmarshal viper configuration into config.Config
	err = v.Unmarshal(&c, decodeWithMixedCasing)

	return err
}

// decodeWithMixedCasing enables support for different kinds of casing in configuration (snake, param, etc)
// This works because Viper prefers CLI flags to config file & default values.
// https://pkg.go.dev/github.com/mitchellh/mapstructure#DecoderConfig.MatchName
func decodeWithMixedCasing(config *mapstructure.DecoderConfig) {
	config.MatchName = func(mapKey string, fieldName string) bool {
		snakedMapKey := strcase.ToSnake(mapKey)
		return strings.EqualFold(snakedMapKey, fieldName)
	}
}
