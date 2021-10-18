package config

import (
	"os"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestFlagProfileName(t *testing.T) {
	name := "profile"
	flag := flagProfileName
	if flag.Name != name {
		t.Fatalf(`Got %s, want %s`, flag.Name, name)
	}

	required := true
	if flag.Required != required {
		t.Fatalf(`Got %t, want %t`, flag.Required, required)
	}
}

func TestFlagVerbose(t *testing.T) {
	name := "verbose"
	flag := flagVerbose
	if flag.Name != name {
		t.Fatalf(`Got %s, want %s`, flag.Name, name)
	}

	required := false
	if flag.Required != required {
		t.Fatalf(`Got %t, want %t`, flag.Required, required)
	}
}

func TestFlagHideArns(t *testing.T) {
	name := "hide-arns"
	flag := flagHideArns
	if flag.Name != name {
		t.Fatalf(`Got %s, want %s`, flag.Name, name)
	}

	required := false
	if flag.Required != required {
		t.Fatalf(`Got %t, want %t`, flag.Required, required)
	}
}

func TestFlagDisableDialog(t *testing.T) {
	name := "disable-dialog"
	flag := flagDisableDialog
	if flag.Name != name {
		t.Fatalf(`Got %s, want %s`, flag.Name, name)
	}

	required := false
	if flag.Required != required {
		t.Fatalf(`Got %t, want %t`, flag.Required, required)
	}
}

func TestFlagDisableRefresh(t *testing.T) {
	name := "disable-refresh"
	flag := flagDisableRefresh
	if flag.Name != name {
		t.Fatalf(`Got %s, want %s`, flag.Name, name)
	}

	required := false
	if flag.Required != required {
		t.Fatalf(`Got %t, want %t`, flag.Required, required)
	}
}

type foo struct {
	name  string
	value string
}

func TestParseFlagsWithDefaults(t *testing.T) {

	want := CredentialProcessFlags{
		ProfileName:    "my-profile",
		Verbose:        false,
		HideArns:       false,
		DisableDialog:  false,
		DisableRefresh: false,
	}

	app := &cli.App{
		Name:  "testing",
		Flags: CredentialProcessFlagsConfiguration,
		Action: func(c *cli.Context) error {
			output := ParseCredentialProcessFlags(c)
			if output != want {
				t.Fatalf(`Unexpected output`)
			}
			return nil
		},
	}

	args := os.Args[0:1]
	args = append(args, "--profile", "my-profile")
	err := app.Run(args)

	if err != nil {
		t.Fatalf(`Could not run test app: %s`, err.Error())
	}
}

func TestParseFlagsWithUserProvided(t *testing.T) {

	want := CredentialProcessFlags{
		ProfileName:    "my-profile",
		Verbose:        true,
		HideArns:       true,
		DisableDialog:  true,
		DisableRefresh: true,
	}

	app := &cli.App{
		Name:  "testing",
		Flags: CredentialProcessFlagsConfiguration,
		Action: func(c *cli.Context) error {
			output := ParseCredentialProcessFlags(c)
			if output != want {
				t.Fatalf(`Unexpected output`)
			}
			return nil
		},
	}

	args := os.Args[0:1]
	args = append(args, "--profile", "my-profile", "--verbose", "--hide-arns", "--disable-dialog", "--disable-refresh")
	err := app.Run(args)

	if err != nil {
		t.Fatalf(`Could not run test app: %s`, err.Error())
	}
}
