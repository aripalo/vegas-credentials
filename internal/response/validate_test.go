package response

import (
	"testing"
	"time"

	"github.com/aripalo/vegas-credentials/internal/config"
	"github.com/aripalo/vegas-credentials/internal/newprofile"
	"github.com/aripalo/vegas-credentials/internal/vegastestapp"
)

func TestValidateCorrect(t *testing.T) {

	r := Response{
		Version:         1,
		AccessKeyID:     "ID",
		SecretAccessKey: "SECRET",
		SessionToken:    "TOKEN",
		Expiration:      time.Now().Add(time.Minute * 5),
	}

	var f config.Flags
	var p newprofile.NewProfile

	app := vegastestapp.New(f, p)

	err := r.Validate(app)
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

	var f config.Flags
	var p newprofile.NewProfile

	app := vegastestapp.New(f, p)

	err := r.Validate(app)
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

	var f config.Flags
	var p newprofile.NewProfile

	app := vegastestapp.New(f, p)

	err := r.Validate(app)
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

	var f config.Flags
	var p newprofile.NewProfile

	app := vegastestapp.New(f, p)

	err := r.Validate(app)
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

	var f config.Flags
	var p newprofile.NewProfile

	app := vegastestapp.New(f, p)

	err := r.Validate(app)
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

	var f config.Flags
	var p newprofile.NewProfile

	app := vegastestapp.New(f, p)

	err := r.Validate(app)
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

	var f config.Flags
	var p newprofile.NewProfile

	app := vegastestapp.New(f, p)

	err := r.ValidateForMandatoryRefresh(app)

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

	var f config.Flags
	var p newprofile.NewProfile

	app := vegastestapp.New(f, p)

	err := r.ValidateForMandatoryRefresh(app)
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

	f := config.Flags{
		DisableMandatoryRefresh: true,
	}
	var p newprofile.NewProfile

	app := vegastestapp.New(f, p)

	err := r.ValidateForMandatoryRefresh(app)

	if err != nil {
		t.Fatalf("Got %q, expected nil", err)
	}
}
