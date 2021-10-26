package interfaces

import (
	"io"

	"github.com/aripalo/vegas-credentials/internal/config"
	"github.com/aripalo/vegas-credentials/internal/newprofile"
)

// AssumeCredentialProcess is an interface used by multiple different internal packages
// to define common interface to move configuration and operation related data around
// to avoid import cycles.
type AssumeCredentialProcess interface {
	GetDestination() io.Writer
	GetFlags() *config.Flags
	GetProfile() *newprofile.NewProfile
}
