package awscreds

import (
	"time"

	"github.com/aripalo/aws-mfa-credential-process/internal/provider"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
)

func Get(p provider.Provider) (*credentials.Credentials, error) {
	profile := p.GetProfile()

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(profile.Region),
		Credentials: credentials.NewSharedCredentials("", profile.SourceProfile),
	})

	if err != nil {
		return nil, err
	}

	// Create the credentials from AssumeRoleProvider to assume the role
	// referenced by the "myRoleARN" ARN. Prompt for MFA token from stdin.
	creds := stscreds.NewCredentials(sess, profile.AssumeRoleArn, func(p *stscreds.AssumeRoleProvider) {
		p.SerialNumber = aws.String(profile.MfaSerial)
		//p.TokenProvider = func() (string, error) { return tokenProvider(flags, profileConfig) }
		p.Duration = time.Duration(profile.DurationSeconds * int(time.Second))
		p.RoleSessionName = profile.RoleSessionName
		p.ExternalID = aws.String(profile.ExternalID)

	})

	return creds, nil
}
