package actions

import (
	"github.com/aripalo/aws-mfa-credential-process/internal/config"
	"github.com/aripalo/aws-mfa-credential-process/internal/credentialprocess"
	"github.com/aripalo/aws-mfa-credential-process/internal/profile"
	"github.com/aripalo/aws-mfa-credential-process/internal/utils"
	"github.com/urfave/cli/v2"
)

// Assume performs the AWS assume role operation with MFA
func Assume(c *cli.Context) error {
	flags := config.ParseCredentialProcessFlags(c)
	sharedInitialization(c.Command.Name, flags.Verbose, flags.DisableDialog)

	utils.SafeLogLn("VERBOSE", flags.Verbose)

	var err error

	// profileConfig holds the configuration data read from ~/.aws/config for given profile
	profileConfig, err := profile.GetProfile(flags.ProfileName)
	if err != nil {
		return err
	}

	// sessionCredentials is a JSON representation of AWS Temporary Session Credentials for the assumed role
	// matches https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sourcing-external.html
	sessionCredentials, err := credentialprocess.GetOutput(flags, profileConfig)
	if err != nil {
		return err
	}

	// Print Temporary Session Credentials into stdout as AWS tools expect for credential_process
	// https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sourcing-external.html
	utils.OutputToAwsCredentialProcess(sessionCredentials)

	return err
}
