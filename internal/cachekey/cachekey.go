package cachekey

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/aripalo/aws-mfa-credential-process/internal/profile"
)

func Get(profileName string, config profile.Profile) (string, error) {
	return formatFileName(profileName, config)
}

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
