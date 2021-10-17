package credentialprocess

import (
	"encoding/json"
	"time"

	"github.com/aripalo/aws-mfa-credential-process/internal/config"
	"github.com/aripalo/aws-mfa-credential-process/internal/mfa"
	"github.com/aripalo/aws-mfa-credential-process/internal/profile"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
)

func getFreshTemporaryCredentials(flags config.CredentialProcessFlags, profileConfig profile.Profile) (json.RawMessage, error) {
	var err error

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(profileConfig.Region),
		Credentials: credentials.NewSharedCredentials("", profileConfig.SourceProfile),
	})

	// Create the credentials from AssumeRoleProvider to assume the role
	// referenced by the "myRoleARN" ARN. Prompt for MFA token from stdin.
	creds := stscreds.NewCredentials(sess, profileConfig.AssumeRoleArn, func(p *stscreds.AssumeRoleProvider) {
		p.SerialNumber = aws.String(profileConfig.MfaSerial)
		p.TokenProvider = func() (string, error) { return tokenProvider(flags, profileConfig) }

		if profileConfig.DurationSeconds != 0 {
			p.Duration = time.Duration(profileConfig.DurationSeconds * int(time.Second))
		}

		if profileConfig.RoleSessionName != "" {
			p.RoleSessionName = profileConfig.RoleSessionName
		}

		if profileConfig.ExternalID != "" {
			p.ExternalID = aws.String(profileConfig.ExternalID)
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

func tokenProvider(flags config.CredentialProcessFlags, profileConfig profile.Profile) (string, error) {
	result, err := mfa.GetTokenResult(flags, profileConfig)
	return result.Value, err
}

func toPrettyJson(data interface{}) (json.RawMessage, error) {
	pretty, err := json.MarshalIndent(data, "", "    ")
	return pretty, err
}
