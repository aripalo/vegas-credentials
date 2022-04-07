package provider

import (
	"fmt"
	"regexp"
)

func validateToken(token string) error {
	result := tokenValidatePattern.MatchString(token)
	if !result {
		return fmt.Errorf("Invalid OATH TOTP MFA Token Code: %q", token)
	}
	return nil
}

var tokenValidatePattern = regexp.MustCompile(`^\d{6}\d*$`)
