package totp

import (
	"context"
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/aripalo/vegas-credentials/internal/msg"
	"github.com/aripalo/vegas-credentials/internal/multinput"
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
	yubikeyLabel string
	resolvers    []multinput.InputResolver
}

// Initialize a new TOTP construct which provides the Get method.
func New(options TotpOptions) TOTP {
	return &Totp{
		yubikeyLabel: options.YubikeyLabel,
		resolvers: []multinput.InputResolver{
			resolver.ConfigureCLI(),
			resolver.ConfigureGUI(options.EnableGui),
			resolver.ConfigureYubikey(yubikey.Options{
				Device:    options.YubikeySerial,
				Account:   options.YubikeyLabel,
				EnableGui: options.EnableGui,
			}),
		},
	}
}

const MfaTimeout time.Duration = 60 * time.Second

// Method responsible for actually querying the TOTP code from end-user.
func (m *Totp) Get() (string, error) {

	// Print some end-user messages
	// TODO detect GUI/Yubikey... requires some refactor
	msg.Prompt("🔑", "Input the Token Code (or touch Yubikey if configured):")

	ctx, cancel := context.WithTimeout(context.Background(), MfaTimeout)
	defer cancel()

	// Use Multinput to query the TOTP from various resolvers (CLI, GUI, Yubikey).
	mi := multinput.New(m.resolvers)
	result, err := mi.Provide(ctx)
	if err != nil {
		return "", err
	}

	// Trim just in case.
	code := strings.TrimSpace(result.Value)

	// Ensure token looks like it should.
	// Final validation done by AWS STS.
	if !isValidToken(code) {
		return code, errors.New("invalid mfa code") // TODO
	}

	return code, nil
}

// Validates the received value looks like a TOTP MFA Token Code.
func isValidToken(value string) bool {
	return regexp.MustCompile(`^\d{6}\d*$`).MatchString(value)
}
