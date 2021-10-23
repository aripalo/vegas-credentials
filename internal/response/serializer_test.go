package response

import (
	_ "embed"
	"strings"
	"testing"
	"time"
)

//go:embed testdata/expected-output.json
var serializationOutput string

func TestSerialize(t *testing.T) {
	exp, err := time.Parse(time.RFC3339, "2020-05-19T18:06:10+00:00")
	if err != nil {
		panic(err)
	}

	r := Response{
		Version:         1,
		AccessKeyID:     "ID",
		SecretAccessKey: "SECRET",
		SessionToken:    "TOKEN",
		Expiration:      exp,
	}

	output, err := r.Serialize()
	if err != nil {
		t.Fatalf("Got %q, want nil", err)
	}

	got := string(output)
	want := strings.ReplaceAll(serializationOutput, "\t", "    ")

	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
}

func TestDeserialize(t *testing.T) {
	var r Response

	err := r.Deserialize([]byte(serializationOutput))
	if err != nil {
		t.Fatalf("Got %q, want nil", err)
	}

	exp, err := time.Parse(time.RFC3339, "2020-05-19T18:06:10Z")
	if err != nil {
		panic(err)
	}

	// The canonical way to strip a monotonic clock reading is to use t = t.Round(0).
	// https://pkg.go.dev/time#hdr-Monotonic_Clocks
	exp = exp.Round(0)

	want := Response{
		Version:         1,
		AccessKeyID:     "ID",
		SecretAccessKey: "SECRET",
		SessionToken:    "TOKEN",
		Expiration:      exp,
	}

	if r.Version != want.Version {
		t.Fatalf("Got Version %q, want %q", r.Version, want.Version)
	}

	if r.AccessKeyID != want.AccessKeyID {
		t.Fatalf("Got AccessKeyID %q, want %q", r.AccessKeyID, want.AccessKeyID)
	}

	if r.SecretAccessKey != want.SecretAccessKey {
		t.Fatalf("Got SecretAccessKey %q, want %q", r.SecretAccessKey, want.SecretAccessKey)
	}

	if r.SessionToken != want.SessionToken {
		t.Fatalf("Got SessionToken %q, want %q", r.SessionToken, want.SessionToken)
	}

	if !r.Expiration.Equal(want.Expiration) {
		t.Fatalf("Got Expiration %q, want %q", r.Expiration, want.Expiration)
	}
}
