package response

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aripalo/aws-mfa-credential-process/internal/utils"
)

// Serialize response into pretty JSON
func (r *Response) Serialize() (json.RawMessage, error) {
	data, err := utils.PrettyJSON(r)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Serialization Error: %s", err.Error()))
	}
	return []byte(data), nil
}

// Deserialize JSON into Response struct
func (r *Response) Deserialize(value json.RawMessage) error {
	err := json.Unmarshal(value, &r)
	if err != nil {
		return errors.New(fmt.Sprintf("Deserialization Error: %s", err.Error()))
	}
	return nil
}
