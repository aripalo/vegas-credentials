package sts

import (
	"time"

	"github.com/aripalo/vegas-credentials/internal/interfaces"
	"github.com/aripalo/vegas-credentials/internal/mfa"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
)

// GetAssumedCredentials retuns a set of temporary session credentials for the assumed role and an expiration time
func GetAssumedCredentials(a interfaces.AssumeCredentialProcess) (credentials.Value, time.Time, error) {
	var value credentials.Value
	var expiration time.Time

	sess, err := getSession(a)
	if err != nil {
		return value, expiration, err
	}
	creds := getCredentials(a, sess)

	// Get() performs the actual assume role operation by fetching temporary session credentials
	value, err = creds.Get()
	if err != nil {
		return value, expiration, err
	}

	// Get temporary session expiration time, must be called after creds.Get()
	expiration, err = creds.ExpiresAt()
	if err != nil {
		return value, expiration, err
	}

	return value, expiration, nil
}

// getSession provides AWS session used to assume the target role
func getSession(a interfaces.AssumeCredentialProcess) (*session.Session, error) {
	p := a.GetProfile()
	return session.NewSession(&aws.Config{
		Region:      aws.String(p.Source.Region),
		Credentials: credentials.NewSharedCredentials("", p.Source.Name),
	})
}

// getCredentials configures how the target role is assumed
func getCredentials(a interfaces.AssumeCredentialProcess, sess *session.Session) *credentials.Credentials {
	p := a.GetProfile()
	return stscreds.NewCredentials(sess, p.Target.RoleArn, func(assume *stscreds.AssumeRoleProvider) {

		// IAM MFA device ARN
		assume.SerialNumber = aws.String(p.Source.MfaSerial)

		// Configures the temporary session duration
		assume.Duration = time.Duration(p.Target.DurationSeconds) * time.Second

		// map our own MFA Token Provider signature to one expected by AWS Go SDK
		assume.TokenProvider = func() (string, error) {
			result, err := mfa.GetToken(a)
			if err != nil {
				return "", err
			}
			return result.Value, nil
		}

		// ExternalID may not be empty, or otherwise AWS Go SDK errors
		if p.Target.ExternalID != "" {
			assume.ExternalID = aws.String(p.Target.ExternalID)
		}

		// RoleSessionName may not be empty, or otherwise AWS Go SDK errors
		if p.Target.RoleSessionName != "" {
			assume.RoleSessionName = p.Target.RoleSessionName
		}
	})
}
