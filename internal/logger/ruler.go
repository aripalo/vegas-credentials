package logger

import (
	"errors"
	"fmt"

	"github.com/aripalo/vegas-credentials/internal/data"
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

// PrintRuler repeats a character as many times as thera are columns in the terminal (if verbose mode)
func PrintRuler(d data.Provider, char string) {
	c := d.GetConfig()
	if c.Verbose {
		ruler := createRuler(char)
		s := d.GetWriteStream()
		c := d.GetConfig()
		if c.NoColor {
			fmt.Fprintln(s, ruler)
		} else {
			fmt.Fprintln(s, textColorDebug.Render(ruler))
		}
	}
}
