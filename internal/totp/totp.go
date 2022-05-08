package totp

import (
	"bytes"
	"context"
	_ "embed"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/aripalo/vegas-credentials/internal/msg"
	"github.com/aripalo/vegas-credentials/internal/multinput"
	"github.com/aripalo/vegas-credentials/internal/tmpl"
	"github.com/aripalo/vegas-credentials/internal/totp/resolver"
	"github.com/aripalo/vegas-credentials/internal/yubikey"
)

type GetCodeInput struct {
	YubikeySerial string
	YubikeyLabel  string
	EnableGui     bool
	EnableYubikey bool
}

const MfaTimeout time.Duration = 60 * time.Second

func GetCode(ctx context.Context, input GetCodeInput) (string, error) {

	ctxWithTimeout, cancel := context.WithTimeout(ctx, MfaTimeout)
	defer cancel()

	var resolvers []multinput.InputResolver

	resolvers = append(resolvers, resolver.CLI)

	if input.EnableGui {
		resolvers = append(resolvers, resolver.GUI)
	}

	if input.EnableYubikey {
		y, err := yubikey.New(yubikey.Options{
			Device:    input.YubikeySerial,
			Account:   input.YubikeyLabel,
			EnableGui: input.EnableGui,
		})

		if err == nil {
			resolvers = append(resolvers, resolver.Yubikey(y))
		}
	}

	message := formatInputMessage(input.EnableGui, input.EnableYubikey)
	msg.Prompt("🔑", message)

	// Use Multinput to query the TOTP from various resolvers (CLI, GUI, Yubikey).
	mi := multinput.New(resolvers)
	result, err := mi.Provide(ctxWithTimeout)
	if err != nil {
		return "", err
	}

	msg.Debug("ℹ️", fmt.Sprintf("MFA: Token received via %s", result.ResolverID))

	// Trim just in case.
	code := strings.TrimSpace(result.Value)

	// Ensure token looks like it should.
	// Final validation done by AWS STS.
	if !isValidToken(code) {
		return code, errors.New("invalid mfa code") // TODO
	}

	return code, nil
}

// Validates the received value looks like a TOTP MFA Token Code:
// 6 digits or more.
func isValidToken(value string) bool {
	return regexp.MustCompile(`^\d{6}\d*$`).MatchString(value)
}

type inputTmplOpts struct {
	GuiEnabled     bool
	YubikeyEnabled bool
}

//go:embed data/mfa-code-message.tmpl
var inputTmpl string

func formatInputMessage(guiEnabled bool, yubikeyEnabled bool) string {
	opts := inputTmplOpts{
		GuiEnabled:     guiEnabled,
		YubikeyEnabled: yubikeyEnabled,
	}
	message := bytes.Buffer{}
	err := tmpl.Write(&message, "mfa-code-input", inputTmpl, opts)
	if err != nil {
		msg.Fatal(err.Error())
	}
	return string(message.Bytes())
}
