package assumable

import (
	"errors"
	"vegas3/internal/awsini"
)

func New(iniConfig any, profileName string) (Assumable, error) {
	var assumable Assumable

	target, err := awsini.NewTargetProfile(iniConfig, profileName)
	if err != nil {
		return assumable, err
	}

	if target.SourceProfile == "" {
		return assumable, errors.New("vegas_source_profile not configured")
	}

	source, err := awsini.NewSourceProfile(iniConfig, target.SourceProfile)
	if err != nil {
		return assumable, err
	}

	assumable = Assumable{
		ProfileName:     profileName,
		MfaSerial:       source.MfaSerial,
		YubikeySerial:   source.YubikeySerial,
		YubikeyLabel:    resolveYubikeyLabel(source.YubikeyLabel, source.MfaSerial),
		Region:          resolveRegion(target.Region, source.Region),
		SourceProfile:   target.SourceProfile,
		RoleArn:         target.RoleArn,
		DurationSeconds: resolveDurationSeconds(target.DurationSeconds),
		RoleSessionName: target.RoleSessionName,
		ExternalID:      target.ExternalID,
	}

	if assumable.MfaSerial == "" {
		return assumable, errors.New("mfa_serial not configured")
	}

	return assumable, nil
}
