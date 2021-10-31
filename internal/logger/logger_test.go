package logger

import (
	"bytes"
	"os"
	"testing"

	"github.com/aripalo/vegas-credentials/internal/config"
	"github.com/aripalo/vegas-credentials/internal/interfaces"
	"github.com/aripalo/vegas-credentials/internal/profile"
	"github.com/aripalo/vegas-credentials/internal/vegastestapp"
	"github.com/gookit/color"
)

type LoggerTestCase struct {
	description string
	flags       config.Flags
	emoji       string
	prefix      string
	message     string
	args        []interface{}
	fn          func(a interfaces.AssumeCredentialProcess, emoji string, prefix string, message string, args ...interface{})
	want        string
}

func TestLogger(t *testing.T) {

	tests := []LoggerTestCase{
		{
			description: "non-verbose debugf",
			flags: config.Flags{
				Verbose: false,
			},
			emoji:   "🚧",
			prefix:  "Test",
			message: "Message",
			fn:      Debugf,
			want:    "",
		},
		{
			description: "verbose debugf",
			flags: config.Flags{
				Verbose: true,
			},
			emoji:   "🚧",
			prefix:  "Test",
			message: "Message",
			fn:      Debugf,
			want:    "🚧 \x1b[90m\x1b[1mTest:\x1b[0m\x1b[0m \x1b[90mMessage\x1b[0m",
		},
	}

	for _, tc := range tests {

		// Handle terminal env (i.e. in CI)
		nocolor := os.Getenv("NO_COLOR")
		term := os.Getenv("TERM")
		level := color.TermColorLevel()
		os.Unsetenv("NO_COLOR")
		os.Setenv("TERM", "xterm-256color")
		os.Setenv("FORCE_COLOR", "on")
		_ = color.ForceSetColorLevel(color.Level256)

		defer func() {
			os.Setenv("NO_COLOR", nocolor)
			os.Setenv("TERM", term)
			os.Unsetenv("FORCE_COLOR")
			color.ForceSetColorLevel(level)
		}()

		t.Run(tc.description, func(t *testing.T) {
			var output bytes.Buffer

			a := &vegastestapp.AssumeAppForTesting{
				Flags:       tc.flags,
				Profile:     profile.Profile{},
				Destination: &output,
			}

			tc.fn(a, tc.emoji, tc.prefix, tc.message, tc.args...)

			got := output.String()

			if got != tc.want {
				t.Fatalf(`Got %q, want %q`, got, tc.want)
			}
		})
	}
}
