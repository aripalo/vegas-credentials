package config

import (
	"os"
	"path/filepath"
)

// TempDir denotes a directory where to store temporary files.
// OS specific. On Unix-based machines follows $TMPDIR.
var TempDir string = func(appName string) string {
	dir := filepath.Join(os.TempDir(), appName)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		panic(err)
	}
	return dir
}(APP_NAME)

// TempFilePath returns a full path for given filename for temporary use.
// The path should not change between consecutive executions.
// On Unix-based machines follows $TMPDIR.
func TempFilePath(filename string) string {
	return filepath.Join(TempDir, filename)
}
