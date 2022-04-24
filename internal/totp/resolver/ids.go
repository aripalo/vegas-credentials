package resolver

import "github.com/aripalo/vegas-credentials/internal/multinput"

const (
	ResolverIdYubikey   multinput.ResolverID = "Yubikey Touch"
	ResolverIdCliStdin  multinput.ResolverID = "CLI Standard Input"
	ResolverIdGuiDialog multinput.ResolverID = "GUI Dialog Prompt"
)
