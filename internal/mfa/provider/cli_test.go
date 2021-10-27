package provider

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/aripalo/vegas-credentials/internal/config"
	"github.com/aripalo/vegas-credentials/internal/profile"
	"github.com/aripalo/vegas-credentials/internal/prompt"
	"github.com/aripalo/vegas-credentials/internal/vegastestapp"
)

func TestCliSuccess(t *testing.T) {

	want := Token{
		Value:    "123456",
		Provider: TOKEN_PROVIDER_CLI_INPUT,
	}

	cliPrompt = func(ctx context.Context, text string) (string, error) {
		return want.Value, nil
	}

	defer func() { cliPrompt = prompt.Cli }()

	f := config.Flags{}
	p := profile.Profile{}

	a := vegastestapp.New(f, p)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	provider := New(a, true)
	go provider.QueryCLI(ctx, a)

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

	if err != nil {
		t.Fatalf(`Got %q, want nil`, err)
	}

	if token.Value != want.Value {
		t.Fatalf(`Got %q want %q`, token.Value, want.Value)
	}

	if string(token.Provider) != string(want.Provider) {
		t.Fatalf(`Got %q want %q`, token.Provider, want.Provider)
	}
}

func TestCliError(t *testing.T) {

	wantErr := "Some error"

	want := Token{
		Value:    "",
		Provider: "", // TODO should it return the provider still? Maybe hard to implement!
	}

	cliPrompt = func(ctx context.Context, text string) (string, error) {
		return want.Value, errors.New(wantErr)
	}

	defer func() { cliPrompt = prompt.Cli }()

	f := config.Flags{}
	p := profile.Profile{}

	a := vegastestapp.New(f, p)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	provider := New(a, true)
	go provider.QueryCLI(ctx, a)

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

	if err.Error() != wantErr {
		t.Fatalf(`Got %q, want %q`, err, wantErr)
	}

	if token.Value != want.Value {
		t.Fatalf(`Got %q want %q`, token.Value, want.Value)
	}

	if string(token.Provider) != string(want.Provider) {
		t.Fatalf(`Got %q want %q`, token.Provider, want.Provider)
	}
}
