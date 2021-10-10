package utils

import (
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/gookit/color"
)

func FormatMessage(messageColor color.Color, emoji string, prefix string, message string) string {
	return fmt.Sprintf("%s%s %s", color.OpBold.Render(emoji), messageColor.Render(color.OpBold.Render(prefix+":")), messageColor.Render(message))
}

const (
	COLOR_DEBUG     = color.FgDarkGray
	COLOR_INFO      = color.FgGray
	COLOR_IMPORTANT = color.FgYellow
	COLOR_ERROR     = color.FgRed
	COLOR_SUCCESS   = color.FgGreen
	COLOR_TITLE     = color.FgLightMagenta
)

func FormatExpirationMessage(expiration time.Time) string {
	return fmt.Sprintf("Valid for ~%s (until %s)", humanize.Time(expiration), expiration.Local().Format("2006-01-02 15:04:05 MST"))
}
