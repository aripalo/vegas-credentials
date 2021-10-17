package awscreds

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/aripalo/aws-mfa-credential-process/internal/application/assume/awscreds/mfa"
	"github.com/aripalo/aws-mfa-credential-process/internal/data"
	"github.com/aripalo/aws-mfa-credential-process/internal/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
)

type Response struct {
	Version         int       `json:"Version"`
	AccessKeyID     string    `json:"AccessKeyId"`
	SecretAccessKey string    `json:"SecretAccessKey"`
	SessionToken    string    `json:"SessionToken"`
	Expiration      time.Time `json:"Expiration"`
}

type CredentialProcess struct {
	response    *Response
	credentials *credentials.Credentials
}

// New initializes session credentials
func (c *CredentialProcess) New(d data.Provider) (*CredentialProcess, error) {
	profile := d.GetProfile()

	credentialprocess := CredentialProcess{}

	//c.response = Response{}
	//c.credentials = credentials.Credentials{}
	//var creds *credentials.Credentials

	logger.Infoln(d, "👷", "Role", profile.AssumeRoleArn)

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(profile.Region),
		Credentials: credentials.NewSharedCredentials("", profile.SourceProfile),
	})

	if err != nil {
		return &credentialprocess, err
	}

	// Create the credentials from AssumeRoleProvider to assume the role
	// referenced by the "myRoleARN" ARN. Prompt for MFA token from stdin.
	creds := stscreds.NewCredentials(sess, profile.AssumeRoleArn, func(p *stscreds.AssumeRoleProvider) {
		p.SerialNumber = aws.String(profile.MfaSerial)
		p.Duration = time.Duration(profile.DurationSeconds * int(time.Second))

		p.TokenProvider = func() (string, error) {
			result, err := mfa.GetToken(d)
			if err != nil {
				return "", err
			}
			return result.Value, nil
		}

		if profile.ExternalID != "" {
			p.ExternalID = aws.String(profile.ExternalID)
		}

		if profile.RoleSessionName != "" {
			p.RoleSessionName = profile.RoleSessionName
		}

	})

	credentialprocess.credentials = creds

	return &credentialprocess, err
}

// Get Temporary Session Credentials
func (c *CredentialProcess) Get() error {
	var err error

	if c.credentials.IsExpired() {
		c.credentials.Expire()
	}
	// TODO mandatory refresh (if not disabled)

	expiration, err := c.credentials.ExpiresAt()
	value, err := c.credentials.Get()

	c.response = &Response{
		Version:         1,
		AccessKeyID:     value.AccessKeyID,
		SecretAccessKey: value.SecretAccessKey,
		SessionToken:    value.SessionToken,
		Expiration:      expiration,
	}

	return err
}

// Print credential_process combatible JSON into stdout
func (c *CredentialProcess) Print() error {
	pretty, err := json.MarshalIndent(c.response, "", "    ")
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, string(pretty))
	return nil
}
