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

	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("ini")
	v.AddConfigPath(filepath.Join(home, ".aws"))

	section := fmt.Sprintf("profile %s", config.Profile)

	v.SetDefault(fmt.Sprintf("%s.duration_seconds", section), config.DurationSeconds)
	v.SetDefault(fmt.Sprintf("%s.yubikey_serial", section), config.YubikeySerial)
	v.SetDefault(fmt.Sprintf("%s.yubikey_label", section), config.YubikeyLabel)

	err = v.ReadInConfig()
	if err != nil {
		return err
	}

	var configurations map[string]Profile

	err = v.Unmarshal(&configurations, decodeWithMixedCasing)
	if err != nil {
		return err
	}

	profileConfig = configurations[section]
	if profileConfig.AssumeRoleArn == "" || profileConfig.SourceProfile == "" {
		return errors.New("Invalid profile")
	}

	*p = profileConfig

	return nil
}

type Profile struct {
	YubikeySerial   string `mapstructure:"yubikey_serial"`
	YubikeyLabel    string `mapstructure:"yubikey_label"`
	SourceProfile   string `mapstructure:"source_profile"`
	AssumeRoleArn   string `mapstructure:"assume_role_arn"`
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
		snakedMapKey := strings.TrimPrefix(strcase.ToSnake(mapKey), "_")
		return strings.EqualFold(snakedMapKey, fieldName)
	}
}
