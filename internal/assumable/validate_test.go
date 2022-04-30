package assumable

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name     string
		input    Opts
		expected error
	}{
		{
			name: "mfa_serial missing",
			input: Opts{
				ProfileName:   "frank@concerts",
				RoleArn:       "arn:aws:iam::222222222222:role/SingerRole",
				SourceProfile: "default",
			},
			expected: errors.New("Profile \"frank@concerts\" does not contain \"mfa_serial\""),
		},
		{
			name: "mfa_serial invalid",
			input: Opts{
				ProfileName:   "frank@concerts",
				RoleArn:       "arn:aws:iam::222222222222:role/SingerRole",
				SourceProfile: "default",
				MfaSerial:     "invalid",
			},
			expected: errors.New("Profile \"frank@concerts\" contains invalid mfa_serial \"invalid\". Must satisfy ^arn:aws:iam:\\d*:\\d{12}:mfa\\/.*$"),
		},
		{
			name: "vegas_source_profile missing",
			input: Opts{
				ProfileName: "frank@concerts",
				RoleArn:     "arn:aws:iam::222222222222:role/SingerRole",
				MfaSerial:   "arn:aws:iam::111111111111:mfa/FrankSinatra",
			},
			expected: errors.New("Profile \"frank@concerts\" does not contain \"vegas_source_profile\""),
		},
		{
			name: "vegas_role_arn missing",
			input: Opts{
				ProfileName:   "frank@concerts",
				SourceProfile: "default",
				MfaSerial:     "arn:aws:iam::111111111111:mfa/FrankSinatra",
			},
			expected: errors.New("Profile \"frank@concerts\" does not contain \"vegas_role_arn\""),
		},
		{
			name: "vegas_role_arn invalid",
			input: Opts{
				ProfileName:   "frank@concerts",
				SourceProfile: "default",
				MfaSerial:     "arn:aws:iam::111111111111:mfa/FrankSinatra",
				RoleArn:       "invalid",
			},
			expected: errors.New("Profile \"frank@concerts\" contains invalid vegas_role_arn \"invalid\". Must satisty ^arn:aws:iam:\\d*:\\d{12}:role\\/[a-zA-Z0-9_+=,.@-]{1,64}$"),
		},
		{
			name: "role_session_name invalid",
			input: Opts{
				ProfileName:     "frank@concerts",
				SourceProfile:   "default",
				MfaSerial:       "arn:aws:iam::111111111111:mfa/FrankSinatra",
				RoleArn:         "arn:aws:iam::222222222222:role/SingerRole",
				RoleSessionName: "invalid//",
			},
			expected: errors.New("Profile \"frank@concerts\" contains invalid role_session_name \"invalid//\". Must satisfy ^[a-zA-Z0-9_+=,.@-]{1,64}$"),
		},
		{
			name: "external_id invalid",
			input: Opts{
				ProfileName:   "frank@concerts",
				SourceProfile: "default",
				MfaSerial:     "arn:aws:iam::111111111111:mfa/FrankSinatra",
				RoleArn:       "arn:aws:iam::222222222222:role/SingerRole",
				ExternalID:    "0",
			},
			expected: errors.New("Profile \"frank@concerts\" contains invalid external_id \"0\". Must satisfy ^[a-zA-Z0-9+=,.@:\\/-]{2,}$"),
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.name)
		t.Run(name, func(t *testing.T) {
			actual := test.input.validate()
			assert.Equal(t, test.expected, actual)
		})
	}
}
