package provider

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"github.com/aripalo/vegas-credentials/internal/interfaces"
)

// yubikeyTokenFindPattern describes the regexp that will match OATH TOPT MFA token code from Yubikey
var yubikeyTokenFindPattern = regexp.MustCompile(`\d{6}\d*$`)

var execCommandContext = exec.CommandContext

func (t *TokenProvider) QueryYubikey(ctx context.Context, a interfaces.AssumeCredentialProcess) {
	var token Token
	var err error

	p := a.GetProfile()
	token.Provider = TOKEN_PROVIDER_YUBIKEY_TOUCH

	cmd := execCommandContext(ctx, "ykman", "--device", p.Source.YubikeySerial, "oath", "accounts", "code", p.Source.YubikeyLabel)
	stdout, err := cmd.Output()

	if err != nil {
		t.errorChan <- &err
	} else {
		value := yubikeyTokenFindPattern.FindString(strings.TrimSpace(string(stdout)))
		if value != "" {
			token.Value = value
			t.tokenChan <- &token
		} else {
			err = fmt.Errorf("Could not get OATH account %s token for Yubikey device %s", p.Source.YubikeyLabel, p.Source.YubikeySerial)
			t.errorChan <- &err
		}

	}
}

// VerifyYubikey tells if Yubikey serial+label configured and given device available, which means we can query Yubikey for token
func VerifyYubikey(ctx context.Context, a interfaces.AssumeCredentialProcess) error {
	var err error
	p := a.GetProfile()

	// if configured, check if given device is available
	cmd := execCommandContext(ctx, "ykman", "list")
	stdout, err := cmd.Output()
	if err != nil {
		return errors.New(YubikeyErrorFail)
	}

	available := strings.Contains(string(stdout), p.Source.YubikeySerial)

	if !available {
		return errors.New(YubikeyErrorNotConnected)
	}
	return nil
}

const (
	YubikeyErrorNotConfigured string = "Yubikey not configured"
	YubikeyErrorFail          string = "Yubikey failed"
	YubikeyErrorNotConnected  string = "Yubikey not connected"
)
