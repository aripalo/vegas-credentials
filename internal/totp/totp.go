package totp

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aripalo/vegas-credentials/internal/msg"
	"github.com/aripalo/vegas-credentials/internal/multinput"
	"github.com/aripalo/vegas-credentials/internal/totp/resolver"
	"github.com/aripalo/vegas-credentials/internal/yubikey"
)

// Options given to GetCode function that controls which inputs are queried
// for OATH TOTP code. CLI input is always queried.
type Options struct {
	YubikeySerial string
	YubikeyLabel  string
	EnableGui     bool
	EnableYubikey bool
}

// Timeout how long to wait for OATH TOTP code.
const timeout time.Duration = 60 * time.Second

// Get OATH TOTP MFA code from various input sources.
func GetCode(ctx context.Context, options Options) (string, error) {

	ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var resolvers []multinput.InputResolver
	resolvers = append(resolvers, resolver.CLI)

	if options.EnableGui {
		resolvers = append(resolvers, resolver.GUI)
	}

	if options.EnableYubikey {
		if y, err := yubikey.New(yubikey.Options{
			Device:    options.YubikeySerial,
			Account:   options.YubikeyLabel,
			EnableGui: options.EnableGui,
		}); err == nil {
			resolvers = append(resolvers, resolver.Yubikey(y))
		}
	}

	message := formatInputMessage(options.EnableGui, options.EnableYubikey)
	msg.Prompt("🔑", message)

	mi := multinput.New(resolvers)
	result, err := mi.Provide(ctxWithTimeout)
	if err != nil {
		return "", err
	}

	msg.Debug("ℹ️", fmt.Sprintf("MFA: Token received via %s", result.ResolverID))

	code := result.Value

	if !isValidToken(code) {
		return code, errors.New("invalid mfa code") // TODO
	}

	return code, nil
}
