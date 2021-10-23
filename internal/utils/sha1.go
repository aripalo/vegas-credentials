package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

// GenerateSHA1 reads string data and creates a SHA1 hash string in hex encoding from it
func GenerateSHA1(data string) (string, error) {
	h := sha1.New()
	_, err := h.Write([]byte(data))
	if err != nil {
		return "", err
	}
	hash := hex.EncodeToString(h.Sum(nil))
	return hash, nil
}
