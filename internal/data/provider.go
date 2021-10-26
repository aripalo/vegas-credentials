package data

import (
	"io"

	"github.com/aripalo/vegas-credentials/internal/config"
	"github.com/aripalo/vegas-credentials/internal/profile"
)

type Provider interface {
	GetWriteStream() io.Writer
	GetConfig() *config.Flags
	GetProfile() *profile.Profile
}
