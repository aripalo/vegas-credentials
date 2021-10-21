package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

// GenerateSHA1 reads string data and creates a SHA1 hash string in hex encoding from it
func GenerateSHA1(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash
}
