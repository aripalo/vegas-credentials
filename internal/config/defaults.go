package config

import (
	"os"
)

type StringDefaultOption struct {
	Name  string
	Value string
	Usage string
}

type IntDefaultOption struct {
	Name  string
	Value int
	Usage string
}

type BoolDefaultOption struct {
	Name  string
	Value bool
	Usage string
}

type DefaultOptions struct {
	Profile                 StringDefaultOption
	DurationSeconds         IntDefaultOption
	Debug                   BoolDefaultOption
	Verbose                 BoolDefaultOption
	HideArns                BoolDefaultOption
	DisableDialog           BoolDefaultOption
	DisableMandatoryRefresh BoolDefaultOption
	NoColor                 BoolDefaultOption
}

var Defaults = DefaultOptions{
	Profile:                 StringDefaultOption{"profile", "", "Which AWS Profile to use from ~/.aws/config"},
	DurationSeconds:         IntDefaultOption{"duration-seconds", 3600, "Default Session duration in seconds"},
	Debug:                   BoolDefaultOption{"debug", false, "Prints out various debugging information"},
	Verbose:                 BoolDefaultOption{"verbose", false, "Verbose output"},
	HideArns:                BoolDefaultOption{"hide-arns", false, "Hide IAM Role & MFA Serial ARNS from output (even on verbose mode)"},
	DisableDialog:           BoolDefaultOption{"disable-dialog", false, "Disable GUI Dialog Prompt and use CLI stdin input instead"},
	DisableMandatoryRefresh: BoolDefaultOption{"disable-mandatory-refresh", false, "Disable Session Credentials refreshing (as defined in Botocore)"},
	NoColor:                 BoolDefaultOption{"no-color", resolveNoColorDefaultValue(), "Disable fancy colored output with emojis"},
}

func resolveNoColorDefaultValue() bool {
	// Check if NO_COLOR set https://no-color.org/
	_, noColorSet := os.LookupEnv("NO_COLOR")
	if noColorSet {
		return true
	}

	// Check if app-specific _NO_COLOR set https://medium.com/@jdxcode/12-factor-cli-apps-dd3c227a0e46
	_, appNoColorSet := os.LookupEnv("AWS_MFA_CREDENTIAL_PROCESS_NO_COLOR")
	if appNoColorSet {
		return true
	}

	// Check if $TERM=dumb https://unix.stackexchange.com/a/43951
	if os.Getenv("TERM") == "dumb" {
		return true
	}

	// Otherwise default NoColor=false (i.e. colors enabled)
	return false
}
