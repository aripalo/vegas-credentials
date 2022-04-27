package assumeopts

import "github.com/aripalo/vegas-credentials/internal/utils"

// Checksum calculates a SHA1 hash from the current configuration.
// Useful for detecting configuration changes e.g. for caching purposes.
func (a *AssumeOpts) Checksum() (string, error) {
	return utils.CalculateChecksum(a)
}
