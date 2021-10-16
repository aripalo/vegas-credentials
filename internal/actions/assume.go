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

	var err error

	profileConfig, err := profile.GetProfile(flags.ProfileName)
	if err != nil {
		return err
	}

	output, err := credentialprocess.GetOutput(flags, profileConfig)
	if err != nil {
		return err
	}

	utils.OutputToAwsCredentialProcess(string(output))

	return err
}
