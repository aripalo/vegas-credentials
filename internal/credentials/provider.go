package credentials

import (
	"time"

	"github.com/aripalo/vegas-credentials/internal/credentials/sts"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
)

func (c *Credentials) BuildProvider(tokenProvider sts.TokenProvider) sts.Provider {
	return func(assume *stscreds.AssumeRoleProvider) {

		// IAM MFA device ARN
		assume.SerialNumber = aws.String(c.cfg.MfaSerial)

		// Configures the temporary session duration
		assume.Duration = time.Duration(c.cfg.DurationSeconds) * time.Second

		// map our own MFA Token Provider signature to one expected by AWS Go SDK
		assume.TokenProvider = tokenProvider

		// ExternalID may not be empty string, or otherwise AWS Go SDK errors
		if c.cfg.ExternalID != "" {
			assume.ExternalID = aws.String(c.cfg.ExternalID)
		}

		// RoleSessionName may not be empty string, or otherwise AWS Go SDK errors
		if c.cfg.RoleSessionName != "" {
			assume.RoleSessionName = c.cfg.RoleSessionName
		}
	}
}
