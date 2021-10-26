package target

import (
	"testing"

	"github.com/aripalo/vegas-credentials/internal/vegastest"
)

func TestFrankMissingSourceProfile(t *testing.T) {
	configPath := vegastest.GetTestdataFilePath("missing-source-profile.ini")
	_, err := loadWithPath("frank@concerts", configPath)
	got := err.Error()
	want := `Profile "frank@concerts" does not contain "vegas_source_profile"`
	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
}
func TestCelineMissingSourceProfile(t *testing.T) {
	configPath := vegastest.GetTestdataFilePath("missing-source-profile.ini")
	_, err := loadWithPath("celine@concerts", configPath)
	got := err.Error()
	want := `Profile "celine@concerts" does not contain "vegas_source_profile"`
	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
}

func TestFrankMissingRoleArn(t *testing.T) {
	configPath := vegastest.GetTestdataFilePath("missing-role-arn.ini")
	_, err := loadWithPath("frank@concerts", configPath)
	got := err.Error()
	want := `Profile "frank@concerts" does not contain "vegas_role_arn"`
	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
}
func TestCelineMissingRoleArn(t *testing.T) {
	configPath := vegastest.GetTestdataFilePath("missing-role-arn.ini")
	_, err := loadWithPath("celine@concerts", configPath)
	got := err.Error()
	want := `Profile "celine@concerts" does not contain "vegas_role_arn"`
	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
}

func TestFrankInvalidRoleArn(t *testing.T) {
	configPath := vegastest.GetTestdataFilePath("invalid-role-arn.ini")
	_, err := loadWithPath("frank@concerts", configPath)
	got := err.Error()
	want := `Profile "frank@concerts" contains invalid vegas_role_arn "arn:aws:iam::00:role/Invalid". Must satisty ^arn:aws:iam:\d*:\d{12}:role\/[a-zA-Z0-9_+=,.@-]{1,64}$`
	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
}
func TestCelineInvalidRoleArn(t *testing.T) {
	configPath := vegastest.GetTestdataFilePath("invalid-role-arn.ini")
	_, err := loadWithPath("celine@concerts", configPath)
	got := err.Error()
	want := `Profile "celine@concerts" contains invalid vegas_role_arn "notvalid". Must satisty ^arn:aws:iam:\d*:\d{12}:role\/[a-zA-Z0-9_+=,.@-]{1,64}$`
	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
}

func TestFrankInvalidRoleSessionName(t *testing.T) {
	configPath := vegastest.GetTestdataFilePath("invalid-role-session-name.ini")
	_, err := loadWithPath("frank@concerts", configPath)
	got := err.Error()
	want := `Profile "frank@concerts" contains invalid role_session_name "invalid//". Must satisfy ^[a-zA-Z0-9_+=,.@-]{1,64}$`
	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
}

func TestCelineInvalidRoleSessionName(t *testing.T) {
	configPath := vegastest.GetTestdataFilePath("invalid-role-session-name.ini")
	_, err := loadWithPath("celine@concerts", configPath)
	got := err.Error()
	want := `Profile "celine@concerts" contains invalid role_session_name "invalidTooLongTooLongTooLongTooLongTooLongTooLongTooLongTooLongTooLongTooLongTooLongTooLongTooLong". Must satisfy ^[a-zA-Z0-9_+=,.@-]{1,64}$`
	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
}

func TestFrankInvalidExternalId(t *testing.T) {
	configPath := vegastest.GetTestdataFilePath("invalid-external-id.ini")
	_, err := loadWithPath("frank@concerts", configPath)
	got := err.Error()
	want := `Profile "frank@concerts" contains invalid external_id "a". Must satisfy ^[a-zA-Z0-9+=,.@:\/-]{2,}$`
	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
}

func TestCelineInvalidExternalId(t *testing.T) {
	configPath := vegastest.GetTestdataFilePath("invalid-external-id.ini")
	_, err := loadWithPath("celine@concerts", configPath)
	got := err.Error()
	want := `Profile "celine@concerts" contains invalid external_id "foo()". Must satisfy ^[a-zA-Z0-9+=,.@:\/-]{2,}$`
	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
}

