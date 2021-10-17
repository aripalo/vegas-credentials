package awscreds

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
)

type Credentials struct {
	AccessKeyID     string `json:"AccessKeyId"`
	SecretAccessKey string `json:"SecretAccessKey"`
}

type TemporaryCredentials struct {
	Credentials
	SessionToken string    `json:"SessionToken"`
	Expiration   time.Time `json:"Expiration"`
}

type CredentialProcessOutput struct {
	TemporaryCredentials
	Version int `json:"Version"`
}

func (o *CredentialProcessOutput) GetOutput(creds *credentials.Credentials) error {
	var err error

	o.Version = 1

	if creds.IsExpired() {
		creds.Expire()
	}

	expiration, err := creds.ExpiresAt()
	value, err := creds.Get()

	o.AccessKeyID = value.AccessKeyID
	o.SecretAccessKey = value.SecretAccessKey
	o.SessionToken = value.SessionToken
	o.Expiration = expiration

	return err
}
