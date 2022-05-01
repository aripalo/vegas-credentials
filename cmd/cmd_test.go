package cmd

import (
	"fmt"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tests := []struct {
		name       string
		input      *cobra.Command
		existsRunE bool
		version    string
	}{
		{
			name:       "vegas-credentials",
			input:      rootCmd,
			existsRunE: false,
			version:    "v0.0.0-development",
		},
		{
			name:       "assume",
			input:      assumeCmd,
			existsRunE: true,
		},
		{
			name:       "config",
			input:      configCmd,
			existsRunE: false,
		},
		{
			name:       "list",
			input:      configListCmd,
			existsRunE: true,
		},
		{
			name:       "show-profile",
			input:      configShowProfileCmd,
			existsRunE: true,
		},
		{
			name:       "version",
			input:      versionCmd,
			existsRunE: true,
		},
		{
			name:       "cache",
			input:      cacheCmd,
			existsRunE: false,
		},
		{
			name:       "clean",
			input:      cacheCleanCmd,
			existsRunE: true,
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.name)
		t.Run(name, func(t *testing.T) {

			assert.Equal(t, test.name, test.input.Name())
			assert.NotNil(t, test.input.PreRun)
			assert.NotNil(t, test.input.PostRun)
			if test.existsRunE {
				assert.NotNil(t, test.input.RunE)
			}
			if test.version != "" {
				assert.Equal(t, test.version, test.input.Version)
			}
			assert.Greater(t, len(test.input.Short), 6)
		})
	}
}
