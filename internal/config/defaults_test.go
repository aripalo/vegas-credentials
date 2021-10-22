package config

import (
	"os"
	"testing"
)

func resetEnv() {
	os.Unsetenv("NO_COLOR")
	os.Unsetenv("AWS_MFA_CREDENTIAL_PROCESS_NO_COLOR")
	os.Unsetenv("TERM")
}

func TestDefaultProfile(t *testing.T) {
	want := ""
	if Defaults.Profile.Value != want {
		t.Fatalf("Got %q, want %q", Defaults.Profile.Value, want)
	}
}

func TestDefaultDurationSeconds(t *testing.T) {
	want := 3600
	if Defaults.DurationSeconds.Value != want {
		t.Fatalf("Got %q, want %q", Defaults.DurationSeconds.Value, want)
	}
}

func TestDefaultVerbose(t *testing.T) {
	want := false
	if Defaults.Verbose.Value != want {
		t.Fatalf("Got %t, want %t", Defaults.Verbose.Value, want)
	}
}

func TestDefaultHideArns(t *testing.T) {
	want := false
	if Defaults.HideArns.Value != want {
		t.Fatalf("Got %t, want %t", Defaults.HideArns.Value, want)
	}
}

func TestDefaultDisableDialog(t *testing.T) {
	want := false
	if Defaults.DisableDialog.Value != want {
		t.Fatalf("Got %t, want %t", Defaults.DisableDialog.Value, want)
	}
}

func TestDefaultDisableRefresh(t *testing.T) {
	want := false
	if Defaults.DisableRefresh.Value != want {
		t.Fatalf("Got %t, want %t", Defaults.DisableRefresh.Value, want)
	}
}

func TestDefaultNoColorWithoutEnv(t *testing.T) {

	defer resetEnv()

	os.Unsetenv("NO_COLOR")
	os.Unsetenv("AWS_MFA_CREDENTIAL_PROCESS_NO_COLOR")
	os.Setenv("TERM", "xterm-256color")

	got := resolveNoColorDefaultValue()
	want := false
	if got != want {
		t.Fatalf("Got %t, want %t", got, want)
	}
}

func TestDefaultNoColorWithNoColorEnv(t *testing.T) {

	defer resetEnv()

	os.Setenv("NO_COLOR", "true")
	os.Unsetenv("AWS_MFA_CREDENTIAL_PROCESS_NO_COLOR")
	os.Setenv("TERM", "xterm-256color")

	got := resolveNoColorDefaultValue()
	want := true
	if got != want {
		t.Fatalf("Got %t, want %t", got, want)
	}
}

func TestDefaultNoColorWithAppNoColorEnv(t *testing.T) {

	defer resetEnv()

	os.Unsetenv("NO_COLOR")
	os.Setenv("AWS_MFA_CREDENTIAL_PROCESS_NO_COLOR", "true")
	os.Setenv("TERM", "xterm-256color")

	got := resolveNoColorDefaultValue()
	want := true
	if got != want {
		t.Fatalf("Got %t, want %t", got, want)
	}
}

func TestDefaultNoColorWithDumbTerm(t *testing.T) {

	defer resetEnv()

	os.Unsetenv("NO_COLOR")
	os.Unsetenv("AWS_MFA_CREDENTIAL_PROCESS_NO_COLOR")
	os.Setenv("TERM", "dumb")

	got := resolveNoColorDefaultValue()
	want := true
	if got != want {
		t.Fatalf("Got %t, want %t", got, want)
	}
}
