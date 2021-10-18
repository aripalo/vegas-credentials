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

// Response defines the output format expected by AWS credential_process
// https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sourcing-external.html
type Response struct {
	Version         int       `json:"Version"`
	AccessKeyID     string    `json:"AccessKeyId"`
	SecretAccessKey string    `json:"SecretAccessKey"`
	SessionToken    string    `json:"SessionToken"`
	Expiration      time.Time `json:"Expiration"`
}

// getSession provides AWS session used to assume the target role
func getSession(d data.Provider) (*session.Session, error) {
	profile := d.GetProfile()
	return session.NewSession(&aws.Config{
		Region:      aws.String(profile.Region),
		Credentials: credentials.NewSharedCredentials("", profile.SourceProfile),
	})
}

// defineCredentials configures how the target role is assumed
func defineCredentials(d data.Provider, sess *session.Session) *credentials.Credentials {
	profile := d.GetProfile()
	return stscreds.NewCredentials(sess, profile.AssumeRoleArn, func(p *stscreds.AssumeRoleProvider) {

		// IAM MFA device ARN
		p.SerialNumber = aws.String(profile.MfaSerial)

		// Configures the temporary session duration
		p.Duration = time.Duration(profile.DurationSeconds) * time.Second

		// map our own MFA Token Provider signature to one expected by AWS Go SDK
		p.TokenProvider = func() (string, error) {
			result, err := mfa.GetToken(d)
			if err != nil {
				return "", err
			}
			return result.Value, nil
		}

		// ExternalID may not be empty, or otherwise AWS Go SDK errors
		if profile.ExternalID != "" {
			p.ExternalID = aws.String(profile.ExternalID)
		}

		// RoleSessionName may not be empty, or otherwise AWS Go SDK errors
		if profile.RoleSessionName != "" {
			p.RoleSessionName = profile.RoleSessionName
		}
	})
}

// getCredentialProcessResponse performs the assume role operation and returns response struct ready for printing
func getCredentialProcessResponse(creds *credentials.Credentials) (*Response, error) {
	var response *Response
	var err error

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

	// format the response
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
func CredentialProcess(d data.Provider) error {
	profile := d.GetProfile()

	logger.Infoln(d, "👷", "Role", profile.AssumeRoleArn)

	sess, err := getSession(d)

	if err != nil {
		return err
	}

	creds := defineCredentials(d, sess)

	response, err := getCredentialProcessResponse(creds)
	if err != nil {
		return err
	}

	logger.DebugJSON(d, "🔧 ", "Response", response)

	print(response)

	return err
}

// Print credential_process combatible JSON into stdout
func print(response *Response) error {
	output, err := utils.PrettyJSON(response)
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
