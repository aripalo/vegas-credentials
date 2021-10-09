package mfa

import (
	"errors"
	"regexp"
)

var tokenPattern = regexp.MustCompile("\\d{6}\\d*")

func validateToken(token string) error {
	result := tokenPattern.Match([]byte(token))
	if result != true {
		return errors.New("Invalid token")
	}
	return nil
}
