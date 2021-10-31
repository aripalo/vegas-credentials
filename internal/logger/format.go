package logger

import (
	"fmt"

	"github.com/aripalo/vegas-credentials/internal/interfaces"
	"github.com/gookit/color"
)

// format the log message
func format(a interfaces.AssumeCredentialProcess, colorize color.Color, emoji string, prefix string, message string) string {
	f := a.GetFlags()
	output := ""

	if emoji != "" && !f.NoColor {
		output = fmt.Sprintf("%s%s ", output, emoji)
	}

	if prefix != "" {
		var p string
		if f.NoColor {
			p = fmt.Sprintf("%s:", prefix)
		} else {
			p = colorize.Render(textBold.Render(fmt.Sprintf("%s:", prefix)))
		}
		output = fmt.Sprintf("%s%s ", output, p)
	}

	if f.NoColor {
		output = fmt.Sprintf("%s%s", output, message)
	} else {
		output = fmt.Sprintf("%s%s", output, colorize.Render(message))
	}

	return output
}
