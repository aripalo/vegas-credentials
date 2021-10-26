package logger

import (
	"fmt"

	"github.com/aripalo/vegas-credentials/internal/interfaces"
	"github.com/gookit/color"
)

const (
	textBold           = color.OpBold
	textColorDebug     = color.FgDarkGray
	textColorInfo      = color.FgGray
	textColorImportant = color.FgYellow
	textColorError     = color.FgRed
	textColorSuccess   = color.FgGreen
	textColorTitle     = color.FgLightMagenta
	textColorPrompt    = color.FgCyan
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
		output = fmt.Sprintf("%s%s", output, p)
	}

	if f.NoColor {
		output = fmt.Sprintf("%s %s", output, message)
	} else {
		output = fmt.Sprintf("%s %s", output, colorize.Render(message))
	}

	return output
}

// Newline prints a newline character
func Newline(a interfaces.AssumeCredentialProcess) {
	s := a.GetDestination()
	fmt.Fprintln(s)
}

// Debugln prints a message with newline if verbose mode enabled (in dark gray if colors enabled).
func Debugln(a interfaces.AssumeCredentialProcess, emoji string, prefix string, message string) {
	s := a.GetDestination()
	f := a.GetFlags()
	if f.Verbose {
		formatted := format(a, textColorDebug, emoji, prefix, message)
		fmt.Fprintln(s, formatted)
	}
}

// Debug prints a message if verbose mode enabled (in dark gray if colors enabled).
func Debug(a interfaces.AssumeCredentialProcess, emoji string, prefix string, message string) {
	s := a.GetDestination()
	f := a.GetFlags()
	if f.Verbose {
		formatted := format(a, textColorDebug, emoji, prefix, message)
		fmt.Fprint(s, formatted)
	}
}

// Debugf prints a formatted message if verbose mode enabled (in dark gray if colors enabled).
func Debugf(a interfaces.AssumeCredentialProcess, emoji string, prefix string, message string, args ...interface{}) {
	s := a.GetDestination()
	f := a.GetFlags()
	if f.Verbose {
		formatted := format(a, textColorDebug, emoji, prefix, message)
		fmt.Fprintf(s, formatted, args...)
	}
}

// Infoln prints a message with newline if verbose mode enabled (in gray if colors enabled).
func Infoln(a interfaces.AssumeCredentialProcess, emoji string, prefix string, message string) {
	s := a.GetDestination()

	formatted := format(a, textColorInfo, emoji, prefix, message)
	fmt.Fprintln(s, formatted)

}

// Info prints a message if verbose mode enabled (in gray if colors enabled).
func Info(a interfaces.AssumeCredentialProcess, emoji string, prefix string, message string) {
	s := a.GetDestination()

	formatted := format(a, textColorInfo, emoji, prefix, message)
	fmt.Fprint(s, formatted)

}

// Infof prints a formatted message if verbose mode enabled (in gray if colors enabled).
func Infof(a interfaces.AssumeCredentialProcess, emoji string, prefix string, message string, args ...interface{}) {
	s := a.GetDestination()

	formatted := format(a, textColorInfo, emoji, prefix, message)
	fmt.Fprintf(s, formatted, args...)

}

// Successln prints a message with newline if verbose mode enabled (in green if colors enabled).
func Successln(a interfaces.AssumeCredentialProcess, emoji string, prefix string, message string) {
	s := a.GetDestination()

	formatted := format(a, textColorSuccess, emoji, prefix, message)
	fmt.Fprintln(s, formatted)

}

// Success prints a message if verbose mode enabled (in green if colors enabled).
func Success(a interfaces.AssumeCredentialProcess, emoji string, prefix string, message string) {
	s := a.GetDestination()

	formatted := format(a, textColorSuccess, emoji, prefix, message)
	fmt.Fprint(s, formatted)

}

// Successf prints a formatted message if verbose mode enabled (in green if colors enabled).
func Successf(a interfaces.AssumeCredentialProcess, emoji string, prefix string, message string, args ...interface{}) {
	s := a.GetDestination()

	formatted := format(a, textColorSuccess, emoji, prefix, message)
	fmt.Fprintf(s, formatted, args...)

}

