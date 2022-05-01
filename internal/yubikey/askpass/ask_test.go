package askpass

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAskPass(t *testing.T) {
	tests := []struct {
		name      string
		enableGui bool
		expected  string
		cliPrompt func(ctx context.Context, text string) (string, error)
		guiPrompt func(ctx context.Context, title string, text string) (string, error)
		err       error
	}{
		{
			name:      "gui-disabled",
			enableGui: false,
			expected:  "p4ssword",
			cliPrompt: func(ctx context.Context, text string) (string, error) {
				return "p4ssword", nil
			},
			guiPrompt: func(ctx context.Context, title string, text string) (string, error) {
				return "should-not-return", nil
			},
		},
		{
			name:      "cli-wins",
			enableGui: true,
			expected:  "p4ssword",
			cliPrompt: func(ctx context.Context, text string) (string, error) {
				return "p4ssword", nil
			},
			guiPrompt: func(ctx context.Context, title string, text string) (string, error) {
				time.Sleep(time.Millisecond * 1000)
				return "should-not-return", nil
			},
		},
		{
			name:      "gui-wins",
			enableGui: true,
			expected:  "p4ssword",
			cliPrompt: func(ctx context.Context, text string) (string, error) {
				time.Sleep(time.Millisecond * 1000)
				return "should-not-return", nil
			},
			guiPrompt: func(ctx context.Context, title string, text string) (string, error) {
				return "p4ssword", nil
			},
		},
		{
			name:      "both-fail",
			enableGui: true,
			expected:  "",
			cliPrompt: func(ctx context.Context, text string) (string, error) {
				return "", errors.New("cli-fail")
			},
			guiPrompt: func(ctx context.Context, title string, text string) (string, error) {
				return "", errors.New("gui-fail")
			},
			err: context.DeadlineExceeded,
		},
		{
			name:      "timeout-exceeded",
			enableGui: true,
			expected:  "",
			cliPrompt: func(ctx context.Context, text string) (string, error) {
				time.Sleep(time.Millisecond * 1000)
				return "p4ssword", nil
			},
			guiPrompt: func(ctx context.Context, title string, text string) (string, error) {
				time.Sleep(time.Millisecond * 1000)
				return "p4ssword", nil
			},
			err: context.DeadlineExceeded,
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.name)
		t.Run(name, func(t *testing.T) {

			cliPrompt = test.cliPrompt
			guiPrompt = test.guiPrompt

			ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
			defer cancel()

			actual, err := AskPassword(ctx, test.enableGui)
			assert.Equal(t, test.err, err)
			assert.Equal(t, test.expected, actual)
		})
	}
}
