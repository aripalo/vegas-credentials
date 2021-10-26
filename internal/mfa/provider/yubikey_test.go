package provider

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/aripalo/vegas-credentials/internal/config"
	"github.com/aripalo/vegas-credentials/internal/newprofile"
	"github.com/aripalo/vegas-credentials/internal/newprofile/source"
	"github.com/aripalo/vegas-credentials/internal/vegastestapp"
)

const GO_WANT_HELPER_PROCESS string = "GO_WANT_HELPER_PROCESS"

func gen(helper string) func(ctx context.Context, command string, args ...string) *exec.Cmd {
	return func(ctx context.Context, command string, args ...string) *exec.Cmd {
		cs := []string{fmt.Sprintf("-test.run=%s", helper), "--", command}
		cs = append(cs, args...)
		cmd := exec.CommandContext(ctx, os.Args[0], cs...)
		cmd.Env = []string{fmt.Sprintf("%s=1", GO_WANT_HELPER_PROCESS)}
		return cmd
	}
}

// Based on https://npf.io/2015/06/testing-exec-command/

const validTestYubikeySerial = "12345678"
const validTestYubikeyLabel = "arn:aws:iam::123456789012:mfa/JaneDoe"
const validTestYubikeyToken = "123456"

func TestFoo(t *testing.T) {
	f := config.Flags{}
	p := newprofile.NewProfile{
		Source: &source.SourceProfile{
			YubikeySerial: validTestYubikeySerial,
			YubikeyLabel:  validTestYubikeyLabel,
		},
	}

	token, err := foo(f, p, gen("TestHelperProcess"))

	if err != nil {
		t.Errorf("Expected nil error, got %#v", err)
	}

	want := validTestYubikeyToken

	if token.Value != want {
		t.Errorf("Got %q, want %q", token.Value, want)
	}
}

func foo(
	f config.Flags,
	p newprofile.NewProfile,
	fakeExecCommandContext func(ctx context.Context, command string, args ...string) *exec.Cmd,
) (Token, error) {
	execCommandContext = fakeExecCommandContext
	defer func() { execCommandContext = exec.CommandContext }()

	d := vegastestapp.New(f, p)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	provider := New(d, true)

	go provider.QueryYubikey(ctx, d)

	var token Token
	var err error

	select {
	case e := <-provider.errorChan:
		err = *e
		//return token, err
	case i := <-provider.tokenChan:
		token = *i
		//return token, err
	case <-ctx.Done():
		err = ctx.Err()
		//return token, err
	}

	return token, err
}

func TestHelperProcess(t *testing.T) {
	if os.Getenv(GO_WANT_HELPER_PROCESS) != "1" {
		return
	}
	result := fmt.Sprintf("%s  %s", validTestYubikeySerial, validTestYubikeyToken)
	fmt.Fprint(os.Stdout, result)
	os.Exit(0)
}
