package logger

import (
	"io"
	"testing"

	"github.com/aripalo/vegas-credentials/internal/config"
	"github.com/aripalo/vegas-credentials/internal/profile"
	"github.com/aripalo/vegas-credentials/internal/vegastestapp"
)

type formatTestCase struct {
	description string
	flags       config.Flags
	emoji       string
	prefix      string
	message     string
	want        string
}

func TestFormat(t *testing.T) {

	tests := []formatTestCase{
		{
			"with all formatting",
			config.Flags{
				NoColor: false,
			},
			"🚧",
			"Test",
			"Message",
			"🚧 \x1b[90m\x1b[1mTest:\x1b[0m\x1b[0m \x1b[90mMessage\x1b[0m",
		},
		{
			"witout emoji",
			config.Flags{
				NoColor: false,
			},
			"",
			"Test",
			"Message",
			"\x1b[90m\x1b[1mTest:\x1b[0m\x1b[0m \x1b[90mMessage\x1b[0m",
		},
		{
			"witout prefix",
			config.Flags{
				NoColor: false,
			},
			"🚧",
			"",
			"Message",
			"🚧 \x1b[90mMessage\x1b[0m",
		},
		{
			"without color",
			config.Flags{
				NoColor: true,
			},
			"🚧",
			"Test",
			"Message",
			"Test: Message",
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			a := &vegastestapp.AssumeAppForTesting{
				Flags:       tc.flags,
				Profile:     profile.Profile{},
				Destination: io.Discard,
			}

			got := format(a, textColorDebug, tc.emoji, tc.prefix, tc.message)

			if got != tc.want {
				t.Fatalf(`Got %q, want %q`, got, tc.want)
			}
		})
	}

}
