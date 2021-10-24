package sourceprofile

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

type SourceProfile struct {
	YubikeySerial string `ini:"vegas_yubikey_serial"`
	YubikeyLabel  string `ini:"vegas_yubikey_label"`
	MfaSerial     string `ini:"mfa_serial"`
	Region        string `ini:"region"`
}

func New(profileName string) (*SourceProfile, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	configPath := filepath.Join(home, ".aws", "config")
	return loadWithPath(profileName, configPath)
}

func loadWithPath(profileName string, configPath string) (*SourceProfile, error) {

	sourceProfile := new(SourceProfile)

	// set defaults
	sourceProfile.Region = os.Getenv("AWS_REGION")

	cfg, err := ini.Load(configPath)
	if err != nil {
		return sourceProfile, err
	}

	profileSection := fmt.Sprintf("profile %s", profileName)

	// default profile should not have "profile " prefix in section title
	if profileName == "default" {
		profileSection = profileName
	}

	section, err := cfg.GetSection(profileSection)
	if err != nil {
		return sourceProfile, err
	}

	err = section.MapTo(sourceProfile)
	if err != nil {
		return sourceProfile, err
	}

	// Use MFA serial (ARN) as the OATH Account Label for Yubikey
	if sourceProfile.YubikeySerial != "" && sourceProfile.YubikeyLabel == "" {
		sourceProfile.YubikeyLabel = sourceProfile.MfaSerial
	}

	if sourceProfile.MfaSerial == "" {
		return sourceProfile, fmt.Errorf(`Missing "mfa_serial" in profile "%s"`, profileName)
	}

	return sourceProfile, nil
}
