package flag

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

type testFlags struct {
	BooleanFlag bool   `mapstructure:"boolean-flag"`
	StringFlag  string `mapstructure:"string-flag"`
}

func Test(t *testing.T) {

	var testCmd = &cobra.Command{
		Use: "test",
	}

	testCmd.Flags().Bool("boolean-flag", true, "")
	testCmd.Flags().StringP("string-flag", "s", "value", "")

	name := "flag parsing"
	input := testCmd
	expected := testFlags{
		BooleanFlag: true,
		StringFlag:  "value",
	}

	t.Run(name, func(t *testing.T) {
		actual, err := Parse(testFlags{}, input)
		assert.Equal(t, err, nil)
		assert.Equal(t, expected, actual)
	})

}
