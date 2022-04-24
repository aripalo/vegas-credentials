package locations

import (
	"os"
	"path/filepath"
)

// Get AWS config file which usually is under .aws folder in user home directory,
// unless user has set $AWS_CONFIG_FILE environment variable.
// https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html#cli-configure-files-where
func getAwsConfigFile() (string, error) {
	if val := os.Getenv("AWS_CONFIG_FILE"); val != "" {
		return val, nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".aws", "config"), nil
}