// Titleln prints a message with newline if verbose mode enabled (in magenta if colors enabled).
func Titleln(a interfaces.AssumeCredentialProcess, emoji string, prefix string, message string) {
	s := a.GetDestination()
	f := a.GetFlags()
	if f.Verbose {
		formatted := format(a, textColorTitle, emoji, prefix, message)
		fmt.Fprintln(s, formatted)
	}
}

// Title prints a message if verbose mode enabled (in magenta if colors enabled).
func Title(a interfaces.AssumeCredentialProcess, emoji string, prefix string, message string) {
	s := a.GetDestination()
	f := a.GetFlags()
	if f.Verbose {
		formatted := format(a, textColorTitle, emoji, prefix, message)
		fmt.Fprint(s, formatted)
	}
}

// Titlef prints a formatted message if verbose mode enabled (in magenta if colors enabled).
func Titlef(a interfaces.AssumeCredentialProcess, emoji string, prefix string, message string, args ...interface{}) {
	s := a.GetDestination()
	f := a.GetFlags()
	if f.Verbose {
		formatted := format(a, textColorTitle, emoji, prefix, message)
		fmt.Fprintf(s, formatted, args...)
	}
}

// Importantln prints a message with newline (in yellow if colors enabled).
func Importantln(a interfaces.AssumeCredentialProcess, emoji string, prefix string, message string) {
	s := a.GetDestination()
	formatted := format(a, textColorImportant, emoji, prefix, message)
	fmt.Fprintln(s, formatted)
}

// Important prints a message (in yellow if colors enabled).
func Important(a interfaces.AssumeCredentialProcess, emoji string, prefix string, message string) {
	s := a.GetDestination()
	formatted := format(a, textColorImportant, emoji, prefix, message)
	fmt.Fprint(s, formatted)
}

// Importantf prints a formatted message (in yellow if colors enabled).
func Importantf(a interfaces.AssumeCredentialProcess, emoji string, prefix string, message string, args ...interface{}) {
	s := a.GetDestination()
	formatted := format(a, textColorImportant, emoji, prefix, message)
	fmt.Fprintf(s, formatted, args...)
}

// Promptln prints a message with newline (in cyan if colors enabled).
func Promptln(a interfaces.AssumeCredentialProcess, emoji string, prefix string, message string) {
	s := a.GetDestination()
	formatted := format(a, textColorPrompt, emoji, prefix, message)
	fmt.Fprintln(s, formatted)
}

// Prompt prints a message (in cyan if colors enabled).
func Prompt(a interfaces.AssumeCredentialProcess, emoji string, prefix string, message string) {
	s := a.GetDestination()
	formatted := format(a, textColorPrompt, emoji, prefix, message)
	fmt.Fprint(s, formatted)
}

// Promptf prints a formatted message (in cyan if colors enabled).
func Promptf(a interfaces.AssumeCredentialProcess, emoji string, prefix string, message string, args ...interface{}) {
	s := a.GetDestination()
	formatted := format(a, textColorPrompt, emoji, prefix, message)
	fmt.Fprintf(s, formatted, args...)
}

// Errorln prints a formatted message (in red if colors enabled).
func Errorln(a interfaces.AssumeCredentialProcess, emoji string, prefix string, message string) {
	s := a.GetDestination()
	formatted := format(a, textColorError, emoji, prefix, message)
	fmt.Fprintln(s, formatted)
}

// Error prints a message (in red if colors enabled).
func Error(a interfaces.AssumeCredentialProcess, emoji string, prefix string, message string) {
	s := a.GetDestination()
	formatted := format(a, textColorError, emoji, prefix, message)
	fmt.Fprint(s, formatted)
}

// Errorf prints a formatted message (in red if colors enabled).
func Errorf(a interfaces.AssumeCredentialProcess, emoji string, prefix string, message string, args ...interface{}) {
	s := a.GetDestination()
	formatted := format(a, textColorError, emoji, prefix, message)
	fmt.Fprintf(s, formatted, args...)
}
