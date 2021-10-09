package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aripalo/goawsmfa/internal/mfa"
	"github.com/aripalo/goawsmfa/internal/profile"
	"github.com/urfave/cli/v2"
)

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
		},
		Action: mainAction,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func mainAction(c *cli.Context) error {

	var err error

	config, err := profile.GetProfile(c.String("profile"))
	fmt.Println(config)
	result, err := mfa.GetTokenResult(config.YubikeySerial, config.YubikeyLabel)
	fmt.Println(fmt.Sprintf("Received Token %s via %s", result.Value, result.Provider))

	return err
}
