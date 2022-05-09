package assumecfg

import (
	"fmt"
	"os"
	"os/user"
	"regexp"
	"runtime"

	"github.com/aripalo/vegas-credentials/internal/msg"
)

const DurationSecondsDefault int = 3600

func resolveYubikeyLabel(yubikeyLabel string, mfaSerial string) string {
	if yubikeyLabel != "" {
		return yubikeyLabel
	}
	return mfaSerial
}

func resolveRegion(targetProfileRegion string, sourceProfileRegion string) string {
	if targetProfileRegion != "" {
		return targetProfileRegion
	}
	if sourceProfileRegion != "" {
		return sourceProfileRegion
	}
	return os.Getenv("AWS_REGION")
}

func resolveDurationSeconds(durationSeconds int) int {
	if durationSeconds > 0 {
		return durationSeconds
	}
	return DurationSecondsDefault
}

var minLength = 2

func resolveRoleSessionName(roleSessionName string) string {
	if len(roleSessionName) >= minLength {
		return roleSessionName
	}

	id := getSessionIdentifier()
	val := format(id)

	return val
}

func getSessionIdentifier() string {

	// Let's first try getting user info
	if u, err := user.Current(); err == nil {

		if u2, err := user.Lookup(u.Username); err == nil {

			msg.Debug("🔴", "uFullname:"+u.Name)
			msg.Debug("🔴", "u2Fullname:"+u2.Name)
			// Return user full name if meaningful
			if len(u2.Name) >= minLength {
				msg.Trace("", "Fallback: Fullname")
				return u2.Name
			}

		}

		// Return user (system) name if meaningful
		if len(u.Username) >= minLength {
			msg.Trace("", "Fallback: Username")
			return u.Username
		}
	}

	// If username not found, return hostname if meaningful
	if hostname, err := os.Hostname(); err == nil {
		if len(hostname) >= minLength {
			msg.Trace("", "Fallback: Hostname")
			return hostname
		}
	}

	msg.Trace("", "Fallback: OS_ARCH")

	// If no other value could not be resolved, use system info
	return fmt.Sprintf("%s_%s\n", runtime.GOOS, runtime.GOARCH)
}

func format(id string) string {
	cleaned := removeDisallowed(id)
	truncated := truncate(cleaned)

	// handle extremely rare situation where the cleaned string is too short
	if len(truncated) <= minLength {
		return fmt.Sprintf("mysession-%s", truncated)
	}

	return truncated
}

// https://aws.amazon.com/blogs/security/easily-control-naming-individual-iam-role-sessions/
var disallowed = `[^a-zA-Z0-9_=,.@-]`
var disallowedRegexp = regexp.MustCompile(disallowed)

func removeDisallowed(value string) string {
	return disallowedRegexp.ReplaceAllString(value, "")
}

func truncate(value string) string {
	if len(value) > 64 {
		return value[:64]
	}
	return value
}
