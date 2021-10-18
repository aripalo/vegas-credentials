package provider

import (
	"context"
	"errors"
	"os/exec"
	"strings"

	"github.com/aripalo/aws-mfa-credential-process/internal/data"
)

func (t *TokenProvider) QueryYubikey(ctx context.Context, d data.Provider) {
	var token Token
	var err error

	p := d.GetProfile()
	token.Provider = TOKEN_PROVIDER_YUBIKEY_TOUCH

	cmd := exec.CommandContext(ctx, "ykman", "--device", p.YubikeySerial, "oath", "accounts", "code", p.YubikeyLabel)
	stdout, err := cmd.Output()
	if err != nil {
		t.errorChan <- &err
	} else {
		value := tokenPattern.FindString(string(stdout))
		token.Value = value
		t.tokenChan <- &token
	}
}

// VerifyYubikey tells if Yubikey serial+label configured and given device available, which means we can query Yubikey for token
func VerifyYubikey(ctx context.Context, d data.Provider) error {
	var err error
	p := d.GetProfile()

	// first check if Yubikey configured
	if p.YubikeySerial == "" || p.YubikeyLabel == "" {
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

const (
	YubikeyErrorNotConfigured string = "Yubikey not configured"
	YubikeyErrorFail          string = "Yubikey failed"
	YubikeyErrorNotConnected  string = "Yubikey not connected"
)
