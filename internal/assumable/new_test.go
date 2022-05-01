package assumable

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
		expected    Opts
		err         error
	}{
		{
			name:        "valid-minimal",
			datasource:  "./testdata/valid-minimal.ini",
			profileName: "frank@concerts",
			expected: Opts{
				ProfileName:     "frank@concerts",
				MfaSerial:       "arn:aws:iam::111111111111:mfa/FrankSinatra",
				YubikeyLabel:    "arn:aws:iam::111111111111:mfa/FrankSinatra",
				RoleArn:         "arn:aws:iam::222222222222:role/SingerRole",
				SourceProfile:   "default",
				DurationSeconds: 3600,
				Checksum:        "d8fdfde29a33d93a04f2c81a014d3558fe09f1c7",
			},
		},
		{
			name:        "invalid-missing-source",
			datasource:  "./testdata/invalid-missing-source.ini",
			profileName: "frank@concerts",
			expected:    Opts{},
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
