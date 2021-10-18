package logger

import (
	"errors"
	"fmt"

	"github.com/aripalo/aws-mfa-credential-process/internal/data"
	"golang.org/x/term"
)

func getWidth() (int, error) {
	if term.IsTerminal(0) {
		width, _, err := term.GetSize(0)
		if err != nil {
			return 0, err
		}
		return width, nil
	} else {
		return 0, errors.New("Not a terminal")
	}
}

func createRuler(char string) string {
	width, err := getWidth()

	if err != nil || width == 0 {
		width = 16
	}

	banner := ""
	for i := 0; i < width; i++ {
		banner += char
	}

	return banner
}

func printRuler(d data.Provider, char string) {
	ruler := createRuler(char)
	s := d.GetWriteStream()
	c := d.GetConfig()
	if c.NoColor {
		fmt.Fprintln(s, ruler)
	} else {
		fmt.Fprintln(s, textColorDebug.Render(ruler))
	}
}
