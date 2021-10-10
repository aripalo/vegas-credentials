package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aripalo/goawsmfa/internal/credentialprocess"
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

	profileName := c.String("profile")
	config, err := profile.GetProfile(profileName)
	fmt.Println(config)

	output, err := credentialprocess.GetOutput(config)
	fmt.Println(string(output))

	return err
}

func toPrettyJson(data interface{}) (string, error) {
	pretty, err := json.MarshalIndent(data, "", "    ")
	return string(pretty), err
}
