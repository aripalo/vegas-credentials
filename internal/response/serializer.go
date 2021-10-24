package response

import (
	"encoding/json"
	"fmt"

	"github.com/aripalo/vegas-credentials/internal/utils"
)

// Serialize response into pretty JSON
func (r *Response) Serialize() (json.RawMessage, error) {
	data, err := utils.PrettyJSON(r)
	if err != nil {
		return nil, fmt.Errorf("Serialization Error: %s", err.Error())
	}
	return []byte(data), nil
}

// Deserialize JSON into Response struct
func (r *Response) Deserialize(value json.RawMessage) error {
	err := json.Unmarshal(value, &r)
	if err != nil {
		return fmt.Errorf("Deserialization Error: %s", err.Error())
	}
	return nil
}
