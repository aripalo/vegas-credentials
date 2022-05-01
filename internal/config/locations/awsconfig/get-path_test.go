package awsconfig

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPath(t *testing.T) {

	t.Run("default unix", func(t *testing.T) {
		actual, err := GetPath()
		homeDir, _ := os.UserHomeDir()
		assert.Equal(t, err, nil)
		assert.Equal(t, filepath.Join(homeDir, ".aws", "config"), actual)
	})

	t.Run("$AWS_CONFIG_FILE set", func(t *testing.T) {
		customLocation := "/tmp/custom/config"
		os.Setenv("AWS_CONFIG_FILE", customLocation)
		actual, err := GetPath()
		assert.Equal(t, err, nil)
		assert.Equal(t, customLocation, actual)
	})

}
