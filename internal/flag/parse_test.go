package flag

import (
	"fmt"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

type testFlags struct {
	BooleanFlag bool   `mapstructure:"boolean-flag"`
	StringFlag  string `mapstructure:"string-flag"`
}

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		input    *cobra.Command
		expected testFlags
		err      error
	}{
		{
			name: "missing flags",
			input: func() *cobra.Command {
				cmd := &cobra.Command{Use: "test"}
				return cmd
			}(),
			expected: testFlags{},
			err:      &mapstructure.Error{Errors: []string{"'' has unset fields: boolean-flag, string-flag"}},
		},
		{
			name: "success",
			input: func() *cobra.Command {
				cmd := &cobra.Command{Use: "test"}
				cmd.Flags().Bool("boolean-flag", true, "")
				cmd.Flags().StringP("string-flag", "s", "value", "")
				return cmd
			}(),
			expected: testFlags{
				BooleanFlag: true,
				StringFlag:  "value",
			},
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.name)
		t.Run(name, func(t *testing.T) {
			actual, err := Parse(testFlags{}, test.input)

			assert.Equal(t, test.err, err)

			assert.Equal(t, test.expected, actual)
		})
	}
}
