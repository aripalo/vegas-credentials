package main

import (
	"encoding/json"
	"os"

	"github.com/aripalo/aws-mfa-credential-process/internal/config"
	"github.com/aripalo/aws-mfa-credential-process/internal/credentialprocess"
	"github.com/aripalo/aws-mfa-credential-process/internal/profile"
	"github.com/aripalo/aws-mfa-credential-process/internal/utils"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "aws-mfa-credential-process",
		Usage: "Caching AWS Credential Process to manage assuming an IAM Role with MFA token from Yubikey and Authenticator App",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Required: true,
				Name:     config.FLAG_PROFILE_NAME,
				Usage:    "profile configuration from .aws/config to use for assuming a role",
			},
			&cli.BoolFlag{
				Required: false,
				Value:    false,
				Name:     config.FLAG_VERBOSE,
				Usage:    "enable verbose output",
			},
			&cli.BoolFlag{
				Required: false,
				Value:    false,
				Name:     config.FLAG_HIDE_ARNS,
				Usage:    "Disable printing Role ARN & MFA Device ARN to console (even on verbose-mode)",
			},
			&cli.BoolFlag{
				Required: false,
				Value:    false,
				Name:     config.FLAG_DISABLE_DIALOG,
				Usage:    "Disable GUI-prompt and enter MFA Token Code via CLI standard input",
			},
			&cli.BoolFlag{
				Required: false,
				Value:    false,
				Name:     config.FLAG_DISABLE_REFRESH,
				Usage:    "Disable automatic session credentials mandatory refresh (600s), as defined by botocore",
			},
		},
		Action: mainAction,
	}

	err := app.Run(os.Args)
	if err != nil {
		utils.PrintError(err)
		os.Exit(1)
	}

}

func mainAction(c *cli.Context) error {

	flags := config.Flags{
		ProfileName:    c.String(config.FLAG_PROFILE_NAME),
		Verbose:        c.Bool(config.FLAG_VERBOSE),
		HideArns:       c.Bool(config.FLAG_HIDE_ARNS),
		DisableDialog:  c.Bool(config.FLAG_DISABLE_DIALOG),
		DisableRefresh: c.Bool(config.FLAG_DISABLE_REFRESH),
	}

	if flags.Verbose {
		utils.PrintBanner()
	}

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

func toPrettyJson(data interface{}) (string, error) {
	pretty, err := json.MarshalIndent(data, "", "    ")
	return string(pretty), err
}
