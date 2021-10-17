package profile

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aripalo/aws-mfa-credential-process/internal/config"
	"gopkg.in/ini.v1"
)

const awsConfigFileLocation string = ".aws/config"

func GetProfile(profileName string) (Profile, error) {

	var path string
	var config *ini.File
	var profile Profile
	var err error

	path, err = resolveConfigPath()
	config, err = loadConfig(path)
	profile, err = loadProfile(config, profileName)

	return profile, err
}

// resolveConfigPath provides the absolute path to AWS config file
func resolveConfigPath() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configPath := filepath.Join(homedir, awsConfigFileLocation)

	return configPath, nil

}

// loadConfig loads an ini-file based configuration from given path
func loadConfig(configPath string) (*ini.File, error) {
	config, err := ini.Load(configPath)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func loadProfile(profileConfig *ini.File, profileName string) (Profile, error) {

	sectionName := fmt.Sprintf("profile %s", profileName)

	profile := &Profile{
		DurationSeconds: config.Config.DurationSeconds,
		YubikeySerial:   config.Config.YubikeySerial,
		YubikeyLabel:    config.Config.YubikeyLabel,
	}

	section, err := profileConfig.GetSection(sectionName)
	if err != nil {
		return *profile, err
	}

	err = section.MapTo(profile)
	if err != nil {
		return *profile, err
	}

	if profile.AssumeRoleArn == "" {
		return *profile, errors.New(fmt.Sprintf("Missing assume_role_arn from profile %s config", profileName))
	}

	if profile.MfaSerial == "" {
		return *profile, errors.New(fmt.Sprintf("Missing mfa_serial from profile %s config", profileName))
	}

	return *profile, nil
}

// TODO validate ARNs

type Profile struct {
	YubikeySerial   string `ini:"yubikey_serial"`
	YubikeyLabel    string `ini:"yubikey_label"`
	SourceProfile   string `ini:"source_profile"`
	AssumeRoleArn   string `ini:"assume_role_arn"`
	MfaSerial       string `ini:"mfa_serial"`
	DurationSeconds int    `ini:"duration_seconds"`
	Region          string `ini:"region"`
	RoleSessionName string `ini:"role_session_name"`
	ExternalID      string `ini:"external_id"`
}
