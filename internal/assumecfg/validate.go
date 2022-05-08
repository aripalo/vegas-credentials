package assumecfg

import (
	"fmt"
	"regexp"
)

func (cfg *AssumeCfg) validate() error {

	if cfg.MfaSerial == "" {
		return fmt.Errorf(`Profile "%s" does not contain "mfa_serial"`, cfg.ProfileName)
	}

	if !iamMfaDevicePattern.MatchString(cfg.MfaSerial) {
		return fmt.Errorf(`Profile "%s" contains invalid mfa_serial "%s". Must satisfy %s`, cfg.ProfileName, cfg.MfaSerial, iamMfaDevicePatternBase)
	}

	if cfg.SourceProfile == "" {
		return fmt.Errorf(`Profile "%s" does not contain "vegas_source_profile"`, cfg.ProfileName)
	}

	if cfg.RoleArn == "" {
		return fmt.Errorf(`Profile "%s" does not contain "vegas_role_arn"`, cfg.ProfileName)
	}

	if !iamRolePattern.MatchString(cfg.RoleArn) {
		return fmt.Errorf(`Profile "%s" contains invalid vegas_role_arn "%s". Must satisty %s`, cfg.ProfileName, cfg.RoleArn, iamRolePatternBase)
	}

	if cfg.RoleSessionName != "" && !iamResourceNamePattern.MatchString(cfg.RoleSessionName) {
		return fmt.Errorf(`Profile "%s" contains invalid role_session_name "%s". Must satisfy %s`, cfg.ProfileName, cfg.RoleSessionName, iamResourceNamePAtternFull)
	}

	if cfg.ExternalID != "" && !externalIdPattern.MatchString(cfg.ExternalID) {
		return fmt.Errorf(`Profile "%s" contains invalid external_id "%s". Must satisfy %s`, cfg.ProfileName, cfg.ExternalID, externalIdPatternBase)
	}

	return nil
}

var iamMfaDevicePatternBase = `^arn:aws:iam:\d*:\d{12}:mfa\/.*$`
var iamMfaDevicePattern = regexp.MustCompile(iamMfaDevicePatternBase)

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
