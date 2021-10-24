// vegastest provides test helper utilities, not real functionality
package vegastest

import (
	"os"
	"path/filepath"
)

func GetCurrentDirectory() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return cwd
}

func GetTestdataFilePath(filename string) string {
	cwd := GetCurrentDirectory()
	return filepath.Join(cwd, "testdata", filename)
}
