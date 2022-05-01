package cache

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		ttl        time.Duration
		sleep      time.Duration
		expected   string
		errMessage string
	}{
		{
			name:     "works",
			input:    "some-value",
			ttl:      time.Millisecond * 1000,
			expected: "some-value",
		},
		{
			name:       "ttl-expired",
			input:      "some-value",
			ttl:        time.Millisecond * 500,
			sleep:      time.Millisecond * 501,
			expected:   "",
			errMessage: "Key not found",
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.name)
		t.Run(name, func(t *testing.T) {

			tempDir, err := ioutil.TempDir("", "vegas-credentials-cache-test")
			if err != nil {
				log.Fatal(err)
			}

			defer os.RemoveAll(tempDir)

			c := New(tempDir)

			err = c.Set(test.name, []byte(test.input), test.ttl)
			require.NoError(t, err)

			time.Sleep(test.sleep)

			result, err := c.Get(test.name)
			if test.errMessage != "" {
				assert.Equal(t, test.errMessage, err.Error())
			}

			actual := string(result)

			assert.Equal(t, test.expected, actual)

			err = c.Remove(test.name)
			require.NoError(t, err)
		})
	}
}
