package logger

import (
	"bytes"
	"errors"
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
	args        []string
	fn          func(a interfaces.AssumeCredentialProcess, emoji string, prefix string, message string)
	fnF         func(a interfaces.AssumeCredentialProcess, emoji string, prefix string, message string, args ...interface{})
	want        string
}

func TestLogger(t *testing.T) {

	tests := []LoggerTestCase{
		{
			description: "non-verbose Debugln",
			flags: config.Flags{
				Verbose: false,
			},
			emoji:   "🚧",
			prefix:  "Test",
			message: "Message",
			fn:      Debugln,
			want:    "",
		},
		{
			description: "verbose Debugln",
			flags: config.Flags{
				Verbose: true,
			},
			emoji:   "🚧",
			prefix:  "Test",
			message: "Message",
			fn:      Debugln,
			want:    "🚧 \x1b[90m\x1b[1mTest:\x1b[0m\x1b[0m \x1b[90mMessage\x1b[0m\n",
		},
		{
			description: "non-verbose Debug",
			flags: config.Flags{
				Verbose: false,
			},
			emoji:   "🚧",
			prefix:  "Test",
			message: "Message",
			fn:      Debug,
			want:    "",
		},
		{
			description: "verbose Debug",
			flags: config.Flags{
				Verbose: true,
			},
			emoji:   "🚧",
			prefix:  "Test",
			message: "Message",
			fn:      Debug,
			want:    "🚧 \x1b[90m\x1b[1mTest:\x1b[0m\x1b[0m \x1b[90mMessage\x1b[0m",
		},
		{
			description: "non-verbose Debugf",
			flags: config.Flags{
				Verbose: false,
			},
			emoji:   "🚧",
			prefix:  "Test",
			message: "Message %s",
			args:    []string{"formatted"},
			fnF:     Debugf,
			want:    "",
		},
		{
			description: "verbose Debugf",
			flags: config.Flags{
				Verbose: true,
			},
			emoji:   "🚧",
			prefix:  "Test",
			message: "Message %s",
			args:    []string{"formatted"},
			fnF:     Debugf,
			want:    "🚧 \x1b[90m\x1b[1mTest:\x1b[0m\x1b[0m \x1b[90mMessage formatted\x1b[0m",
		},
		{
			description: "Infoln",
			flags:       config.Flags{},
			emoji:       "🚧",
			prefix:      "Test",
			message:     "Message",
			fn:          Infoln,
			want:        "🚧 \x1b[90m\x1b[1mTest:\x1b[0m\x1b[0m \x1b[90mMessage\x1b[0m\n",
		},
		{
			description: "Info",
			flags:       config.Flags{},
			emoji:       "🚧",
			prefix:      "Test",
			message:     "Message",
			fn:          Info,
			want:        "🚧 \x1b[90m\x1b[1mTest:\x1b[0m\x1b[0m \x1b[90mMessage\x1b[0m",
		},
		{
			description: "Infof",
			flags:       config.Flags{},
			emoji:       "🚧",
			prefix:      "Test",
			message:     "Message %s",
			args:        []string{"formatted"},
			fnF:         Infof,
			want:        "🚧 \x1b[90m\x1b[1mTest:\x1b[0m\x1b[0m \x1b[90mMessage formatted\x1b[0m",
		},

		{
			description: "Successln",
			flags:       config.Flags{},
			emoji:       "🚧",
			prefix:      "Test",
			message:     "Message",
			fn:          Successln,
			want:        "🚧 \x1b[32m\x1b[1mTest:\x1b[0m\x1b[0m \x1b[32mMessage\x1b[0m\n",
		},
		{
			description: "Success",
			flags:       config.Flags{},
			emoji:       "🚧",
			prefix:      "Test",
			message:     "Message",
			fn:          Success,
			want:        "🚧 \x1b[32m\x1b[1mTest:\x1b[0m\x1b[0m \x1b[32mMessage\x1b[0m",
		},
		{
			description: "Successf",
			flags:       config.Flags{},
			emoji:       "🚧",
			prefix:      "Test",
			message:     "Message %s",
			args:        []string{"formatted"},
			fnF:         Successf,
			want:        "🚧 \x1b[32m\x1b[1mTest:\x1b[0m\x1b[0m \x1b[32mMessage formatted\x1b[0m",
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

			if tc.fnF == nil && tc.fn == nil {
				panic(errors.New("No test function defined"))
			}

			if tc.fnF != nil {
				b := make([]interface{}, len(tc.args))
				for i := range tc.args {
					b[i] = tc.args[i]
				}

				tc.fnF(a, tc.emoji, tc.prefix, tc.message, b...)
			}

			if tc.fn != nil {
				tc.fn(a, tc.emoji, tc.prefix, tc.message)
			}

			got := output.String()

			if got != tc.want {
				t.Fatalf(`Got %q, want %q`, got, tc.want)
			}
		})
	}
}
