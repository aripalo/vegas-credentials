package assumeopts

import (
	"errors"

	"github.com/aripalo/vegas-credentials/internal/awsini"
	"github.com/aripalo/vegas-credentials/internal/utils"
)

// New returns a struct representing all the information required to assume an
// IAM role with MFA. Effectively it parses the given dataSource
// (either file name with string type or raw data in []byte) and finds the
// correct configuration by looking up the given profileName.
func New[D awsini.DataSource](dataSource D, profileName string) (Opts, error) {
	var opts Opts
	var role awsini.Role
	var user awsini.User

	err := awsini.LoadProfile(dataSource, profileName, &role)
	if err != nil {
		return opts, err
	}

	if role.SourceProfile == "" {
		return opts, errors.New("vegas_source_profile not configured")
	}

	err = awsini.LoadProfile(dataSource, role.SourceProfile, &user)
	if err != nil {
		return opts, err
	}

	opts = Opts{
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

	if opts.MfaSerial == "" {
		return opts, errors.New("mfa_serial not configured")
	}

	checksum, err := utils.CalculateChecksum(opts)
	if err != nil {
		return opts, err
	}

	opts.Checksum = checksum

	return opts, nil
}
