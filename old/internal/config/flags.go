package config

import (
	"github.com/urfave/cli/v2"
)

// CredentialProcessFlags is a simplified struct for accessing the provided CLI flags within the program
type CredentialProcessFlags struct {
	// ProfileName defines the profile read from ~/.aws/config
	ProfileName string

	// Verbose defines the output verbosity
	Verbose bool

	// HideArns defines if Role and MFA Serial ARN should be visible in output
	HideArns bool

	// DisableDialog defines if CLI input should be used instead of the GUI Dialog Prompt
	DisableDialog bool

	// DisableRrefresh defines if Temporary Session Credentials should be refreshed according to BotoCore mandatory refresh
	DisableRefresh bool
}

type DeleteCacheFlags struct {
	// ProfileName defines the profile read from ~/.aws/config
	ProfileName string

	// Verbose defines the output verbosity
	Verbose bool

	// DisableDialog defines if CLI input should be used instead of the GUI Dialog Prompt
	DisableDialog bool
}

// flagProfileName defines a CLI flag configuration
var flagProfileName *cli.StringFlag = &cli.StringFlag{
	Required: true,
	Name:     "profile",
	Aliases:  []string{"p"},
	Usage:    "profile configuration from .aws/config to use for assuming a role",
}

// flagProfileNameOptional defines a CLI flag configuration
var flagProfileNameOptional *cli.StringFlag = &cli.StringFlag{
	Required: false,
	Name:     "profile",
	Aliases:  []string{"p"},
	Usage:    "profile configuration to delete from cache",
}

// flagVerbose defines a CLI flag configuration
var flagVerbose *cli.BoolFlag = &cli.BoolFlag{
	Required: false,
	Value:    Config.Verbose,
	Name:     "verbose",
	Usage:    "enable verbose output",
}

// flagHideArns defines a CLI flag configuration
var flagHideArns *cli.BoolFlag = &cli.BoolFlag{
	Required: false,
	Value:    Config.HideArns,
	Name:     "hide-arns",
	Usage:    "Disable printing Role ARN & MFA Device ARN to console (even on verbose-mode)",
}

// flagDisableDialog defines a CLI flag configuration
var flagDisableDialog *cli.BoolFlag = &cli.BoolFlag{
	Required: false,
	Value:    Config.DisableDialog,
	Name:     "disable-dialog",
	Usage:    "Disable GUI-prompt and enter MFA Token Code via CLI standard input",
}

// flagDisableRefresh defines a CLI flag configuration
var flagDisableRefresh *cli.BoolFlag = &cli.BoolFlag{
	Required: false,
	Value:    Config.DisableRefresh,
	Name:     "disable-refresh",
	Usage:    "Disable automatic session credentials mandatory refresh (600s), as defined by botocore",
}

// CredentialProcessFlagsConfiguration holds the flag configuration provided for the CLI App
var CredentialProcessFlagsConfiguration = []cli.Flag{
	flagProfileName,
	flagVerbose,
	flagHideArns,
	flagDisableDialog,
	flagDisableRefresh,
}

// DeleteCacheFlagsConfiguration holds the flag configuration provided for the CLI App
var DeleteCacheFlagsConfiguration = []cli.Flag{
	flagProfileNameOptional,
	flagVerbose,
	flagDisableDialog,
}

// ParseCredentialProcessFlags reads CLI flag values provided in runtime and returns a simplified flags struct containing the parsed values
func ParseCredentialProcessFlags(c *cli.Context) CredentialProcessFlags {
	return CredentialProcessFlags{
		ProfileName:    c.String(flagProfileName.Name),
		Verbose:        c.Bool(flagVerbose.Name),
		HideArns:       c.Bool(flagHideArns.Name),
		DisableDialog:  c.Bool(flagDisableDialog.Name),
		DisableRefresh: c.Bool(flagDisableRefresh.Name),
	}
}

// ParseDeleteCacheFlags reads CLI flag values provided in runtime and returns a simplified flags struct containing the parsed values
func ParseDeleteCacheFlags(c *cli.Context) DeleteCacheFlags {
	return DeleteCacheFlags{
		ProfileName:   c.String(flagProfileNameOptional.Name),
		Verbose:       c.Bool(flagVerbose.Name),
		DisableDialog: c.Bool(flagDisableDialog.Name),
	}
}