func TestFrankValidMinimal(t *testing.T) {
	configPath := vegastest.GetTestdataFilePath("valid-minimal.ini")
	got, err := loadWithPath("frank@concerts", configPath)
	if err != nil {
		t.Fatalf("Got %q, want %q", err, "nil")
	}
	wantProfile := "default"
	if got.SourceProfile != wantProfile {
		t.Fatalf("Got %q, want %q", got.SourceProfile, wantProfile)
	}
	wantRoleArn := "arn:aws:iam::222222222222:role/SingerRole"
	if got.RoleArn != wantRoleArn {
		t.Fatalf("Got %q, want %q", got.RoleArn, wantRoleArn)
	}
}
func TestCelineValidMinimal(t *testing.T) {
	configPath := vegastest.GetTestdataFilePath("valid-minimal.ini")
	got, err := loadWithPath("celine@concerts", configPath)
	if err != nil {
		t.Fatalf("Got %q, want %q", err, "nil")
	}
	wantProfile := "celine"
	if got.SourceProfile != wantProfile {
		t.Fatalf("Got %q, want %q", got.SourceProfile, wantProfile)
	}
	wantRoleArn := "arn:aws:iam::222222222222:role/SingerRole"
	if got.RoleArn != wantRoleArn {
		t.Fatalf("Got %q, want %q", got.RoleArn, wantRoleArn)
	}
}

func TestFrankValidFull(t *testing.T) {
	configPath := vegastest.GetTestdataFilePath("valid-full.ini")
	got, err := loadWithPath("frank@concerts", configPath)
	if err != nil {
		t.Fatalf("Got %q, want %q", err, "nil")
	}
	wantProfile := "default"
	if got.SourceProfile != wantProfile {
		t.Fatalf("Got %q, want %q", got.SourceProfile, wantProfile)
	}
	wantRoleArn := "arn:aws:iam::222222222222:role/SingerRole"
	if got.RoleArn != wantRoleArn {
		t.Fatalf("Got %q, want %q", got.RoleArn, wantRoleArn)
	}
	wantRegion := "us-west-1"
	if got.Region != wantRegion {
		t.Fatalf("Got %q, want %q", got.Region, wantRegion)
	}
	wantDurationSeconds := 4383
	if got.DurationSeconds != wantDurationSeconds {
		t.Fatalf("Got %q, want %q", got.DurationSeconds, wantDurationSeconds)
	}
	wantRoleSessionName := "SinatraAtTheSands"
	if got.RoleSessionName != wantRoleSessionName {
		t.Fatalf("Got %q, want %q", got.RoleSessionName, wantRoleSessionName)
	}
	wantExternalId := "0093624694724"
	if got.ExternalID != wantExternalId {
		t.Fatalf("Got %q, want %q", got.ExternalID, wantExternalId)
	}

}
func TestCelineValidFull(t *testing.T) {
	configPath := vegastest.GetTestdataFilePath("valid-full.ini")
	got, err := loadWithPath("celine@concerts", configPath)
	if err != nil {
		t.Fatalf("Got %q, want %q", err, "nil")
	}
	wantProfile := "celine"
	if got.SourceProfile != wantProfile {
		t.Fatalf("Got %q, want %q", got.SourceProfile, wantProfile)
	}
	wantRoleArn := "arn:aws:iam::222222222222:role/SingerRole"
	if got.RoleArn != wantRoleArn {
		t.Fatalf("Got %q, want %q", got.RoleArn, wantRoleArn)
	}
	wantRegion := "ca-central-1"
	if got.Region != wantRegion {
		t.Fatalf("Got %q, want %q", got.Region, wantRegion)
	}
	wantDurationSeconds := 3536
	if got.DurationSeconds != wantDurationSeconds {
		t.Fatalf("Got %q, want %q", got.DurationSeconds, wantDurationSeconds)
	}
	wantRoleSessionName := "ANewDayLiveInLasVegas"
	if got.RoleSessionName != wantRoleSessionName {
		t.Fatalf("Got %q, want %q", got.RoleSessionName, wantRoleSessionName)
	}
	wantExternalId := "0886971371697"
	if got.ExternalID != wantExternalId {
		t.Fatalf("Got %q, want %q", got.ExternalID, wantExternalId)
	}
}
