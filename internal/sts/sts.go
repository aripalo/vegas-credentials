package sts

import (
	"time"

	"github.com/aripalo/aws-mfa-credential-process/internal/data"
	"github.com/aripalo/aws-mfa-credential-process/internal/mfa"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
)

// GetAssumedCredentials retuns a set of temporary session credentials for the assumed role and an expiration time
func GetAssumedCredentials(d data.Provider) (credentials.Value, time.Time, error) {
	var value credentials.Value
	var expiration time.Time

	sess, err := getSession(d)
	if err != nil {
		return value, expiration, err
	}
	creds := getCredentials(d, sess)

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
func getSession(d data.Provider) (*session.Session, error) {
	p := d.GetProfile()
	return session.NewSession(&aws.Config{
		Region:      aws.String(p.Region),
		Credentials: credentials.NewSharedCredentials("", p.SourceProfile),
	})
}

// getCredentials configures how the target role is assumed
func getCredentials(d data.Provider, sess *session.Session) *credentials.Credentials {
	p := d.GetProfile()
	return stscreds.NewCredentials(sess, p.RoleArn, func(assume *stscreds.AssumeRoleProvider) {

		// IAM MFA device ARN
		assume.SerialNumber = aws.String(p.MfaSerial)

		// Configures the temporary session duration
		assume.Duration = time.Duration(p.DurationSeconds) * time.Second

		// map our own MFA Token Provider signature to one expected by AWS Go SDK
		assume.TokenProvider = func() (string, error) {
			result, err := mfa.GetToken(d)
			if err != nil {
				return "", err
			}
			return result.Value, nil
		}

		// ExternalID may not be empty, or otherwise AWS Go SDK errors
		if p.ExternalID != "" {
			assume.ExternalID = aws.String(p.ExternalID)
		}

		// RoleSessionName may not be empty, or otherwise AWS Go SDK errors
		if p.RoleSessionName != "" {
			assume.RoleSessionName = p.RoleSessionName
		}
	})
}
