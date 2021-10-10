package profile

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

func GetProfile(profileName string) (Profile, error) {

	var err error

	homedir, err := os.UserHomeDir()
	configPath := filepath.Join(homedir, ".aws/config")

	cfg, err := ini.Load(configPath)

	sectionName := fmt.Sprintf("profile %s", profileName)

	profile := &Profile{
		DurationSeconds: 3600, // default to 1 hour as AWS does
	}

	err = cfg.Section(sectionName).MapTo(profile)

	return *profile, err

}

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
