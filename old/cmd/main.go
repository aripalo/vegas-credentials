package main

import (
	"fmt"
	"os"

	"github.com/aripalo/aws-mfa-credential-process/internal/actions"
	"github.com/aripalo/aws-mfa-credential-process/internal/config"
	"github.com/aripalo/aws-mfa-credential-process/internal/utils"
	"github.com/urfave/cli/v2"
)

// version information used as output for --version, overwritten by goreleaser with real version value
// https://goreleaser.com/customization/build/
var version = "development"

// Application Entrypoint
func main() {

	// Do some global configuration initialization
	config.Init()

	// Just print the version string, nothing else
	// https://github.com/urfave/cli/blob/master/docs/v2/manual.md#version-flag
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println(c.App.Version)
	}

	// Configure application
	app := &cli.App{
		Name:    "aws-mfa-credential-process",
		Version: version,
		Usage:   "Caching AWS Credential Process to manage assuming an IAM Role with MFA token from Yubikey and Authenticator App",
		Commands: []*cli.Command{
			{
				Name:   "assume",
				Usage:  "Perform AWS credential_process",
				Flags:  config.CredentialProcessFlagsConfiguration,
				Action: actions.Assume,
			},
			{
				Name:   "delete-cache",
				Usage:  "Deletes all temporary session credentials from cache",
				Flags:  config.DeleteCacheFlagsConfiguration,
				Action: actions.DeleteCache,
			},
		},
	}

	// Run the application
	err := app.Run(os.Args)

	// Handle errors
	if err != nil {
		utils.PrintError(err)
		os.Exit(1)
	}
}
