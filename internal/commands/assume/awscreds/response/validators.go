package response

import (
	"errors"
	"fmt"
	"time"

	"github.com/aripalo/aws-mfa-credential-process/internal/data"
	"github.com/dustin/go-humanize"
)

// Validate ensures the response is of correct format
func (r *Response) Validate(d data.Provider) error {

	if r.Version != AWS_CREDENTIAL_PROCESS_VERSION {
		return errors.New("Incorrect Version")
	}

	if r.AccessKeyID == "" {
		return errors.New("Missing AccessKeyID")
	}

	if r.SecretAccessKey == "" {
		return errors.New("Missing SecretAccessKey")
	}

	if r.SessionToken == "" {
		return errors.New("Missing SessionToken")
	}

	now := time.Now()

	if r.Expiration.Before(now) {
		return errors.New(fmt.Sprintf("Expired %s\n", humanize.RelTime(r.Expiration, now, "ago", "in future")))
	}

	return nil
}

// ValidateForMandatoryRefresh ensures response is within "mandatory refresh" duration as per BotoCore
// https://github.com/boto/botocore/blob/221ffa67a567df99ee78d7ae308c0e12d7eeeea7/botocore/credentials.py#L350-L355
func (r *Response) ValidateForMandatoryRefresh(d data.Provider) error {

	c := d.GetConfig()

	if c.DisableRefresh {
		return nil
	}

	now := time.Now()
	count := 10 * 60
	limit := now.Add(time.Duration(-count) * time.Second)

	if r.Expiration.Before(limit) {
		return errors.New(fmt.Sprintf("Mandatory refresh required because expiration in %s\n", humanize.Time(r.Expiration)))
	}

	return nil
}
