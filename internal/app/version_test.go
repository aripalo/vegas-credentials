package app

import (
	"bytes"
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var version = "v0.0.0-development"

func TestVersion(t *testing.T) {
	tests := []struct {
		name     string
		input    VersionFlags
		expected string
	}{
		{
			name:     "short version",
			input:    VersionFlags{Full: false},
			expected: fmt.Sprintf("%s\n", version),
		},
		{
			name:     "full version",
			input:    VersionFlags{Full: true},
			expected: fmt.Sprintf("vegas-credentials version %s %s/%s\n", version, runtime.GOOS, runtime.GOARCH),
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.name)
		t.Run(name, func(t *testing.T) {
			var output bytes.Buffer
			a := &App{dest: &output}
			err := a.Version(test.input)
			require.NoError(t, err)
			actual := string(output.Bytes())
			assert.Equal(t, test.expected, actual)
		})
	}
}
