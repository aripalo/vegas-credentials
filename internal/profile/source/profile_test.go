package source

import (
	"os"
	"strings"
	"testing"

	"github.com/aripalo/vegas-credentials/internal/vegastest"
)

const AWS_REGION_ENVIRONMENT_VARIABLE string = "AWS_REGION"

func TestMissingConfig(t *testing.T) {
	configPath := vegastest.GetTestdataFilePath("ENOENT")
	_, err := loadWithPath("default", configPath)
	got := err.Error()
	want := "no such file or directory"

	if !strings.Contains(got, want) {
		t.Fatalf("Got %q, expected to contain %q", got, want)
	}
}

func TestDefaultProfileMissing(t *testing.T) {
	configPath := vegastest.GetTestdataFilePath("missing-profile.ini")
	_, err := loadWithPath("default", configPath)
	got := err.Error()
	want := `section "default" does not exist`

	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
}

func TestNamedProfileMissing(t *testing.T) {
	configPath := vegastest.GetTestdataFilePath("missing-profile.ini")
	_, err := loadWithPath("celine", configPath)
	got := err.Error()
	want := `section "profile celine" does not exist`

	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
}

func TestDefaultProfileMissingMfaSerial(t *testing.T) {
	configPath := vegastest.GetTestdataFilePath("missing-mfa-serial.ini")
	_, err := loadWithPath("default", configPath)
	got := err.Error()
	want := `Missing "mfa_serial" in profile "default"`

	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
}

func TestNamedProfileMissingMfaSerial(t *testing.T) {
	configPath := vegastest.GetTestdataFilePath("missing-mfa-serial.ini")
	_, err := loadWithPath("celine", configPath)
	got := err.Error()
	want := `Missing "mfa_serial" in profile "celine"`

	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
}

func TestDefaultProfileWithOutRegion(t *testing.T) {
	configPath := vegastest.GetTestdataFilePath("without-yubikey.ini")
	result, err := loadWithPath("default", configPath)
	if err != nil {
		t.Fatalf("Got unexpected error %q", err)
	}
	got := result.Region
	want := ``

	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
}

func TestNamedProfileWithOutRegion(t *testing.T) {
	configPath := vegastest.GetTestdataFilePath("without-yubikey.ini")
	result, err := loadWithPath("celine", configPath)
	if err != nil {
		t.Fatalf("Got unexpected error %q", err)
	}
	got := result.Region
	want := ``

	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
}

func TestDefaultProfileWithRegionFromEnvironment(t *testing.T) {
	regionEnvValue := "ap-southeast-1"
	os.Setenv(AWS_REGION_ENVIRONMENT_VARIABLE, regionEnvValue)
	defer os.Unsetenv(AWS_REGION_ENVIRONMENT_VARIABLE)
	configPath := vegastest.GetTestdataFilePath("without-yubikey.ini")
	result, err := loadWithPath("default", configPath)
	if err != nil {
		t.Fatalf("Got unexpected error %q", err)
	}
	got := result.Region
	want := regionEnvValue

	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
}

func TestNamedProfileWithRegionFromEnvironment(t *testing.T) {
	regionEnvValue := "ap-southeast-1"
	os.Setenv(AWS_REGION_ENVIRONMENT_VARIABLE, regionEnvValue)
	defer os.Unsetenv(AWS_REGION_ENVIRONMENT_VARIABLE)
	configPath := vegastest.GetTestdataFilePath("without-yubikey.ini")
	result, err := loadWithPath("celine", configPath)
	if err != nil {
		t.Fatalf("Got unexpected error %q", err)
	}
	got := result.Region
	want := regionEnvValue

	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
}

func TestDefaultProfileWithRegion(t *testing.T) {
	regionEnvValue := "ap-southeast-1"
	os.Setenv(AWS_REGION_ENVIRONMENT_VARIABLE, regionEnvValue)
	defer os.Unsetenv(AWS_REGION_ENVIRONMENT_VARIABLE)
	configPath := vegastest.GetTestdataFilePath("with-region.ini")
	result, err := loadWithPath("default", configPath)
	if err != nil {
		t.Fatalf("Got unexpected error %q", err)
	}
	got := result.Region
	want := "us-west-1"

	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
}

