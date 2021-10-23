package response

import (
	"bytes"
	_ "embed"
	"strings"
	"testing"
	"time"
)

//go:embed testdata/expected-output.json
var expectedOutput string

func TestOutput(t *testing.T) {

	exp, err := time.Parse(time.RFC3339, "2020-05-19T18:06:10+00:00")
	if err != nil {
		panic(err)
	}

	var output bytes.Buffer

	r := Response{
		destination:     &output,
		Version:         1,
		AccessKeyID:     "ID",
		SecretAccessKey: "SECRET",
		SessionToken:    "TOKEN",
		Expiration:      exp,
	}

	err = r.Output()
	if err != nil {
		t.Fatalf("Got %q, want nil", err)
	}

	got := output.String()
	want := strings.ReplaceAll(expectedOutput, "\t", "    ")

	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}

}
