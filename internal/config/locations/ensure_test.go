package locations

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/aripalo/vegas-credentials/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestLocatioEnsure(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		tmpDir := os.TempDir()
		appDir := filepath.Join(tmpDir, "vegas-credentials-testing-appdir")
		_ = os.Remove(appDir)
		actual := EnsureWithinDir(appDir, config.AppName)
		assert.Equal(t, filepath.Join(appDir, config.AppName), actual)

		_, err := os.Stat(appDir)
		assert.Equal(t, err, nil)

		err = os.RemoveAll(appDir)
		assert.Equal(t, err, nil)

		_, err = os.Stat(appDir)
		assert.Equal(t, os.IsNotExist(err), true)
	})

}
