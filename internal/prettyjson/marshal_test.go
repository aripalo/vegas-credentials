package prettyjson

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	BooleanValue bool
	StringValue  string
}

func TestMarshal(t *testing.T) {
	tests := []struct {
		name       string
		input      any
		expected   string
		errMessage string
	}{
		{
			name:       "input unsupported",
			input:      make(chan int),
			expected:   "",
			errMessage: "json: unsupported type: chan int",
		},
		{
			name: "input valid",
			input: testData{
				BooleanValue: true,
				StringValue:  "value",
			},
			expected: "{\n    \"BooleanValue\": true,\n    \"StringValue\": \"value\"\n}",
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.name)
		t.Run(name, func(t *testing.T) {
			actual, err := Marshal(test.input)
			if test.errMessage != "" {
				assert.Equal(t, test.errMessage, err.Error())
			}
			assert.Equal(t, test.expected, actual)
		})
	}
}
