package mutex

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sync"
	"testing"

	"github.com/alexflint/go-filemutex"
	"github.com/stretchr/testify/require"
)

func TestLock(t *testing.T) {
	tests := []struct {
		name     string
		count    int
		input    string
		expected []int
	}{
		{
			name:     "",
			count:    4,
			input:    "",
			expected: []int{1, 2, 3, 4},
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.name)
		t.Run(name, func(t *testing.T) {

			dir = filepath.Join(os.TempDir(), "vegas-credentials-testing-lock")
			err := os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				panic(err)
			}

			var wg sync.WaitGroup
			wg.Add(test.count)

			for i := 0; i < test.count; i++ {
				go func() {
					defer wg.Done()

					unlock, err := Lock()
					require.NoError(t, err)

					fm2, err := filemutex.New(path.Join(dir, fileName))
					require.NoError(t, err)
					err = fm2.TryLock()
					require.Equal(t, AlreadyLocked, err)

					err = unlock()
					require.NoError(t, err)
				}()
			}

			wg.Wait()
		})
	}
}

var AlreadyLocked = errors.New("lock already acquired")
