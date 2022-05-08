package totp

import (
	"bytes"
	_ "embed"

	"github.com/aripalo/vegas-credentials/internal/msg"
	"github.com/aripalo/vegas-credentials/internal/tmpl"
)

type inputTmplOpts struct {
	EnableGui     bool
	EnableYubikey bool
}

//go:embed data/mfa-code-message.tmpl
var inputTmpl string

func formatInputMessage(enableGui bool, enableYubikey bool) string {
	opts := inputTmplOpts{
		EnableGui:     enableGui,
		EnableYubikey: enableYubikey,
	}
	message := bytes.Buffer{}
	err := tmpl.Write(&message, "mfa-code-input", inputTmpl, opts)
	if err != nil {
		msg.Fatal(err.Error())
	}
	return string(message.Bytes())
}
