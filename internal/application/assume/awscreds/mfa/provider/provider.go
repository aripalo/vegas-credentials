package provider

import (
	"context"
	"regexp"
	"time"

	"github.com/aripalo/aws-mfa-credential-process/internal/data"
)

// Token contains the OATH TOPT MFA token value and information about which Porivder Type gave the result
type Token struct {
	// OATH TOPT MFA Token Code received from MFA Provider
	Value string
	// OATH TOPT MFA Provider Type that provided the Token Code Value
	Provider Type
}

// Type defines which MFA OATH TOPT Provider used
type Type string

const (
	// Yubikey Touch OATH TOTP Hardware Token.
	TOKEN_PROVIDER_YUBIKEY_TOUCH Type = "Yubikey Touch"

	// User provided OATH TOPT Token via CLI stdin: Copy-paste or manual input from Authenticator App.
	TOKEN_PROVIDER_CLI_INPUT Type = "CLI input"

	// User provided OATH TOPT Token via GUI Dialog Prompt stdin: Copy-paste or manual input from Authenticator App.
	TOKEN_PROVIDER_GUI_DIALOG_PROMPT Type = "GUI Dialog Prompt"
)

// TokenProvider defines the struct upon which different providers define their methods
type TokenProvider struct {
	tokenChan chan *Token
	errorChan chan *error
}

// MFA_TIMEOUT_SECONDS configures global timeout for the Provide method
const MFA_TIMEOUT_SECONDS = 60

// Provide OATH TOPT MFA Token from supported providers
func (t *TokenProvider) Provide(d data.Provider) (Token, error) {
	t.tokenChan = make(chan *Token, 1)
	t.errorChan = make(chan *error, 1)

	var token Token
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), MFA_TIMEOUT_SECONDS*time.Second)
	defer cancel()

	if UseGui(d) {
		go t.QueryGUI(ctx, d)
	} else {
		go t.QueryCLI(ctx, d)
	}

	if HasYubikey(d) {
		go t.QueryYubikey(ctx, d)
	}

	select {
	case i := <-t.tokenChan:
		token = *i
		return token, err
	case <-ctx.Done():
		return token, ctx.Err()
	}
}

// UseGui tells if GUI Dialog Prompt should be used or not (and fallback to CLI stdin input)
func UseGui(d data.Provider) bool {
	c := d.GetConfig()
	return !c.DisableDialog
}

// HasYubikey telss if Yubikey serial+label configured, which means we can query Yubikey for token
func HasYubikey(d data.Provider) bool {
	p := d.GetProfile()
	return p.YubikeySerial != "" && p.YubikeyLabel != ""
}

// tokenPattern describes the regexp that will match OATH TOPT MFA token code
var tokenPattern = regexp.MustCompile("\\d{6}\\d*")
