package mfa

type Result struct {
	Value    string
	Provider TokenProvider
}

type TokenProvider string

const (
	TOKEN_PROVIDER_YUBIKEY TokenProvider = "Yubikey"
	TOKEN_PROVIDER_CLI     TokenProvider = "CLI"
	TOKEN_PROVIDER_DIALOG  TokenProvider = "Dialog Prompt"
)
