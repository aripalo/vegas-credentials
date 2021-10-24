package provider

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/aripalo/aws-mfa-credential-process/internal/config"
	"github.com/aripalo/aws-mfa-credential-process/internal/profile"
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
	c := config.Config{}
	p := profile.Profile{
		YubikeySerial: validTestYubikeySerial,
		YubikeyLabel:  validTestYubikeyLabel,
	}

	token, err := foo(c, p, gen("TestHelperProcess"))

	if err != nil {
		t.Errorf("Expected nil error, got %#v", err)
	}

	want := validTestYubikeyToken

	if token.Value != want {
		t.Errorf("Got %q, want %q", token.Value, want)
	}
}

func foo(
	c config.Config,
	p profile.Profile,
	fakeExecCommandContext func(ctx context.Context, command string, args ...string) *exec.Cmd,
) (Token, error) {
	execCommandContext = fakeExecCommandContext
	defer func() { execCommandContext = exec.CommandContext }()

	d := NewDpForTest(c, p)

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
	fmt.Fprintf(os.Stdout, result)
	os.Exit(0)
}

type DpForTest struct {
	c config.Config
	p profile.Profile
	w io.Writer
}

func (d *DpForTest) GetWriteStream() io.Writer {
	return d.w
}

func (d *DpForTest) GetProfile() *profile.Profile {
	return &d.p
}

func (d *DpForTest) GetConfig() *config.Config {
	return &d.c
}

func NewDpForTest(c config.Config, p profile.Profile) *DpForTest {
	return &DpForTest{
		c: c,
		p: p,
		w: io.Discard,
	}
}
