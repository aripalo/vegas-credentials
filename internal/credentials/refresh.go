package credentials

import (
	"time"
)

// The time at which all threads will block waiting for refreshed credentials.
// As defined by Botocore:
// https://github.com/boto/botocore/blob/221ffa67a567df99ee78d7ae308c0e12d7eeeea7/botocore/credentials.py#L355
const MandatoryRefreshTimeout int = 10 * 60

// Checks if credentials need to be refreshed as defined by Botocore:
// https://github.com/boto/botocore/blob/221ffa67a567df99ee78d7ae308c0e12d7eeeea7/botocore/credentials.py#L353-L355
func (c *Credentials) isRefreshNeeded() bool {
	return c.secondsRemaining() < MandatoryRefreshTimeout
}

// Returns seconds remaining until credentials are expired.
func (c *Credentials) secondsRemaining() int {
	now := time.Now()
	diff := c.Expiration.Sub(now)
	return int(diff.Seconds())
}

// Checks if credentials already expired.
func (c *Credentials) isExpired(now time.Time) bool {
	return c.Expiration.Before(now)
}
