package ykman

import (
	"fmt"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "not available",
			input:    "notavailalbe",
			expected: "",
		},
		{
			name:  "available",
			input: "ls",
			expected: func() string {
				result, err := exec.LookPath("ls")
				if err != nil || result == "" {
					return "incorrect"
				}
				return result
			}(),
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.name)
		t.Run(name, func(t *testing.T) {
			executable = test.input
			actual := GetPath()
			assert.Equal(t, test.expected, actual)
		})
	}
}
