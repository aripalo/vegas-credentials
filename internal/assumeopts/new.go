package assumeopts

import (
	"errors"

	"github.com/aripalo/vegas-credentials/internal/awsini"
)

// New returns a struct representing all the information required to assume an
// IAM role with MFA. Effectively it parses the given dataSource
// (either file name with string type or raw data in []byte) and finds the
// correct configuration by looking up the given profileName.
func New[D awsini.DataSource](dataSource D, profileName string) (AssumeOpts, error) {
	var assumable AssumeOpts
	var role awsini.Role
	var user awsini.User

	err := awsini.LoadProfile(dataSource, profileName, &role)
	if err != nil {
		return assumable, err
	}

	if role.SourceProfile == "" {
		return assumable, errors.New("vegas_source_profile not configured")
	}

	err = awsini.LoadProfile(dataSource, role.SourceProfile, &user)
	if err != nil {
		return assumable, err
	}

	assumable = AssumeOpts{
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

	if assumable.MfaSerial == "" {
		return assumable, errors.New("mfa_serial not configured")
	}

	return assumable, nil
}
