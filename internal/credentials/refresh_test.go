package credentials

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var refreshNotNeeded int = MandatoryRefreshTimeout + 1
var refreshNeeded int = MandatoryRefreshTimeout - 1

func TestIsRefreshNeeded(t *testing.T) {
	tests := []struct {
		name     string
		input    Credentials
		expected bool
	}{
		{
			name: "not needed",
			input: Credentials{
				Version:         1,
				AccessKeyID:     "ID",
				SecretAccessKey: "SECRET",
				SessionToken:    "TOKEN",
				Expiration:      time.Now().Add(time.Second * time.Duration(refreshNotNeeded)),
			},
			expected: false,
		},
		{
			name: "needed",
			input: Credentials{
				Version:         1,
				AccessKeyID:     "ID",
				SecretAccessKey: "SECRET",
				SessionToken:    "TOKEN",
				Expiration:      time.Now().Add(time.Second * time.Duration(refreshNeeded)),
			},
			expected: true,
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.name)
		t.Run(name, func(t *testing.T) {
			actual := test.input.isRefreshNeeded()
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestSecondsRemaining(t *testing.T) {
	tests := []struct {
		name     string
		input    Credentials
		expected int
	}{
		{
			name: "now",
			input: Credentials{
				Version:         1,
				AccessKeyID:     "ID",
				SecretAccessKey: "SECRET",
				SessionToken:    "TOKEN",
				Expiration:      time.Now(),
			},
			expected: 0,
		},
		{
			name: "60s remaining",
			input: Credentials{
				Version:         1,
				AccessKeyID:     "ID",
				SecretAccessKey: "SECRET",
				SessionToken:    "TOKEN",
				Expiration:      time.Now().Add(time.Second * 60),
			},
			expected: 59,
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.name)
		t.Run(name, func(t *testing.T) {
			actual := test.input.secondsRemaining()
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestIsExpired(t *testing.T) {
	tests := []struct {
		name     string
		input    Credentials
		now      time.Time
		expected bool
	}{
		{
			name: "not expired",
			input: Credentials{
				Version:         1,
				AccessKeyID:     "ID",
				SecretAccessKey: "SECRET",
				SessionToken:    "TOKEN",
				Expiration:      time.Now().Add(time.Second * 1),
			},
			now:      time.Now(),
			expected: false,
		},
		{
			name: "expired",
			input: Credentials{
				Version:         1,
				AccessKeyID:     "ID",
				SecretAccessKey: "SECRET",
				SessionToken:    "TOKEN",
				Expiration:      time.Now().Add(time.Second * -1),
			},
			now:      time.Now(),
			expected: true,
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.name)
		t.Run(name, func(t *testing.T) {
			actual := test.input.isExpired(test.now)
			assert.Equal(t, test.expected, actual)
		})
	}
}
