package main

import (
	"encoding/json"
	"os"

	"github.com/aripalo/goawsmfa/internal/credentialprocess"
	"github.com/aripalo/goawsmfa/internal/profile"
	"github.com/aripalo/goawsmfa/internal/utils"
	"github.com/urfave/cli/v2"
)

const border string = "======================================================"
const thinBorder string = "------------------------------------------------------"

func main() {
	app := &cli.App{
		Name:  "aws-mfa-assume-credential-process",
		Usage: "Caching AWS Credential Process to manage assuming an IAM Role with MFA token from Yubikey and Authenticator App",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Required: true,
				Name:     "profile",
				Usage:    "profile configuration from .aws/config to use for assuming a role",
			},
			&cli.BoolFlag{
				Required: false,
				Value:    false,
				Name:     "verbose",
				Usage:    "enable verbose output",
			},
			&cli.BoolFlag{
				Required: false,
				Value:    false,
				Name:     "hide-arns",
				Usage:    "Set to true to disable printing Role ARN & MFA Device ARN to console",
			},
		},
		Action: mainAction,
	}

	err := app.Run(os.Args)
	if err != nil {
		//log.Fatal(err)
		utils.SafeLog(utils.TextRed(err))
		utils.SafeLog()
		utils.SafeLog(utils.TextBold(utils.TextGray(border)))
		utils.SafeLog(utils.TextBold(utils.TextRed("❌  credential_process failed")))
		utils.SafeLog(utils.TextBold(utils.TextGray(border)))
		os.Exit(1)
	}

}

func mainAction(c *cli.Context) error {

	profileName := c.String("profile")
	verboseOutput := c.Bool("verbose")
	hideArns := c.Bool("hide-arns")

	if verboseOutput {
		utils.SafeLog(utils.TextGrayDark(utils.CreateRuler("=")))
		utils.SafeLog()
		utils.SafeLog(utils.FormatMessage(utils.COLOR_TITLE, "", "credential_process", "AWS MFA Assume Credential Process"))
		utils.SafeLog()
		utils.SafeLog(utils.FormatMessage(utils.COLOR_DEBUG, "📝 ", "Author", "@aripalo"))
		utils.SafeLog(utils.FormatMessage(utils.COLOR_DEBUG, "🖥  ", "Repository", "https://github.com/aripalo/aws-mfa-assume-credential-process"))
		utils.SafeLog()
		utils.SafeLog(utils.TextGrayDark(utils.CreateRuler("-")))
	}

	var err error

	config, err := profile.GetProfile(profileName)
	if err != nil {
		utils.SafeLog(err)
		return err
	}

	output, err := credentialprocess.GetOutput(verboseOutput, profileName, hideArns, config)
	if err != nil {
		utils.SafeLog(err)
		return err
	}

	utils.OutputToAwsCredentialProcess(string(output))

	return err
}

func toPrettyJson(data interface{}) (string, error) {
	pretty, err := json.MarshalIndent(data, "", "    ")
	return string(pretty), err
}
