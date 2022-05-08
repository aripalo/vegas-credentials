package prettyjson

import (
	"encoding/json"
)

// Marshals structs into stringified JSON with indenting.
func Marshal(input any) (string, error) {
	output, err := json.MarshalIndent(input, "", "    ")
	if err != nil {
		return "", err
	}
	return string(output), nil
}
