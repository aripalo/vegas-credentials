package awsini

import (
	"fmt"

	"gopkg.in/ini.v1"
)

// Profile represents a type constraint of which structs the Load function
// can actually load out from ~/.aws/config ini.
type Profile interface {
	*User | *Role
}

// DataSource represents the source where to load ini-configuration.
// Can be either file name with string type or raw data in []byte.
// https://pkg.go.dev/gopkg.in/ini.v1@v1.66.4#Load
type DataSource interface {
	string | []byte
}

// LoadProfile parses an AWS profile configuration the given data source
// (either file name string or file content as []byte) for given profile name.
// Effectively unmarshals the result into a struct.
func LoadProfile[P Profile, D DataSource](dataSource D, profileName string, profile P) error {

	cfg, err := ini.Load(dataSource)
	if err != nil {
		return err
	}

	title := formatTitle(profileName)

	section, err := cfg.GetSection(title)
	if err != nil {
		return err
	}

	err = section.MapTo(profile)
	if err != nil {
		return err
	}

	return nil
}

// Resolve the AWS config ini section title for given profile name.
// By default each profile should be configured with the "profile " prefix,
// except for default which does not have that prefix.
func formatTitle(profileName string) string {
	if profileName == "default" {
		return profileName
	}
	return fmt.Sprintf("profile %s", profileName)
}
