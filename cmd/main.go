package main

import (
	"os"

	"github.com/aripalo/aws-mfa-credential-process/internal/config"
	"github.com/aripalo/aws-mfa-credential-process/internal/credentialprocess"
	"github.com/aripalo/aws-mfa-credential-process/internal/profile"
	"github.com/aripalo/aws-mfa-credential-process/internal/securestorage"
	"github.com/aripalo/aws-mfa-credential-process/internal/utils"
	"github.com/urfave/cli/v2"
)

var Debug = true

func main() {

	app := &cli.App{
		Name:   "aws-mfa-credential-process",
		Usage:  "Caching AWS Credential Process to manage assuming an IAM Role with MFA token from Yubikey and Authenticator App",
		Flags:  config.FlagsConfiguration,
		Action: mainAction,
	}

	err := app.Run(os.Args)
	if err != nil {
		utils.PrintError(err)
		os.Exit(1)
	}

}

func mainAction(c *cli.Context) error {

	flags := config.ParseFlags(c)

	if flags.Verbose {
		utils.PrintBanner()
	}

	config.Init()
	securestorage.Init(flags)

	var err error

	profileConfig, err := profile.GetProfile(flags.ProfileName)
	if err != nil {
		//utils.SafeLog(err)
		return err
	}

	output, err := credentialprocess.GetOutput(flags, profileConfig)
	if err != nil {
		//utils.SafeLog(err)
		return err
	}

	utils.OutputToAwsCredentialProcess(string(output))

	return err
}
