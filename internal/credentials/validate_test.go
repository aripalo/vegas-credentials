package credentials

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name     string
		input    Credentials
		expected error
	}{
		{
			name: "correct",
			input: Credentials{
				Version:         1,
				AccessKeyID:     "ID",
				SecretAccessKey: "SECRET",
				SessionToken:    "TOKEN",
				Expiration:      time.Now().Add(time.Minute * 5),
			},
			expected: nil,
		},
		{
			name: "incorrect version",
			input: Credentials{
				Version:         0,
				AccessKeyID:     "ID",
				SecretAccessKey: "SECRET",
				SessionToken:    "TOKEN",
				Expiration:      time.Now().Add(time.Minute * 5),
			},
			expected: errors.New("Incorrect Version"),
		},
		{
			name: "access key id missing",
			input: Credentials{
				Version:         1,
				SecretAccessKey: "SECRET",
				SessionToken:    "TOKEN",
				Expiration:      time.Now().Add(time.Minute * 5),
			},
			expected: errors.New("Missing AccessKeyID"),
		},
		{
			name: "secret access key missing",
			input: Credentials{
				Version:      1,
				AccessKeyID:  "ID",
				SessionToken: "TOKEN",
				Expiration:   time.Now().Add(time.Minute * 5),
			},
			expected: errors.New("Missing SecretAccessKey"),
		},
		{
			name: "session token missing",
			input: Credentials{
				Version:         1,
				AccessKeyID:     "ID",
				SecretAccessKey: "SECRET",
				Expiration:      time.Now().Add(time.Minute * 5),
			},
			expected: errors.New("Missing SessionToken"),
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.name)
		t.Run(name, func(t *testing.T) {
			actual := test.input.Validate()
			assert.Equal(t, test.expected, actual)
		})
	}
}
