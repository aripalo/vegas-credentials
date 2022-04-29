package checksum

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
)

// Generate reads data and creates a Generate hash string in hex encoding from it
func Generate(input any) (string, error) {

	data, err := json.Marshal(input)
	if err != nil {
		return "", err
	}

	h := sha1.New()
	_, err = h.Write(data)
	if err != nil {
		return "", err
	}
	hash := hex.EncodeToString(h.Sum(nil))
	return hash, nil
}
