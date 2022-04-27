package credentials

import (
	"time"

	"github.com/aripalo/vegas-credentials/internal/sts"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
)

func (c *Credentials) buildProvider(tokenProvider sts.TokenProvider) sts.Provider {
	return func(assume *stscreds.AssumeRoleProvider) {

		// IAM MFA device ARN
		assume.SerialNumber = aws.String(c.opts.MfaSerial)

		// Configures the temporary session duration
		assume.Duration = time.Duration(c.opts.DurationSeconds) * time.Second

		// map our own MFA Token Provider signature to one expected by AWS Go SDK
		assume.TokenProvider = tokenProvider

		// ExternalID may not be empty string, or otherwise AWS Go SDK errors
		if c.opts.ExternalID != "" {
			assume.ExternalID = aws.String(c.opts.ExternalID)
		}

		// RoleSessionName may not be empty string, or otherwise AWS Go SDK errors
		if c.opts.RoleSessionName != "" {
			assume.RoleSessionName = c.opts.RoleSessionName
		}
	}
}
