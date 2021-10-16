package cachekey

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/aripalo/aws-mfa-credential-process/internal/profile"
)

// Get is responsible for creating a unique cache key for given profile configuration, therefore ensuring mutated profile configuration will not use previous cached data
func Get(profileName string, config profile.Profile) (string, error) {
	configString, err := configToString(config)
	combination := combineStrings(profileName, configString)
	hash := generateSha1Hash(combination)
	return hash, err
}

// generateSha1Hash reads byte array of data and creates a SHA1 hash string from it
func generateSha1Hash(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}

// combineStrings combines two strings
func combineStrings(a string, b string) string {
	return fmt.Sprintf("%s%s", a, b)
}

// configToString convertts profile config into stringified JSON
func configToString(config profile.Profile) (string, error) {
	result, err := json.Marshal(config)
	return string(result), err
}
