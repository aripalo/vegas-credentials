package assumeopts

import (
	"time"

	"github.com/aripalo/vegas-credentials/internal/sts"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
)

func (a *AssumeOpts) BuildAssumeRoleProvider(tokenProvider sts.TokenProvider) sts.AssumeRoleProvider {
	return func(assume *stscreds.AssumeRoleProvider) {

		// IAM MFA device ARN
		assume.SerialNumber = aws.String(a.MfaSerial)

		// Configures the temporary session duration
		assume.Duration = time.Duration(a.DurationSeconds) * time.Second

		// map our own MFA Token Provider signature to one expected by AWS Go SDK
		assume.TokenProvider = tokenProvider

		// ExternalID may not be empty string, or otherwise AWS Go SDK errors
		if a.ExternalID != "" {
			assume.ExternalID = aws.String(a.ExternalID)
		}

		// RoleSessionName may not be empty string, or otherwise AWS Go SDK errors
		if a.RoleSessionName != "" {
			assume.RoleSessionName = a.RoleSessionName
		}
	}
}
