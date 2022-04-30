package cache

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKey(t *testing.T) {
	tests := []struct {
		name     string
		prefix   string
		checksum string
		expected string
	}{
		{
			name:     "valid",
			prefix:   "foo",
			checksum: "bar",
			expected: "foo__bar",
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.name)
		t.Run(name, func(t *testing.T) {
			actual := Key(test.prefix, test.checksum)
			assert.Equal(t, test.expected, actual)
		})
	}
}
