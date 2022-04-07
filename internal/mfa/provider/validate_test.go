package provider

import (
	"fmt"
	"testing"
)

func TestValidateCorrect(t *testing.T) {
	token := "123456"
	err := validateToken(token)
	if err != nil {
		t.Fatalf("Got %q, want nil", err)
	}
}

func TestValidateTooShort(t *testing.T) {
	token := "12345"
	err := validateToken(token)
	got := err.Error()
	want := fmt.Sprintf("Invalid OATH TOTP MFA Token Code: %q", token)
	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
}

func TestValidateWeird(t *testing.T) {
	token := "123 45"
	err := validateToken(token)
	got := err.Error()
	want := fmt.Sprintf("Invalid OATH TOTP MFA Token Code: %q", token)
	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
}

func TestValidateContainsStrings(t *testing.T) {
	token := "foobar123456"
	err := validateToken(token)
	got := err.Error()
	want := fmt.Sprintf("Invalid OATH TOTP MFA Token Code: %q", token)
	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
}

func TestValidateString(t *testing.T) {
	token := "foobar"
	err := validateToken(token)
	got := err.Error()
	want := fmt.Sprintf("Invalid OATH TOTP MFA Token Code: %q", token)
	if got != want {
		t.Fatalf("Got %q, want %q", got, want)
	}
}
