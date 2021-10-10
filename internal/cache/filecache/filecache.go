package filecache

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aripalo/goawsmfa/internal/profile"
)

func Get(profileName string, config profile.Profile) ([]byte, error) {
	var err error
	filepath, err := getCachePath(profileName, config)
	cached, err := os.ReadFile(filepath)
	return cached, err
}

func Save(profileName string, config profile.Profile, data []byte) error {
	var err error
	filepath, err := getCachePath(profileName, config)
	err = os.WriteFile(filepath, data, os.ModePerm)
	return err
}

func Remove(profileName string, config profile.Profile) error {
	var err error
	filepath, err := getCachePath(profileName, config)
	err = os.Remove(filepath)
	return err
}

func getAndEnsureCachePath() (string, error) {
	homedir, err := os.UserHomeDir()
	cachePath := filepath.Join(homedir, toolDirName, cacheDirName)
	err = os.MkdirAll(cachePath, os.ModePerm)
	return cachePath, err
}

func getCachePathForKey(key string) (string, error) {
	cachePath, err := getAndEnsureCachePath()
	filePath := filepath.Join(cachePath, key)
	return filePath, err
}

func getCachePath(profileName string, config profile.Profile) (string, error) {
	filename, err := formatFileName(profileName, config)
	filepath, err := getCachePathForKey(filename)
	return filepath, err
}

const (
	toolDirName  string = ".awstool-todo"
	cacheDirName string = "cache"
)

func formatFileName(profileName string, config profile.Profile) (string, error) {
	configString, err := configToString(config)
	combination := fmt.Sprintf("%s%s", profileName, configString)
	filename := b64.StdEncoding.EncodeToString([]byte(combination))
	shortened := filename[:200]
	return shortened, err
}

func configToString(config profile.Profile) (string, error) {
	result, err := json.Marshal(config)
	return string(result), err
}
