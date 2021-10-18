package awscreds

import (
	"fmt"
	"os"
	"time"

	"github.com/aripalo/aws-mfa-credential-process/internal/application/assume/awscreds/mfa"
	"github.com/aripalo/aws-mfa-credential-process/internal/data"
	"github.com/aripalo/aws-mfa-credential-process/internal/logger"
	"github.com/aripalo/aws-mfa-credential-process/internal/utils"
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
	response *Response
}

func getSession(d data.Provider) (*session.Session, error) {
	profile := d.GetProfile()
	return session.NewSession(&aws.Config{
		Region:      aws.String(profile.Region),
		Credentials: credentials.NewSharedCredentials("", profile.SourceProfile),
	})
}

func defineCredentials(d data.Provider, sess *session.Session) *credentials.Credentials {
	profile := d.GetProfile()
	return stscreds.NewCredentials(sess, profile.AssumeRoleArn, func(p *stscreds.AssumeRoleProvider) {
		p.SerialNumber = aws.String(profile.MfaSerial)

		p.Duration = time.Duration(profile.DurationSeconds) * time.Second

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
}

func getCredentialProcessResponse(creds *credentials.Credentials) (*Response, error) {
	var response *Response
	var err error

	value, err := creds.Get()
	if err != nil {
		return response, err
	}

	expiration, err := creds.ExpiresAt()
	if err != nil {
		return response, err
	}

	response = &Response{
		Version:         1,
		AccessKeyID:     value.AccessKeyID,
		SecretAccessKey: value.SecretAccessKey,
		SessionToken:    value.SessionToken,
		Expiration:      expiration,
	}

	return response, err
}

// New initializes session credentials
func (c *CredentialProcess) New(d data.Provider) (*CredentialProcess, error) {
	profile := d.GetProfile()

	credentialprocess := CredentialProcess{}

	logger.Infoln(d, "👷", "Role", profile.AssumeRoleArn)

	sess, err := getSession(d)

	if err != nil {
		return &credentialprocess, err
	}

	creds := defineCredentials(d, sess)

	response, err := getCredentialProcessResponse(creds)
	if err != nil {
		return &credentialprocess, err
	}

	c.response = response

	return &credentialprocess, err
}

// Print credential_process combatible JSON into stdout
func (c *CredentialProcess) Print() error {
	output, err := utils.PrettyJSON(c.response)
	if err != nil {
		return err
	}

	outputToAwsCredentialProcess(output)

	return nil
}

// OutputToAwsCredentialProcess prints to stdout so aws credential_process can read it
// https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sourcing-external.html
func outputToAwsCredentialProcess(output string) {
	fmt.Fprintf(os.Stdout, output)
}
