package profile

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/aripalo/aws-mfa-credential-process/internal/config"
	"github.com/iancoleman/strcase"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

func (p *Profile) Load(config *config.Config) error {
	var profileConfig Profile
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	profileConfig, err = loadWithPath(profileConfig, config, filepath.Join(home, ".aws"))
	if err != nil {
		return err
	}

	*p = profileConfig

	return nil
}

func loadWithPath(profileConfig Profile, config *config.Config, configPath string) (Profile, error) {

	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("ini")
	v.AddConfigPath(configPath)

	section := fmt.Sprintf("profile %s", config.Profile)

	v.SetDefault(fmt.Sprintf("%s.duration_seconds", section), config.DurationSeconds)
	v.SetDefault(fmt.Sprintf("%s.yubikey_serial", section), config.YubikeySerial)
	v.SetDefault(fmt.Sprintf("%s.yubikey_label", section), config.YubikeyLabel)

	err := v.ReadInConfig()
	if err != nil {
		return profileConfig, err
	}

	var configurations map[string]Profile

	err = v.Unmarshal(&configurations, decodeWithMixedCasing)
	if err != nil {
		return profileConfig, err
	}

	profileConfig = configurations[section]
	if profileConfig.RoleArn == "" || profileConfig.SourceProfile == "" {
		return profileConfig, errors.New(fmt.Sprintf("Invalid Profile Configuration for %s", config.Profile))
	}

	return profileConfig, nil
}

type Profile struct {
	YubikeySerial   string `mapstructure:"yubikey_serial"`
	YubikeyLabel    string `mapstructure:"yubikey_label"`
	SourceProfile   string `mapstructure:"source_profile"`
	RoleArn         string `mapstructure:"role_arn"`
	MfaSerial       string `mapstructure:"mfa_serial"`
	DurationSeconds int    `mapstructure:"duration_seconds"`
	Region          string `mapstructure:"region"`
	RoleSessionName string `mapstructure:"role_session_name"`
	ExternalID      string `mapstructure:"external_id"`
}

// decodeWithMixedCasing enables support for different kinds of casing in configuration (snake, param, etc)
// This works because Viper prefers CLI flags to config file & default values.
// https://pkg.go.dev/github.com/mitchellh/mapstructure#DecoderConfig.MatchName
func decodeWithMixedCasing(config *mapstructure.DecoderConfig) {
	config.MatchName = func(mapKey string, fieldName string) bool {

		// Underscore prefixes are used for making Terraform to work
		// as TF fails if source_profile defined, so we support TF by allowing _source_profile
		prefixRemovedMapKey := strings.TrimPrefix(mapKey, "_")

		// Convert to snake_case (in case the key was provided in other casing)
		snakedMapKey := strcase.ToSnake(prefixRemovedMapKey)

		// Handle specific undocumented "feature" ... which may be removed later
		if snakedMapKey == "assume_role_arn" {
			snakedMapKey = "role_arn"
		}

		// EqualFold is the default MatchName function
		return strings.EqualFold(snakedMapKey, fieldName)
	}
}
