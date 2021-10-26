package profile

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/aripalo/vegas-credentials/internal/config"
)

var c = &config.Flags{
	Profile:         "my-profile",
	DurationSeconds: 3600,
}

func TestLoadProfileMinimal(t *testing.T) {

	configPath := getTestdataFilePath("valid-minimal-config")

	var profileConfig Profile

	profile, err := loadWithPath(profileConfig, c, configPath)

	if err != nil {
		t.Fatalf("Error loading config: %s", err.Error())
	}

	want := Profile{
		YubikeySerial:   "",
		YubikeyLabel:    "",
		SourceProfile:   "default",
		RoleArn:         "arn:aws:iam::123456789012:role/Demo",
		MfaSerial:       "arn:aws:iam::123456789012:mfa/JaneDoeMFA",
		DurationSeconds: 3600,
		Region:          "",
		RoleSessionName: "",
		ExternalID:      "",
	}

	matchProfiles(t, profile, want)

}

func TestLoadProfileFull(t *testing.T) {

	configPath := getTestdataFilePath("valid-full-config")

	var profileConfig Profile

	profile, err := loadWithPath(profileConfig, c, configPath)

	if err != nil {
		t.Fatalf("Error loading config: %s", err.Error())
	}

	want := Profile{
		YubikeySerial:   "123456",
		YubikeyLabel:    "foobar",
		SourceProfile:   "default",
		RoleArn:         "arn:aws:iam::123456789012:role/Demo",
		MfaSerial:       "arn:aws:iam::123456789012:mfa/JaneDoeMFA",
		DurationSeconds: 900,
		Region:          "eu-west-1",
		RoleSessionName: "my-session",
		ExternalID:      "extid123",
	}

	matchProfiles(t, profile, want)
}

func TestLoadProfileInvalid(t *testing.T) {

	configPath := getTestdataFilePath("invalid-config")

	var profileConfig Profile

	_, err := loadWithPath(profileConfig, c, configPath)

	want := fmt.Sprintf("Invalid Profile Configuration for %s", c.Profile)

	if err.Error() != want {
		t.Fatalf("Got error %q, want error %q", err.Error(), want)
	}
}

func TestLoadProfileMissing(t *testing.T) {
	configPath := getTestdataFilePath("missing-profile")

	var profileConfig Profile

	profile, err := loadWithPath(profileConfig, c, configPath)

	if err == nil {
		t.Fatalf("Expected error, but got %q", profile)
	}

	want := fmt.Sprintf("Invalid Profile Configuration for %s", c.Profile)

	if err.Error() != want {
		t.Fatalf("Error loading profile: %s", err.Error())
	}
}

func getCurrentDirectory() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return cwd
}

// TODO rename
func getTestdataFilePath(file string) string {
	cwd := getCurrentDirectory()
	return filepath.Join(cwd, "testdata", file)
}

func matchProfiles(t *testing.T, profile Profile, want Profile) {
	if profile.YubikeySerial != want.YubikeySerial {
		t.Fatalf("Got YubikeySerial %q, want %q", profile.YubikeySerial, want.YubikeySerial)
	}

	if profile.YubikeyLabel != want.YubikeyLabel {
		t.Fatalf("Got YubikeyLabel %q, want %q", profile.YubikeyLabel, want.YubikeyLabel)
	}

	if profile.SourceProfile != want.SourceProfile {
		t.Fatalf("Got SourceProfile %q, want %q", profile.SourceProfile, want.SourceProfile)
	}

	if profile.RoleArn != want.RoleArn {
		t.Fatalf("Got RoleArn %q, want %q", profile.RoleArn, want.RoleArn)
	}

	if profile.MfaSerial != want.MfaSerial {
		t.Fatalf("Got MfaSerial %q, want %q", profile.MfaSerial, want.MfaSerial)
	}

	if profile.DurationSeconds != want.DurationSeconds {
		t.Fatalf("Got DurationSeconds %q, want %q", profile.DurationSeconds, want.DurationSeconds)
	}

	if profile.Region != want.Region {
		t.Fatalf("Got Region %q, want %q", profile.Region, want.Region)
	}

	if profile.RoleSessionName != want.RoleSessionName {
		t.Fatalf("Got RoleSessionName %q, want %q", profile.RoleSessionName, want.RoleSessionName)
	}

	if profile.ExternalID != want.ExternalID {
		t.Fatalf("Got ExternalID %q, want %q", profile.ExternalID, want.ExternalID)
	}
}
