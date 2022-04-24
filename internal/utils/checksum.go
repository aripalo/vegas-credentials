package utils

import "encoding/json"

func CalculateChecksum(data any) (string, error) {
	stringified, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return SHA1(string(stringified))
}
