package credentials

import (
	"encoding/json"
	"fmt"

	"github.com/aripalo/vegas-credentials/internal/prettyjson"
)

// Serialize response into pretty JSON
func (r *Credentials) Serialize() (json.RawMessage, error) {
	data, err := prettyjson.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("Serialization Error: %s", err.Error())
	}
	return []byte(data), nil
}

// Deserialize JSON into Response struct
func (r *Credentials) Deserialize(value json.RawMessage) error {
	err := json.Unmarshal(value, &r)
	if err != nil {
		return fmt.Errorf("Deserialization Error: %s", err.Error())
	}
	return nil
}
