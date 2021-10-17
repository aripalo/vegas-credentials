package profile

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aripalo/aws-mfa-credential-process/internal/config"
	"github.com/spf13/viper"
)

func Get(config *config.Config) (ProfileConfig, error) {
	var profileConfig ProfileConfig

	home, err := os.UserHomeDir()
	if err != nil {
		return profileConfig, err
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
		return profileConfig, err
	}

	var configurations map[string]ProfileConfig

	err = v.Unmarshal(&configurations)
	if err != nil {
		return profileConfig, err
	}

	profileConfig = configurations[section]
	if profileConfig.AssumeRoleArn == "" || profileConfig.SourceProfile == "" {
		return profileConfig, errors.New("Invalid profile")
	}

	return profileConfig, nil
}

type ProfileConfig struct {
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
