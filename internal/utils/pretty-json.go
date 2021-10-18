package utils

import (
	"encoding/json"
)

func PrettyJSON(input interface{}) string {

	output, err := json.MarshalIndent(input, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(output)

}
