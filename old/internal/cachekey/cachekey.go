package cachekey

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"strings"

	"github.com/aripalo/aws-mfa-credential-process/internal/profile"
)

const Separator = "__"

// Get is responsible for creating a unique cache key for given profile configuration, therefore ensuring mutated profile configuration will not use previous cached data
func Get(profileName string, config profile.Profile) (string, error) {
	configString, err := configToString(config)
	hash := generateSha1Hash(configString)
	key := combineStrings(profileName, Separator, hash)
	return key, err
}

// generateSha1Hash reads byte array of data and creates a SHA1 hash string from it
func generateSha1Hash(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}

// combineStrings combines two strings
func combineStrings(items ...string) string {
	return strings.Join(items, "")
}

// configToString convertts profile config into stringified JSON
func configToString(config profile.Profile) (string, error) {
	result, err := json.Marshal(config)
	return string(result), err
}
