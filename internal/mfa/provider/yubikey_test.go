package provider

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/aripalo/vegas-credentials/internal/config"
	"github.com/aripalo/vegas-credentials/internal/profile"
	"github.com/aripalo/vegas-credentials/internal/profile/source"
	"github.com/aripalo/vegas-credentials/internal/vegastestapp"
)

func TestYkmanSuccess(t *testing.T) {
	f := config.Flags{}
	p := profile.Profile{
		Source: &source.SourceProfile{
			YubikeySerial: validTestYubikeySerial,
			YubikeyLabel:  validTestYubikeyLabel,
		},
	}

	a := vegastestapp.New(f, p)

	token, err := getYubikeyTokenFromMock(a, setupMock("TestYkmanMockSuccess"))

	if err != nil {
		t.Errorf("Expected nil error, got %#v", err)
	}

	want := validTestYubikeyToken

	if token.Value != want {
		t.Errorf("Got %q, want %q", token.Value, want)
	}
}

// Ykman test helpers
// Based on https://npf.io/2015/06/testing-exec-command/
// ------------------

const GO_WANT_HELPER_PROCESS string = "GO_WANT_HELPER_PROCESS"

func setupMock(helper string) func(ctx context.Context, command string, args ...string) *exec.Cmd {
	return func(ctx context.Context, command string, args ...string) *exec.Cmd {
		cs := []string{fmt.Sprintf("-test.run=%s", helper), "--", command}
		cs = append(cs, args...)
		cmd := exec.CommandContext(ctx, os.Args[0], cs...)
		cmd.Env = []string{fmt.Sprintf("%s=1", GO_WANT_HELPER_PROCESS)}
		return cmd
	}
}

const validTestYubikeySerial = "12345678"
const validTestYubikeyLabel = "arn:aws:iam::123456789012:mfa/JaneDoe"
const validTestYubikeyToken = "123456"

func getYubikeyTokenFromMock(
	a *vegastestapp.AssumeAppForTesting,
	mock func(ctx context.Context, command string, args ...string) *exec.Cmd,
) (Token, error) {
	execCommandContext = mock
	defer func() { execCommandContext = exec.CommandContext }()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	provider := New(a, true)

	go provider.QueryYubikey(ctx, a)

	var token Token
	var err error

	select {
	case e := <-provider.errorChan:
		err = *e
	case i := <-provider.tokenChan:
		token = *i
	case <-ctx.Done():
		err = ctx.Err()
	}

	return token, err
}

// Ykman Mocks
// -----------

func TestYkmanMockSuccess(t *testing.T) {
	if os.Getenv(GO_WANT_HELPER_PROCESS) != "1" {
		return
	}
	result := fmt.Sprintf("%s  %s", validTestYubikeySerial, validTestYubikeyToken)
	fmt.Fprint(os.Stdout, result)
	os.Exit(0)
}
