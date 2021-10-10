package credentialprocess

import (
	"encoding/json"
	"time"

	"github.com/aripalo/goawsmfa/internal/mfa"
	"github.com/aripalo/goawsmfa/internal/profile"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
)

func getFreshTemporaryCredentials(config profile.Profile) (json.RawMessage, error) {
	var err error

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(config.Region),
		Credentials: credentials.NewSharedCredentials("", config.SourceProfile),
	})

	// Create the credentials from AssumeRoleProvider to assume the role
	// referenced by the "myRoleARN" ARN. Prompt for MFA token from stdin.
	creds := stscreds.NewCredentials(sess, config.AssumeRoleArn, func(p *stscreds.AssumeRoleProvider) {
		p.SerialNumber = aws.String(config.MfaSerial)
		p.TokenProvider = func() (string, error) { return tokenProvider(config) }

		if config.DurationSeconds != 0 {
			p.Duration = time.Duration(*aws.Int(config.DurationSeconds))
		}

		if config.RoleSessionName != "" {
			p.RoleSessionName = config.RoleSessionName
		}

		if config.ExternalID != "" {
			p.ExternalID = aws.String(config.ExternalID)
		}
	})

	credentials, err := creds.Get()
	if err != nil {
		return nil, err
	}

	expiration, err := creds.ExpiresAt()

	response := &CredentialProcessResponse{
		Version:         1,
		AccessKeyID:     credentials.AccessKeyID,
		SecretAccessKey: credentials.SecretAccessKey,
		SessionToken:    credentials.SessionToken,
		Expiration:      expiration,
	}

	pretty, err := toPrettyJson(response)

	return pretty, err
}

func tokenProvider(config profile.Profile) (string, error) {
	result, err := mfa.GetTokenResult(config.YubikeySerial, config.YubikeyLabel)
	return result.Value, err
}

func toPrettyJson(data interface{}) (json.RawMessage, error) {
	pretty, err := json.MarshalIndent(data, "", "    ")
	return pretty, err
}
