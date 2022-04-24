package awsini

import (
	"fmt"

	"gopkg.in/ini.v1"
)

func resolveSectionTitle(profileName string) string {
	// default profile should not have "profile " prefix in section title
	if profileName == "default" {
		return profileName
	} else {
		return fmt.Sprintf("profile %s", profileName)
	}
}

func load(source interface{}, profileName string, target any) error {

	cfg, err := ini.Load(source)
	if err != nil {
		return err
	}

	sectionTitle := resolveSectionTitle(profileName)

	section, err := cfg.GetSection(sectionTitle)
	if err != nil {
		return err
	}

	err = section.MapTo(target)
	if err != nil {
		return err
	}

	return nil
}
