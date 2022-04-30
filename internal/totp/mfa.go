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
	//yubikeyLabel string
	yubikeyEnabled bool
	resolvers      []multinput.InputResolver
}

// Initialize a new TOTP construct which provides the Get method.
func New(options TotpOptions) TOTP {

	var resolvers []multinput.InputResolver

	resolvers = append(resolvers, resolver.CLI)

	if options.EnableGui {
		resolvers = append(resolvers, resolver.GUI)
	}

	y, err := yubikey.New(yubikey.Options{
		Device:    options.YubikeySerial,
		Account:   options.YubikeyLabel,
		EnableGui: options.EnableGui,
	})

	yubikeyEnabled := false

	if err == nil {
		yubikeyEnabled = true
		resolvers = append(resolvers, resolver.Yubikey(y))
	}

	return &Totp{
		//yubikeyLabel: options.YubikeyLabel,
		yubikeyEnabled: yubikeyEnabled,
		resolvers:      resolvers,
	}
}

const MfaTimeout time.Duration = 60 * time.Second

// Method responsible for actually querying the TOTP code from end-user.
func (m *Totp) Get() (string, error) {

	// Print some end-user messages
	if m.yubikeyEnabled {
		msg.Prompt("🔑", "Input the Token Code or touch Yubikey:")
	} else {
		msg.Prompt("🔑", "Input the Token Code:")
	}

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
