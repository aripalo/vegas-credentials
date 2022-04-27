package sts

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
)

type TokenProvider func() (string, error)
type AssumeRoleProvider func(assume *stscreds.AssumeRoleProvider)

type GetCredentialsResponse struct {
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
	Expiration      time.Time
}

func GetCredentials(
	profile string,
	region string,
	roleArn string,
	provider AssumeRoleProvider,

) (GetCredentialsResponse, error) {
	var response GetCredentialsResponse

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewSharedCredentials("", profile),
	})
	if err != nil {
		return response, err
	}

	creds := stscreds.NewCredentials(sess, roleArn, provider)

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

	response = GetCredentialsResponse{
		AccessKeyID:     value.AccessKeyID,
		SecretAccessKey: value.SecretAccessKey,
		SessionToken:    value.SessionToken,
		Expiration:      expiration,
	}

	return response, nil
}
