package totp

import "regexp"

// String value representing a regexp that matches valid OATH TOTP token code.
// Essentially 6 digits or more.
const oathTotpRegexpString string = `^\d{6}\d*$`

// Validates the received value looks like a TOTP MFA Token Code:
// 6 digits or more. Final validation done by AWS STS.
func isValidToken(value string) bool {
	return regexp.MustCompile(oathTotpRegexpString).MatchString(value)
}
