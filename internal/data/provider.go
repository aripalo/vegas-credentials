package data

import (
	"io"

	"github.com/aripalo/aws-mfa-credential-process/internal/config"
	"github.com/aripalo/aws-mfa-credential-process/internal/profile"
)

type Provider interface {
	GetWriteStream() io.Writer
	GetConfig() *config.Config
	GetProfile() *profile.Profile
}
