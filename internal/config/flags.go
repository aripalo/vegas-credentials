package config

import "github.com/urfave/cli/v2"

// Flags is a simplified struct for accessing the provided CLI flags within the program
type Flags struct {
	// ProfileName defines the profile read from ~/.aws/config
	ProfileName string

	// Verbose defines the output verbosity
	Verbose bool

	// HideArns defines if Role and MFA Serial ARN should be visible in output
	HideArns bool

	// DisableDialog defines if CLI MFA input should be used instead of the GUI Dialog Prompt
	DisableDialog bool

	// DisableRrefresh defines if Temporary Session Credentials should be refreshed according to BotoCore mandatory refresh
	DisableRefresh bool
}

// flagProfileName defines a CLI flag configuration
var flagProfileName *cli.StringFlag = &cli.StringFlag{
	Required: true,
	Name:     "profile",
	Usage:    "profile configuration from .aws/config to use for assuming a role",
}

// flagVerbose defines a CLI flag configuration
var flagVerbose *cli.BoolFlag = &cli.BoolFlag{
	Required: false,
	Value:    false,
	Name:     "verbose",
	Usage:    "enable verbose output",
}

// flagHideArns defines a CLI flag configuration
var flagHideArns *cli.BoolFlag = &cli.BoolFlag{
	Required: false,
	Value:    false,
	Name:     "hide-arns",
	Usage:    "Disable printing Role ARN & MFA Device ARN to console (even on verbose-mode)",
}

// flagDisableDialog defines a CLI flag configuration
var flagDisableDialog *cli.BoolFlag = &cli.BoolFlag{
	Required: false,
	Value:    false,
	Name:     "disable-dialog",
	Usage:    "Disable GUI-prompt and enter MFA Token Code via CLI standard input",
}

// flagDisableRefresh defines a CLI flag configuration
var flagDisableRefresh *cli.BoolFlag = &cli.BoolFlag{
	Required: false,
	Value:    false,
	Name:     "disable-refresh",
	Usage:    "Disable automatic session credentials mandatory refresh (600s), as defined by botocore",
}

// FlagsConfiguration holds the flag configuration provided for the CLI App
var FlagsConfiguration = []cli.Flag{
	flagProfileName,
	flagVerbose,
	flagHideArns,
	flagDisableDialog,
	flagDisableRefresh,
}

// ParseFlags reads CLI flag values provided in runtime and returns a simplified Flags struct containing the parsed values
func ParseFlags(c *cli.Context) Flags {
	flags := Flags{
		ProfileName:    c.String(flagProfileName.Name),
		Verbose:        c.Bool(flagVerbose.Name),
		HideArns:       c.Bool(flagHideArns.Name),
		DisableDialog:  c.Bool(flagDisableDialog.Name),
		DisableRefresh: c.Bool(flagDisableRefresh.Name),
	}
	return flags
}
