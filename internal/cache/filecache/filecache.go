package filecache

import (
	"bytes"
	"compress/gzip"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/aripalo/aws-mfa-credential-process/internal/profile"
	"github.com/aripalo/aws-mfa-credential-process/internal/utils"
)

func Get(profileName string, config profile.Profile) ([]byte, error) {
	var err error
	filepath, err := getCachePath(profileName, config)
	cached, err := os.ReadFile(filepath)
	decompressed, err := decompress(cached)
	return decompressed, err
}

func Save(profileName string, config profile.Profile, data []byte) error {
	var err error
	compressed, err := compress(data)
	filepath, err := getCachePath(profileName, config)
	err = os.WriteFile(filepath, compressed, os.ModePerm)
	if err != nil {
		utils.SafeLogLn("FILE ERR: ", err)
	}
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
	toolDirName  string = ".awstool-todo" // TODO
	cacheDirName string = "cache"
)

func formatFileName(profileName string, config profile.Profile) (string, error) {
	configString, err := configToString(config)
	combination := fmt.Sprintf("%s%s", profileName, configString)
	hash := generateSha1Hash([]byte(combination))
	return hash, err
}

func generateSha1Hash(data []byte) string {
	s := string(data)
	h := sha1.New()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}

func configToString(config profile.Profile) (string, error) {
	result, err := json.Marshal(config)
	return string(result), err
}

func compress(data []byte) ([]byte, error) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(data); err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func decompress(compressed []byte) ([]byte, error) {
	reader := bytes.NewReader(compressed)
	gz, err := gzip.NewReader(reader)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(gz)
	if err != nil {
		return nil, err
	}

	return data, nil
}
