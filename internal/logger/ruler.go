package logger

import (
	"errors"

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

func CreateRuler(char string) string {
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
