package response

import (
	"io"
	"testing"
	"time"

	"github.com/aripalo/vegas-credentials/internal/config"
	"github.com/aripalo/vegas-credentials/internal/profile"
)

func TestValidateCorrect(t *testing.T) {

	r := Response{
		Version:         1,
		AccessKeyID:     "ID",
		SecretAccessKey: "SECRET",
		SessionToken:    "TOKEN",
		Expiration:      time.Now().Add(time.Minute * 5),
	}

	var c config.Flags
	var p profile.Profile

	d := NewDpForTest(c, p)

	err := r.Validate(d)
	if err != nil {
		t.Fatalf("Got %q, expected nil", err)
	}
}

func TestValidateIncorrectVersion(t *testing.T) {

	r := Response{
		Version:         0,
		AccessKeyID:     "ID",
		SecretAccessKey: "SECRET",
		SessionToken:    "TOKEN",
		Expiration:      time.Now().Add(time.Minute * 5),
	}

	var c config.Flags
	var p profile.Profile

	d := NewDpForTest(c, p)

	err := r.Validate(d)
	got := err.Error()
	want := "Incorrect Version"

	if got != want {
		t.Fatalf("Got %q, expected %q", got, want)
	}
}

func TestValidateAccesKeyIdMissing(t *testing.T) {

	r := Response{
		Version:         1,
		SecretAccessKey: "SECRET",
		SessionToken:    "TOKEN",
		Expiration:      time.Now().Add(time.Minute * 5),
	}

	var c config.Flags
	var p profile.Profile

	d := NewDpForTest(c, p)

	err := r.Validate(d)
	got := err.Error()
	want := "Missing AccessKeyID"

	if got != want {
		t.Fatalf("Got %q, expected %q", got, want)
	}
}

func TestValidateSecretAccesKeyMissing(t *testing.T) {

	r := Response{
		Version:      1,
		AccessKeyID:  "ID",
		SessionToken: "TOKEN",
		Expiration:   time.Now().Add(time.Minute * 5),
	}

	var c config.Flags
	var p profile.Profile

	d := NewDpForTest(c, p)

	err := r.Validate(d)
	got := err.Error()
	want := "Missing SecretAccessKey"

	if got != want {
		t.Fatalf("Got %q, expected %q", got, want)
	}
}

func TestValidateSessionTokenMissing(t *testing.T) {

	r := Response{
		Version:         1,
		AccessKeyID:     "ID",
		SecretAccessKey: "SECRET",
		Expiration:      time.Now().Add(time.Minute * 5),
	}

	var c config.Flags
	var p profile.Profile

	d := NewDpForTest(c, p)

	err := r.Validate(d)
	got := err.Error()
	want := "Missing SessionToken"

	if got != want {
		t.Fatalf("Got %q, expected %q", got, want)
	}
}

func TestValidateExpired(t *testing.T) {

	r := Response{
		Version:         1,
		AccessKeyID:     "ID",
		SecretAccessKey: "SECRET",
		SessionToken:    "TOKEN",
		Expiration:      time.Now().Add(time.Minute * -5),
	}

	var c config.Flags
	var p profile.Profile

	d := NewDpForTest(c, p)

	err := r.Validate(d)
	got := err.Error()
	want := "Expired 5 minutes ago"

	if got != want {
		t.Fatalf("Got %q, expected %q", got, want)
	}
}

func TestValidateMandatoryRefreshNotRequired(t *testing.T) {

	r := Response{
		Version:         1,
		AccessKeyID:     "ID",
		SecretAccessKey: "SECRET",
		SessionToken:    "TOKEN",
		Expiration:      time.Now().Add(time.Duration(-1*9*60) * time.Second),
	}

	var c config.Flags
	var p profile.Profile

	d := NewDpForTest(c, p)

	err := r.ValidateForMandatoryRefresh(d)

	if err != nil {
		t.Fatalf("Got %q, expected nil", err)
	}
}

func TestValidateMandatoryRefreshRequired(t *testing.T) {

	r := Response{
		Version:         1,
		AccessKeyID:     "ID",
		SecretAccessKey: "SECRET",
		SessionToken:    "TOKEN",
		Expiration:      time.Now().Add(time.Duration(-1*11*60) * time.Second),
	}

	var c config.Flags
	var p profile.Profile

	d := NewDpForTest(c, p)

	err := r.ValidateForMandatoryRefresh(d)
	got := err.Error()
	want := "Mandatory refresh required because expiration in 11 minutes ago"

	if got != want {
		t.Fatalf("Got %q, expected %q", got, want)
	}
}

func TestValidateMandatoryRefreshDisabled(t *testing.T) {

	r := Response{
		Version:         1,
		AccessKeyID:     "ID",
		SecretAccessKey: "SECRET",
		SessionToken:    "TOKEN",
		Expiration:      time.Now().Add(time.Duration(-1*11*60) * time.Second),
	}

	c := config.Flags{
		DisableMandatoryRefresh: true,
	}
	var p profile.Profile

	d := NewDpForTest(c, p)

	err := r.ValidateForMandatoryRefresh(d)

	if err != nil {
		t.Fatalf("Got %q, expected nil", err)
	}
}

type DpForTest struct {
	c config.Flags
	p profile.Profile
	w io.Writer
}

func (d *DpForTest) GetWriteStream() io.Writer {
	return d.w
}

func (d *DpForTest) GetProfile() *profile.Profile {
	return &d.p
}

func (d *DpForTest) GetConfig() *config.Flags {
	return &d.c
}

func NewDpForTest(c config.Flags, p profile.Profile) *DpForTest {
	return &DpForTest{
		c: c,
		p: p,
		w: io.Discard,
	}
}
