package locations

import (
	"demo/internal/config"
	"os"
	"path/filepath"
)

// Ensure a directory with app name exists in given baseDir path.
func mustEnsureAppDir(baseDir string) string {
	dir := filepath.Join(baseDir, config.AppName)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		panic(err)
	}
	return dir
}
