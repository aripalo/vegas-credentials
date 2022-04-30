package credentials

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testData struct {
	BooleanValue bool
	StringValue  string
}

var expirationLayout = "2006-01-02T15:04:05Z"
var expirationTime = "2022-04-30T23:05:12Z"

var serialized = strings.TrimSpace(fmt.Sprintf(`
{
    "Version": 1,
    "AccessKeyId": "AKIAIOSFODNN7EXAMPLE",
    "SecretAccessKey": "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
    "SessionToken": "EXAMPLEzie5aiZielooFaequeix8fuaxaib1eex8teibuad0sho5aimie7shae8reed2pahn8eGieph5eik9Ju2anieGahfeipe8eephie7eelaeTaemo5eicohvu5Daec8aengei8queefiephom2ohgh1ou6shah8toni2phe3Eajoopheishi2BeeZ0fixeeph6pheeNge3mieV8ohs9iu1aechie3aoCeeheikuaweeY3ui0bai1hai8uem2yeingaciopii3aiz6iiTiKahdowienohveiw6oofangae0Quaishae5buashughapowei6ohphah2Zoo",
    "Expiration": "%s"
}`, expirationTime))

func TestSerialize(t *testing.T) {

	expiration, err := time.Parse(expirationLayout, expirationTime)
	require.NoError(t, err)

	tests := []struct {
		name       string
		input      Credentials
		expected   string
		errMessage string
	}{
		{
			name: "input valid",
			input: Credentials{
				Version:         1,
				AccessKeyID:     "AKIAIOSFODNN7EXAMPLE",
				SecretAccessKey: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
				SessionToken:    "EXAMPLEzie5aiZielooFaequeix8fuaxaib1eex8teibuad0sho5aimie7shae8reed2pahn8eGieph5eik9Ju2anieGahfeipe8eephie7eelaeTaemo5eicohvu5Daec8aengei8queefiephom2ohgh1ou6shah8toni2phe3Eajoopheishi2BeeZ0fixeeph6pheeNge3mieV8ohs9iu1aechie3aoCeeheikuaweeY3ui0bai1hai8uem2yeingaciopii3aiz6iiTiKahdowienohveiw6oofangae0Quaishae5buashughapowei6ohphah2Zoo",
				Expiration:      expiration,
			},
			expected: serialized,
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.name)
		t.Run(name, func(t *testing.T) {
			actual, err := test.input.Serialize()
			if test.errMessage != "" {
				assert.Equal(t, test.errMessage, err.Error())
			}
			assert.Equal(t, test.expected, string(actual))
		})
	}
}

func TestDeserialize(t *testing.T) {

	expiration, err := time.Parse(expirationLayout, expirationTime)
	require.NoError(t, err)

	tests := []struct {
		name       string
		input      string
		expected   Credentials
		errMessage string
	}{
		{
			name:  "input valid",
			input: serialized,
			expected: Credentials{
				Version:         1,
				AccessKeyID:     "AKIAIOSFODNN7EXAMPLE",
				SecretAccessKey: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
				SessionToken:    "EXAMPLEzie5aiZielooFaequeix8fuaxaib1eex8teibuad0sho5aimie7shae8reed2pahn8eGieph5eik9Ju2anieGahfeipe8eephie7eelaeTaemo5eicohvu5Daec8aengei8queefiephom2ohgh1ou6shah8toni2phe3Eajoopheishi2BeeZ0fixeeph6pheeNge3mieV8ohs9iu1aechie3aoCeeheikuaweeY3ui0bai1hai8uem2yeingaciopii3aiz6iiTiKahdowienohveiw6oofangae0Quaishae5buashughapowei6ohphah2Zoo",
				Expiration:      expiration,
			},
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.name)
		t.Run(name, func(t *testing.T) {
			c := Credentials{}
			err := c.Deserialize(json.RawMessage(test.input))
			if test.errMessage != "" {
				assert.Equal(t, test.errMessage, err.Error())
			}
			assert.Equal(t, test.expected, c)
		})
	}
}
