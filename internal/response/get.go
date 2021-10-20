package response

import (
	"time"

	"github.com/aripalo/aws-mfa-credential-process/internal/data"
	"github.com/aripalo/aws-mfa-credential-process/internal/logger"
	"github.com/aripalo/aws-mfa-credential-process/internal/mfa"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Get Temporary Session Credentials response for AWS credential_process
func (r *Response) Get(d data.Provider) error {

	var err error

	sess, err := getSession(d)
	if err != nil {
		return err
	}

	creds := defineCredentials(d, sess)

	rNew, err := getResponse(creds)
	if err != nil {
		return err
	}

	r.Version = rNew.Version
	r.AccessKeyID = rNew.AccessKeyID
	r.SecretAccessKey = rNew.SecretAccessKey
	r.SessionToken = rNew.SessionToken
	r.Expiration = rNew.Expiration

	logger.DebugJSON(d, "🔧 ", "New Response", r)

	return nil
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

	return stscreds.NewCredentials(sess, profile.RoleArn, func(p *stscreds.AssumeRoleProvider) {

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

// getResponse performs the assume role operation and returns response struct ready for printing
func getResponse(creds *credentials.Credentials) (*Response, error) {

	var r *Response
	var err error

	// Get() performs the actual assume role operation by fetching temporary session credentials
	value, err := creds.Get()
	if err != nil {
		return r, err
	}

	// Get temporary session expiration time, must be called after creds.Get()
	expiration, err := creds.ExpiresAt()
	if err != nil {
		return r, err
	}

	// format the response
	r = &Response{
		Version:         AWS_CREDENTIAL_PROCESS_VERSION,
		AccessKeyID:     value.AccessKeyID,
		SecretAccessKey: value.SecretAccessKey,
		SessionToken:    value.SessionToken,
		Expiration:      expiration,
	}

	return r, err
}
