package credentials

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOutput(t *testing.T) {

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

			var output bytes.Buffer

			test.input.output = &output

			err := test.input.Output()
			if test.errMessage != "" {
				assert.Equal(t, test.errMessage, err.Error())
			}

			actual := output.Bytes()

			assert.Equal(t, test.expected, string(actual))
		})
	}
}
