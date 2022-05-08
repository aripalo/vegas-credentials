package sts

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Function responsible of assuming the IAM Role.
type Provider func(assume *stscreds.AssumeRoleProvider)

// Function called once STS requires OATH TOTP MFA Token
// during the AssumeRoleProvider execution.
type TokenProvider func() (string, error)

// Required configuration to request STS Credentials.
type Request struct {
	Profile  string
	Region   string
	RoleArn  string
	Provider Provider
}

// STS Credentials response.
type Response struct {
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
	Expiration      time.Time
}

// Request STS Credentials.
func GetCredentials(request Request) (Response, error) {
	var response Response

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(request.Region),
		Credentials: credentials.NewSharedCredentials("", request.Profile),
	})
	if err != nil {
		return response, err
	}

	creds := stscreds.NewCredentials(sess, request.RoleArn, request.Provider)

	// Get() performs the actual assume role operation by fetching temporary session credentials
	value, err := creds.Get()
	if err != nil {
		return response, err
	}

	// Get temporary session expiration time, must be called after creds.Get()
	expiration, err := creds.ExpiresAt()
	if err != nil {
		return response, err
	}

	response = Response{
		AccessKeyID:     value.AccessKeyID,
		SecretAccessKey: value.SecretAccessKey,
		SessionToken:    value.SessionToken,
		Expiration:      expiration,
	}

	return response, nil
}
