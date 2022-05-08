package assumecfg

import (
	"fmt"

	"github.com/aripalo/vegas-credentials/internal/assumecfg/awsini"
	"github.com/aripalo/vegas-credentials/internal/checksum"
)

// New returns a struct representing all the information required to assume an
// IAM role with MFA. Effectively it parses the given dataSource
// (either file name with string type or raw data in []byte) and finds the
// correct configuration by looking up the given profileName.
func New[D awsini.DataSource](dataSource D, profileName string) (AssumeCfg, error) {
	var cfg AssumeCfg
	var role awsini.Role
	var user awsini.User

	err := awsini.LoadProfile(dataSource, profileName, &role)
	if err != nil {
		return cfg, err
	}

	if role.SourceProfile == "" {
		return cfg, fmt.Errorf(`Profile "%s" does not contain "vegas_source_profile"`, profileName)
	}

	err = awsini.LoadProfile(dataSource, role.SourceProfile, &user)
	if err != nil {
		return cfg, err
	}

	cfg = AssumeCfg{
		ProfileName:     profileName,
		MfaSerial:       user.MfaSerial,
		YubikeySerial:   user.YubikeySerial,
		YubikeyLabel:    resolveYubikeyLabel(user.YubikeyLabel, user.MfaSerial),
		Region:          resolveRegion(role.Region, user.Region),
		SourceProfile:   role.SourceProfile,
		RoleArn:         role.RoleArn,
		DurationSeconds: resolveDurationSeconds(role.DurationSeconds),
		RoleSessionName: role.RoleSessionName,
		ExternalID:      role.ExternalID,
	}

	err = cfg.validate()
	if err != nil {
		return cfg, err
	}

	checksum, err := checksum.Generate(cfg)
	if err != nil {
		return cfg, err
	}

	cfg.Checksum = checksum

	return cfg, nil
}
