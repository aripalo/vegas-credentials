package prettyjson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	BooleanFlag bool
	StringFlag  string
}

func Test(t *testing.T) {

	name := "pretty json marshal"
	input := testData{
		BooleanFlag: true,
		StringFlag:  "value",
	}
	expected := "{\n    \"BooleanFlag\": true,\n    \"StringFlag\": \"value\"\n}"

	t.Run(name, func(t *testing.T) {
		actual, err := Marshal(input)
		assert.Equal(t, err, nil)
		assert.Equal(t, expected, actual)
	})

}
