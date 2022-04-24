package utils

import (
	"encoding/json"
)

func PrettyJSON(input interface{}) (string, error) {
	output, err := json.MarshalIndent(input, "", "    ")
	if err != nil {
		return "", err
	}
	return string(output), nil
}
