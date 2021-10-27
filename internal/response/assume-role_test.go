package response

import (
	"errors"
	"testing"
	"time"

	"github.com/aripalo/vegas-credentials/internal/config"
	"github.com/aripalo/vegas-credentials/internal/interfaces"
	"github.com/aripalo/vegas-credentials/internal/profile"
	"github.com/aripalo/vegas-credentials/internal/sts"
	"github.com/aripalo/vegas-credentials/internal/vegastestapp"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

func TestAssumeRoleSuccess(t *testing.T) {

	value := credentials.Value{
		AccessKeyID:     "AKIAIOSFODNN7EXAMPLE",
		SecretAccessKey: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
		SessionToken:    "EXAMPLE/Ahb5owu2Eigaiv5ceilaeM2ohGhaiCh4",
	}

	expiration := time.Now().Add(-900 * time.Second)

	getAssumedCredentials = func(a interfaces.AssumeCredentialProcess) (credentials.Value, time.Time, error) {

		return value, expiration, nil
	}

	defer func() { getAssumedCredentials = sts.GetAssumedCredentials }()

	r := new(Response)

	f := config.Flags{}
	p := profile.Profile{}
	a := vegastestapp.New(f, p)

	err := r.AssumeRole(a)

	if err != nil {
		t.Fatalf(`Got %q, want nil`, err)
	}

	if r.AccessKeyID != value.AccessKeyID {
		t.Fatalf(`Got %q, want %q`, r.AccessKeyID, value.AccessKeyID)
	}

	if r.SecretAccessKey != value.SecretAccessKey {
		t.Fatalf(`Got %q, want %q`, r.SecretAccessKey, value.SecretAccessKey)
	}

	if r.SessionToken != value.SessionToken {
		t.Fatalf(`Got %q, want %q`, r.SessionToken, value.SessionToken)
	}
}

func TestAssumeRoleFailure(t *testing.T) {

	var value credentials.Value
	wantErr := "AccessDeniedException"

	expiration := time.Now().Add(-900 * time.Second)

	getAssumedCredentials = func(a interfaces.AssumeCredentialProcess) (credentials.Value, time.Time, error) {

		return value, expiration, errors.New(wantErr)
	}

	defer func() { getAssumedCredentials = sts.GetAssumedCredentials }()

	r := new(Response)

	f := config.Flags{}
	p := profile.Profile{}
	a := vegastestapp.New(f, p)

	err := r.AssumeRole(a)

	if err.Error() != wantErr {
		t.Fatalf(`Got %q, want %q`, err, wantErr)
	}

	if r.AccessKeyID != "" {
		t.Fatalf(`Got %q, want ""`, r.AccessKeyID)
	}

	if r.SecretAccessKey != "" {
		t.Fatalf(`Got %q, want ""`, r.SecretAccessKey)
	}

	if r.SessionToken != "" {
		t.Fatalf(`Got %q, want ""`, r.SessionToken)
	}
}
