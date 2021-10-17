package logger

import (
	"fmt"

	"github.com/aripalo/aws-mfa-credential-process/internal/data"
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
func format(d data.Provider, colorize color.Color, emoji string, prefix string, message string) string {
	c := d.GetConfig()
	output := ""

	if emoji != "" {
		if c.DisableColor == false {
			output = fmt.Sprintf("%s%s", output, emoji)
		}
	}

	if prefix != "" {
		var p string
		if c.DisableColor {
			p = prefix
		} else {
			p = textBold.Render(prefix)
		}
		output = fmt.Sprintf("%s %s", output, p)
	}

	output = fmt.Sprintf("%s %s", output, message)

	if c.DisableColor {
		return output
	} else {
		return colorize.Render(output)
	}
}

// Debugln prints a message with newline if verbose mode enabled (in dark gray if colors enabled).
func Debugln(d data.Provider, emoji string, prefix string, message string) {
	s := d.GetWriteStream()
	c := d.GetConfig()
	if c.Verbose {
		formatted := format(d, textColorDebug, emoji, prefix, message)
		fmt.Fprintln(s, formatted)
	}
}

// Debug prints a message if verbose mode enabled (in dark gray if colors enabled).
func Debug(d data.Provider, emoji string, prefix string, message string) {
	s := d.GetWriteStream()
	c := d.GetConfig()
	if c.Verbose {
		formatted := format(d, textColorDebug, emoji, prefix, message)
		fmt.Fprint(s, formatted)
	}
}

// Debugf prints a formatted message if verbose mode enabled (in dark gray if colors enabled).
func Debugf(d data.Provider, emoji string, prefix string, message string, args ...interface{}) {
	s := d.GetWriteStream()
	c := d.GetConfig()
	if c.Verbose {
		formatted := format(d, textColorDebug, emoji, prefix, message)
		fmt.Fprintf(s, formatted, args...)
	}
}

// Infoln prints a message with newline if verbose mode enabled (in gray if colors enabled).
func Infoln(d data.Provider, emoji string, prefix string, message string) {
	s := d.GetWriteStream()
	c := d.GetConfig()
	if c.Verbose {
		formatted := format(d, textColorInfo, emoji, prefix, message)
		fmt.Fprintln(s, formatted)
	}
}

// Info prints a message if verbose mode enabled (in gray if colors enabled).
func Info(d data.Provider, emoji string, prefix string, message string) {
	s := d.GetWriteStream()
	c := d.GetConfig()
	if c.Verbose {
		formatted := format(d, textColorInfo, emoji, prefix, message)
		fmt.Fprint(s, formatted)
	}
}

// Infof prints a formatted message if verbose mode enabled (in gray if colors enabled).
func Infof(d data.Provider, emoji string, prefix string, message string, args ...interface{}) {
	s := d.GetWriteStream()
	c := d.GetConfig()
	if c.Verbose {
		formatted := format(d, textColorInfo, emoji, prefix, message)
		fmt.Fprintf(s, formatted, args...)
	}
}

// Successln prints a message with newline if verbose mode enabled (in green if colors enabled).
func Successln(d data.Provider, emoji string, prefix string, message string) {
	s := d.GetWriteStream()
	c := d.GetConfig()
	if c.Verbose {
		formatted := format(d, textColorSuccess, emoji, prefix, message)
		fmt.Fprintln(s, formatted)
	}
}

// Success prints a message if verbose mode enabled (in green if colors enabled).
func Success(d data.Provider, emoji string, prefix string, message string) {
	s := d.GetWriteStream()
	c := d.GetConfig()
	if c.Verbose {
		formatted := format(d, textColorSuccess, emoji, prefix, message)
		fmt.Fprint(s, formatted)
	}
}

// Successf prints a formatted message if verbose mode enabled (in green if colors enabled).
func Successf(d data.Provider, emoji string, prefix string, message string, args ...interface{}) {
	s := d.GetWriteStream()
	c := d.GetConfig()
	if c.Verbose {
		formatted := format(d, textColorSuccess, emoji, prefix, message)
		fmt.Fprintf(s, formatted, args...)
	}
}

// Titleln prints a message with newline if verbose mode enabled (in magenta if colors enabled).
func Titleln(d data.Provider, emoji string, prefix string, message string) {
	s := d.GetWriteStream()
	c := d.GetConfig()
	if c.Verbose {
		formatted := format(d, textColorTitle, emoji, prefix, message)
		fmt.Fprintln(s, formatted)
	}
}

// Title prints a message if verbose mode enabled (in magenta if colors enabled).
func Title(d data.Provider, emoji string, prefix string, message string) {
	s := d.GetWriteStream()
	c := d.GetConfig()
	if c.Verbose {
		formatted := format(d, textColorTitle, emoji, prefix, message)
		fmt.Fprint(s, formatted)
	}
}

// Titlef prints a formatted message if verbose mode enabled (in magenta if colors enabled).
func Titlef(d data.Provider, emoji string, prefix string, message string, args ...interface{}) {
	s := d.GetWriteStream()
	c := d.GetConfig()
	if c.Verbose {
		formatted := format(d, textColorTitle, emoji, prefix, message)
		fmt.Fprintf(s, formatted, args...)
	}
}

// Importantln prints a message with newline (in yellow if colors enabled).
func Importantln(d data.Provider, emoji string, prefix string, message string) {
	s := d.GetWriteStream()
	formatted := format(d, textColorImportant, emoji, prefix, message)
	fmt.Fprintln(s, formatted)
}

// Important prints a message (in yellow if colors enabled).
func Important(d data.Provider, emoji string, prefix string, message string) {
	s := d.GetWriteStream()
	formatted := format(d, textColorImportant, emoji, prefix, message)
	fmt.Fprint(s, formatted)
}

// Importantf prints a formatted message (in yellow if colors enabled).
func Importantf(d data.Provider, emoji string, prefix string, message string, args ...interface{}) {
	s := d.GetWriteStream()
	formatted := format(d, textColorImportant, emoji, prefix, message)
	fmt.Fprintf(s, formatted, args...)
}

// Promptln prints a message with newline (in cyan if colors enabled).
func Promptln(d data.Provider, emoji string, prefix string, message string) {
	s := d.GetWriteStream()
	formatted := format(d, textColorPrompt, emoji, prefix, message)
	fmt.Fprintln(s, formatted)
}

// Prompt prints a message (in cyan if colors enabled).
func Prompt(d data.Provider, emoji string, prefix string, message string) {
	s := d.GetWriteStream()
	formatted := format(d, textColorPrompt, emoji, prefix, message)
	fmt.Fprint(s, formatted)
}

// Promptf prints a formatted message (in cyan if colors enabled).
func Promptf(d data.Provider, emoji string, prefix string, message string, args ...interface{}) {
	s := d.GetWriteStream()
	formatted := format(d, textColorPrompt, emoji, prefix, message)
	fmt.Fprintf(s, formatted, args...)
}

// Errorln prints a formatted message (in red if colors enabled).
func Errorln(d data.Provider, emoji string, prefix string, message string) {
	s := d.GetWriteStream()
	formatted := format(d, textColorError, emoji, prefix, message)
	fmt.Fprintln(s, formatted)
}

// Error prints a message (in red if colors enabled).
func Error(d data.Provider, emoji string, prefix string, message string) {
	s := d.GetWriteStream()
	formatted := format(d, textColorError, emoji, prefix, message)
	fmt.Fprint(s, formatted)
}

// Errorf prints a formatted message (in red if colors enabled).
func Errorf(d data.Provider, emoji string, prefix string, message string, args ...interface{}) {
	s := d.GetWriteStream()
	formatted := format(d, textColorError, emoji, prefix, message)
	fmt.Fprintf(s, formatted, args...)
}
