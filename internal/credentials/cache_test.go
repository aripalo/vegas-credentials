package credentials

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	"github.com/aripalo/vegas-credentials/internal/assumecfg"
	"github.com/aripalo/vegas-credentials/internal/cache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {

	tempDir, err := ioutil.TempDir("", "vegas-credentials-cache-test-credentials")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(tempDir)

	c := cache.New(tempDir)

	cfg := assumecfg.AssumeCfg{
		ProfileName: "foo",
		Checksum:    "bar",
	}

	tests := []struct {
		name  string
		input Credentials
	}{
		{
			name: "",
			input: Credentials{
				repo:            c,
				cfg:             cfg,
				Version:         1,
				AccessKeyID:     "ID",
				SecretAccessKey: "SECRET",
				SessionToken:    "TOKEN",
				Expiration:      time.Now().Add(time.Minute * 5),
			},
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.name)
		t.Run(name, func(t *testing.T) {

			fmt.Println("RUNNING")

			err = test.input.saveToCache()
			require.NoError(t, err)

			fmt.Println("SAVED")

			actual := &Credentials{repo: c, cfg: cfg}

			fmt.Println(actual)

			err = actual.readFromCache()

			fmt.Println("READ")

			require.NoError(t, err)

			assert.Equal(t, test.input.Version, actual.Version)
			assert.Equal(t, test.input.AccessKeyID, actual.AccessKeyID)
			assert.Equal(t, test.input.SecretAccessKey, actual.SecretAccessKey)
			assert.Equal(t, test.input.SessionToken, actual.SessionToken)

			t1 := test.input.Expiration.Format(time.RFC3339)
			t2 := actual.Expiration.Format(time.RFC3339)

			assert.Equal(t, t1, t2)
		})
	}
}