func TestNamedProfileWithRegion(t *testing.T) {
	regionEnvValue := "ap-southeast-1"
	os.Setenv(AWS_REGION_ENVIRONMENT_VARIABLE, regionEnvValue)
	defer os.Unsetenv(AWS_REGION_ENVIRONMENT_VARIABLE)
	configPath := vegastest.GetTestdataFilePath("with-region.ini")
	result, err := loadWithPath("celine", configPath)
	if err != nil {
		t.Fatalf("Got unexpected error %q", err)
	}
	got := result.Region
	want := "ca-central-1"

	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
}

func TestDefaultProfileWithoutYubikey(t *testing.T) {
	configPath := vegastest.GetTestdataFilePath("without-yubikey.ini")
	result, err := loadWithPath("default", configPath)
	if err != nil {
		t.Fatalf("Got unexpected error %q", err)
	}
	got := result.MfaSerial
	want := "arn:aws:iam::111111111111:mfa/FrankSinatra"

	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
}

func TestNamedProfileWithoutYubikey(t *testing.T) {
	configPath := vegastest.GetTestdataFilePath("without-yubikey.ini")
	result, err := loadWithPath("celine", configPath)
	if err != nil {
		t.Fatalf("Got unexpected error %q", err)
	}
	got := result.MfaSerial
	want := "arn:aws:iam::111111111111:mfa/CelineDion"

	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
}

func TestDefaultProfileWithYubikeySerial(t *testing.T) {
	configPath := vegastest.GetTestdataFilePath("with-yubikey-serial.ini")
	result, err := loadWithPath("default", configPath)
	if err != nil {
		t.Fatalf("Got unexpected error %q", err)
	}
	gotSerial := result.YubikeySerial
	wantSerial := "12345678"

	if gotSerial != wantSerial {
		t.Fatalf("Got %q, want %q", gotSerial, wantSerial)
	}

	gotLabel := result.YubikeyLabel
	wantLabel := result.MfaSerial
	if gotLabel != wantLabel {
		t.Fatalf("Got %q, want %q", gotLabel, wantLabel)
	}
}

func TestNamedProfileWithYubikeySerial(t *testing.T) {
	configPath := vegastest.GetTestdataFilePath("with-yubikey-serial.ini")
	result, err := loadWithPath("celine", configPath)
	if err != nil {
		t.Fatalf("Got unexpected error %q", err)
	}
	gotSerial := result.YubikeySerial
	wantSerial := "87654321"

	if gotSerial != wantSerial {
		t.Fatalf("Got %q, want %q", gotSerial, wantSerial)
	}

	gotLabel := result.YubikeyLabel
	wantLabel := result.MfaSerial
	if gotLabel != wantLabel {
		t.Fatalf("Got %q, want %q", gotLabel, wantLabel)
	}
}

func TestDefaultProfileWithYubikeyLabel(t *testing.T) {
	configPath := vegastest.GetTestdataFilePath("with-yubikey-label.ini")
	result, err := loadWithPath("default", configPath)
	if err != nil {
		t.Fatalf("Got unexpected error %q", err)
	}
	gotSerial := result.YubikeySerial
	wantSerial := "12345678"

	if gotSerial != wantSerial {
		t.Fatalf("Got %q, want %q", gotSerial, wantSerial)
	}

	gotLabel := result.YubikeyLabel
	wantLabel := "aws-frank"
	if gotLabel != wantLabel {
		t.Fatalf("Got %q, want %q", gotLabel, wantLabel)
	}
}

func TestNamedProfileWithYubikeyLabel(t *testing.T) {
	configPath := vegastest.GetTestdataFilePath("with-yubikey-label.ini")
	result, err := loadWithPath("celine", configPath)
	if err != nil {
		t.Fatalf("Got unexpected error %q", err)
	}
	gotSerial := result.YubikeySerial
	wantSerial := "87654321"

	if gotSerial != wantSerial {
		t.Fatalf("Got %q, want %q", gotSerial, wantSerial)
	}

	gotLabel := result.YubikeyLabel
	wantLabel := "aws-celine"
	if gotLabel != wantLabel {
		t.Fatalf("Got %q, want %q", gotLabel, wantLabel)
	}
}
