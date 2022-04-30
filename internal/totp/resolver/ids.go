package resolver

import "github.com/aripalo/vegas-credentials/internal/multinput"

const (
	ResolverIdYubikey   multinput.ResolverID = "Yubikey touch"
	ResolverIdCliStdin  multinput.ResolverID = "CLI stdin"
	ResolverIdGuiDialog multinput.ResolverID = "GUI prompt"
)
