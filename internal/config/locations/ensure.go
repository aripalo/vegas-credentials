package locations

import (
	"os"
	"path/filepath"
)

// Creates a given directory under baseDir and returns the absolute path.
func EnsureWithinDir(baseDir string, dirName string) string {
	abs := filepath.Join(baseDir, dirName)
	err := os.MkdirAll(abs, os.ModePerm)
	if err != nil {
		panic(err)
	}
	return abs
}
