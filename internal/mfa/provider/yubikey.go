package provider

import (
	"context"
	"errors"
	"os/exec"
	"regexp"
	"strings"

	"github.com/aripalo/aws-mfa-credential-process/internal/data"
	"github.com/aripalo/aws-mfa-credential-process/internal/profile"
)

// yubikeyTokenFindPattern describes the regexp that will match OATH TOPT MFA token code from Yubikey
var yubikeyTokenFindPattern = regexp.MustCompile(`\d{6}\d*$`)

func (t *TokenProvider) QueryYubikey(ctx context.Context, d data.Provider) {
	var token Token
	var err error

	p := d.GetProfile()
	token.Provider = TOKEN_PROVIDER_YUBIKEY_TOUCH

	label := getYubikeyLabel(p)

	cmd := exec.CommandContext(ctx, "ykman", "--device", p.YubikeySerial, "oath", "accounts", "code", label)
	stdout, err := cmd.Output()
	if err != nil {
		t.errorChan <- &err
	} else {
		value := yubikeyTokenFindPattern.FindString(strings.TrimSpace(string(stdout)))
		token.Value = value
		t.tokenChan <- &token
	}
}

// VerifyYubikey tells if Yubikey serial+label configured and given device available, which means we can query Yubikey for token
func VerifyYubikey(ctx context.Context, d data.Provider) error {
	var err error
	p := d.GetProfile()

	label := getYubikeyLabel(p)

	// first check if Yubikey configured
	if p.YubikeySerial == "" || label == "" {
		return errors.New(YubikeyErrorNotConfigured)
	}

	// if configured, check if given device is available
	cmd := exec.CommandContext(ctx, "ykman", "list")
	stdout, err := cmd.Output()
	if err != nil {
		return errors.New(YubikeyErrorFail)
	}

	available := strings.Contains(string(stdout), p.YubikeySerial)

	if !available {
		return errors.New(YubikeyErrorNotConnected)
	}
	return nil
}

func getYubikeyLabel(p *profile.Profile) string {
	if p.YubikeyLabel != "" {
		return p.YubikeyLabel
	}
	return p.MfaSerial
}

const (
	YubikeyErrorNotConfigured string = "Yubikey not configured"
	YubikeyErrorFail          string = "Yubikey failed"
	YubikeyErrorNotConnected  string = "Yubikey not connected"
)
