package assumecfg

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAssumable(t *testing.T) {
	tests := []struct {
		name        string
		datasource  string
		profileName string
		expected    AssumeCfg
		err         error
	}{
		{
			name:        "valid-minimal",
			datasource:  "./testdata/valid-minimal.ini",
			profileName: "frank@concerts",
			expected: AssumeCfg{
				ProfileName:     "frank@concerts",
				MfaSerial:       "arn:aws:iam::111111111111:mfa/FrankSinatra",
				YubikeyLabel:    "arn:aws:iam::111111111111:mfa/FrankSinatra",
				RoleArn:         "arn:aws:iam::222222222222:role/SingerRole",
				RoleSessionName: "SinatraAtTheSands",
				SourceProfile:   "default",
				DurationSeconds: 3600,
				Checksum:        "d5656b38d196c58de4ab1e2cdd5e2715611e9282",
			},
		},
		{
			name:        "invalid-missing-source",
			datasource:  "./testdata/invalid-missing-source.ini",
			profileName: "frank@concerts",
			expected:    AssumeCfg{},
			err:         errors.New("Profile \"frank@concerts\" does not contain \"vegas_source_profile\""),
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.name)
		t.Run(name, func(t *testing.T) {
			actual, err := New(test.datasource, test.profileName)
			assert.Equal(t, test.err, err)
			assert.Equal(t, test.expected, actual)
		})
	}
}
