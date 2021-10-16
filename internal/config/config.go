package config

import (
	"os"
	"path/filepath"
)

const (
	// PRODUCT_NAME defines product name used in various outputs
	PRODUCT_NAME string = "aws-mfa-credential-process"

	PRODUCT_CONFIG_LOCATION = "." + PRODUCT_NAME
)

// Init creates the main config folder with keyring folder
func Init() {

	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	configPath := filepath.Join(homedir, PRODUCT_CONFIG_LOCATION)

	os.MkdirAll(configPath, os.ModePerm)
}
