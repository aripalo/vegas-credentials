package assumable

import (
	"os"
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
