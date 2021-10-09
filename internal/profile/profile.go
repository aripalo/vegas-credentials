package profile

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

func GetProfile(profileName string) (Profile, error) {

	var err error

	homedir, err := os.UserHomeDir()
	configPath := fmt.Sprintf("%s/.aws/config", homedir)

	cfg, err := ini.Load(configPath)

	sectionName := fmt.Sprintf("profile %s", profileName)

	profile := new(Profile)

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
