package checksum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		name       string
		input      any
		expected   string
		errMessage string
	}{
		{
			name:       "input nil",
			input:      nil,
			expected:   "",
			errMessage: "checksum: nil input given",
		},
		{
			name:       "input unsupported",
			input:      make(chan int),
			expected:   "",
			errMessage: "json: unsupported type: chan int",
		},
		{
			name:     "input string",
			input:    "foobar",
			expected: "5f6f3065208dde5f4624d7dfafc36a296a526590", // https://passwordsgenerator.net/sha1-hash-generator/
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.name)
		t.Run(name, func(t *testing.T) {
			actual, err := Generate(test.input)
			if test.errMessage != "" {
				assert.Equal(t, test.errMessage, err.Error())
			}
			assert.Equal(t, test.expected, actual)
		})
	}
}
