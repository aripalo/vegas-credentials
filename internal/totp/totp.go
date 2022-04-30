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

type TotpOptions struct {
	YubikeySerial string
	YubikeyLabel  string
	EnableGui     bool
}

// Returned interface.
type TOTP interface {
	Get() (string, error)
}

// Describes the internal configuration/state.
type Totp struct {
	guiEnabled     bool
	yubikeyEnabled bool
	resolvers      []multinput.InputResolver
}

// Initialize a new TOTP construct which provides the Get method.
func New(options TotpOptions) TOTP {
	guiEnabled := options.EnableGui
	yubikeyEnabled := false

	var resolvers []multinput.InputResolver

	resolvers = append(resolvers, resolver.CLI)

	if guiEnabled {
		resolvers = append(resolvers, resolver.GUI)
	}

	y, err := yubikey.New(yubikey.Options{
		Device:    options.YubikeySerial,
		Account:   options.YubikeyLabel,
		EnableGui: guiEnabled,
	})

	if err == nil {
		yubikeyEnabled = true
		resolvers = append(resolvers, resolver.Yubikey(y))
	}

	return &Totp{
		guiEnabled:     guiEnabled,
		yubikeyEnabled: yubikeyEnabled,
		resolvers:      resolvers,
	}
}

const MfaTimeout time.Duration = 60 * time.Second

//go:embed data/mfa-code-message.tmpl
var inputTmpl string

type inputTmplOpts struct {
	GuiEnabled     bool
	YubikeyEnabled bool
}

// Method responsible for actually querying the TOTP code from end-user.
func (m *Totp) Get() (string, error) {

	message := formatInputMessage(m.guiEnabled, m.yubikeyEnabled)
	msg.Prompt("🔑", message)

	ctx, cancel := context.WithTimeout(context.Background(), MfaTimeout)
	defer cancel()

	// Use Multinput to query the TOTP from various resolvers (CLI, GUI, Yubikey).
	mi := multinput.New(m.resolvers)
	result, err := mi.Provide(ctx)
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
