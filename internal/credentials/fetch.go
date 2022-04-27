package credentials

import (
	"errors"
	"fmt"
	"time"

	"github.com/aripalo/vegas-credentials/internal/sts"

	"github.com/dgraph-io/badger/v3"
	"github.com/dustin/go-humanize"
)

func (c *Credentials) FetchFromCache() error {

	err := c.readFromCache()
	if err != nil {
		if err == badger.ErrKeyNotFound {
			return errors.New("Not found")
		}
		return err
	}

	err = c.Validate()
	if err != nil {
		return err
	}

	now := time.Now()

	if c.isExpired(now) {
		return fmt.Errorf("Expired %s", humanize.RelTime(c.Expiration, now, "ago", "in future"))
	}

	if c.isRefreshNeeded() {
		return fmt.Errorf("Refresh required because expiration in %s", humanize.Time(c.Expiration))
	}

	return nil
}

func (c *Credentials) FetchFromAWS() error {

	r, err := sts.GetCredentials(sts.Request{
		Profile:  c.opts.SourceProfile,
		Region:   c.opts.Region,
		RoleArn:  c.opts.RoleArn,
		Provider: c.opts.AssumeRoleProvider,
	})

	if err != nil {
		return err
	}

	c.AccessKeyID = r.AccessKeyID
	c.SecretAccessKey = r.SecretAccessKey
	c.SessionToken = r.SessionToken
	c.Expiration = r.Expiration

	err = c.Validate()
	if err != nil {
		return err
	}

	err = c.saveToCache()
	if err != nil {
		return err
	}

	return nil
}
