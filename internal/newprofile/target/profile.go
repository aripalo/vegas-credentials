package target

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/aripalo/vegas-credentials/internal/config"
	"gopkg.in/ini.v1"
)

type TargetProfile struct {
	SourceProfile   string `ini:"vegas_source_profile"`
	RoleArn         string `ini:"vegas_role_arn"`
	DurationSeconds int    `ini:"duration_seconds"`
	Region          string `ini:"region"`
	RoleSessionName string `ini:"role_session_name"`
	ExternalID      string `ini:"external_id"`
}

func New(targetName string) (*TargetProfile, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	configPath := filepath.Join(home, ".aws", "config")
	return loadWithPath(targetName, configPath)
}

func loadWithPath(targetName string, configPath string) (*TargetProfile, error) {

	t := new(TargetProfile)

	// set default duration
	t.DurationSeconds = config.Defaults.DurationSeconds.Value

	cfg, err := ini.Load(configPath)
	if err != nil {
		return t, err
	}

	profileSection := fmt.Sprintf("profile %s", targetName)

	section, err := cfg.GetSection(profileSection)
	if err != nil {
		return t, err
	}

	err = section.MapTo(t)
	if err != nil {
		return t, err
	}

	if t.SourceProfile == "" {
		return t, fmt.Errorf(`Profile "%s" does not contain "vegas_source_profile"`, targetName)
	}

	if t.RoleArn == "" {
		return t, fmt.Errorf(`Profile "%s" does not contain "vegas_role_arn"`, targetName)
	}

	if !iamRolePattern.MatchString(t.RoleArn) {
		return t, fmt.Errorf(`Profile "%s" contains invalid vegas_role_arn "%s". Must satisty %s`, targetName, t.RoleArn, iamRolePatternBase)
	}

	if t.RoleSessionName != "" && !iamResourceNamePattern.MatchString(t.RoleSessionName) {
		return t, fmt.Errorf(`Profile "%s" contains invalid role_session_name "%s". Must satisfy %s`, targetName, t.RoleSessionName, iamResourceNamePAtternFull)
	}

	if t.ExternalID != "" && !externalIdPattern.MatchString(t.ExternalID) {
		return t, fmt.Errorf(`Profile "%s" contains invalid external_id "%s". Must satisfy %s`, targetName, t.ExternalID, externalIdPatternBase)
	}

	return t, nil
}

/*
User, Role or Role Session Names can be maximum 64 characters.

Names of users, groups, roles, policies, instance profiles, and server certificates must be alphanumeric, including the following common characters: plus (+), equal (=), comma (,), period (.), at (@), underscore (_), and hyphen (-).

https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html#reference_iam-quotas-entity-length
*/
var iamResourceNamePatternBase = `[a-zA-Z0-9_+=,.@-]{1,64}`
var iamResourceNamePAtternFull = fmt.Sprintf("^%s$", iamResourceNamePatternBase)
var iamResourceNamePattern = regexp.MustCompile(iamResourceNamePAtternFull)

/*
Names of users, groups, roles, policies, instance profiles, and server certificates must be alphanumeric, including the following common characters: plus (+), equal (=), comma (,), period (.), at (@), underscore (_), and hyphen (-).

https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html
*/
var iamRolePatternBase = fmt.Sprintf(`^arn:aws:iam:\d*:\d{12}:role\/%s$`, iamResourceNamePatternBase)
var iamRolePattern = regexp.MustCompile(iamRolePatternBase)

/*
The external ID value that a third party uses to assume a role must have a minimum of 2 characters and a maximum of 1,224 characters. The value must be alphanumeric without white space. It can also include the following symbols: plus (+), equal (=), comma (,), period (.), at (@), colon (:), forward slash (/), and hyphen (-).

In reality the maximum length is 1224, but that causes "Quantifier range is too large" error in golang so let's leave that validation to AWS.

https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html
*/
var externalIdPatternBase = `^[a-zA-Z0-9+=,.@:\/-]{2,}$`
var externalIdPattern = regexp.MustCompile(externalIdPatternBase)
