package awsini

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadUserProfile(t *testing.T) {
	tests := []struct {
		name        string
		datasource  string
		profileName string
		expected    User
		err         error
	}{
		{
			name:        "missing profile - default",
			datasource:  "./testdata/user/missing-profile.ini",
			profileName: "default",
			expected:    User{},
			err:         errors.New("section \"default\" does not exist"),
		},
		{
			name:        "missing profile - celine",
			datasource:  "./testdata/user/missing-profile.ini",
			profileName: "celine",
			expected:    User{},
			err:         errors.New("section \"profile celine\" does not exist"),
		},
		{
			name:        "basic - default",
			datasource:  "./testdata/user/basic.ini",
			profileName: "default",
			expected: User{
				MfaSerial: "arn:aws:iam::111111111111:mfa/FrankSinatra",
			},
		},
		{
			name:        "basic - celine",
			datasource:  "./testdata/user/basic.ini",
			profileName: "celine",
			expected: User{
				MfaSerial: "arn:aws:iam::111111111111:mfa/CelineDion",
			},
		},
		{
			name:        "with region - default",
			datasource:  "./testdata/user/with-region.ini",
			profileName: "default",
			expected: User{
				MfaSerial: "arn:aws:iam::111111111111:mfa/FrankSinatra",
				Region:    "us-west-1",
			},
		},
		{
			name:        "with region - celine",
			datasource:  "./testdata/user/with-region.ini",
			profileName: "celine",
			expected: User{
				MfaSerial: "arn:aws:iam::111111111111:mfa/CelineDion",
				Region:    "ca-central-1",
			},
		},
		{
			name:        "with yubikey serial - default",
			datasource:  "./testdata/user/with-yubikey-serial.ini",
			profileName: "default",
			expected: User{
				MfaSerial:     "arn:aws:iam::111111111111:mfa/FrankSinatra",
				YubikeySerial: "12345678",
			},
		},
		{
			name:        "with yubikey serial - celine",
			datasource:  "./testdata/user/with-yubikey-serial.ini",
			profileName: "celine",
			expected: User{
				MfaSerial:     "arn:aws:iam::111111111111:mfa/CelineDion",
				YubikeySerial: "87654321",
			},
		},
		{
			name:        "with yubikey label - default",
			datasource:  "./testdata/user/with-yubikey-label.ini",
			profileName: "default",
			expected: User{
				MfaSerial:     "arn:aws:iam::111111111111:mfa/FrankSinatra",
				YubikeySerial: "12345678",
				YubikeyLabel:  "aws-frank",
			},
		},
		{
			name:        "with yubikey label - celine",
			datasource:  "./testdata/user/with-yubikey-label.ini",
			profileName: "celine",
			expected: User{
				MfaSerial:     "arn:aws:iam::111111111111:mfa/CelineDion",
				YubikeySerial: "87654321",
				YubikeyLabel:  "aws-celine",
			},
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.name)
		t.Run(name, func(t *testing.T) {
			var profile User
			err := LoadProfile(test.datasource, test.profileName, &profile)
			assert.Equal(t, test.err, err)
			assert.Equal(t, test.expected, profile)
		})
	}
}

func TestLoadRoleProfile(t *testing.T) {
	tests := []struct {
		name        string
		datasource  string
		profileName string
		expected    Role
		err         error
	}{
		{
			name:        "valid minimal - frank@concerts",
			datasource:  "./testdata/role/valid-minimal.ini",
			profileName: "frank@concerts",
			expected: Role{
				RoleArn:       "arn:aws:iam::222222222222:role/SingerRole",
				SourceProfile: "default",
			},
		},
		{
			name:        "valid minimal - celine@concerts",
			datasource:  "./testdata/role/valid-minimal.ini",
			profileName: "celine@concerts",
			expected: Role{
				RoleArn:       "arn:aws:iam::222222222222:role/SingerRole",
				SourceProfile: "celine",
			},
		},
		{
			name:        "valid full - frank@concerts",
			datasource:  "./testdata/role/valid-full.ini",
			profileName: "frank@concerts",
			expected: Role{
				RoleArn:         "arn:aws:iam::222222222222:role/SingerRole",
				SourceProfile:   "default",
				Region:          "us-west-1",
				DurationSeconds: 4383,
				RoleSessionName: "SinatraAtTheSands",
				ExternalID:      "0093624694724",
			},
		},
		{
			name:        "valid full - celine@concerts",
			datasource:  "./testdata/role/valid-full.ini",
			profileName: "celine@concerts",
			expected: Role{
				RoleArn:         "arn:aws:iam::222222222222:role/SingerRole",
				SourceProfile:   "celine",
				Region:          "ca-central-1",
				DurationSeconds: 3536,
				RoleSessionName: "ANewDayLiveInLasVegas",
				ExternalID:      "0886971371697",
			},
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.name)
		t.Run(name, func(t *testing.T) {
			var profile Role
			err := LoadProfile(test.datasource, test.profileName, &profile)
			assert.Equal(t, test.err, err)
			assert.Equal(t, test.expected, profile)
		})
	}
}
