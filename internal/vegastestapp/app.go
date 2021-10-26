package vegastestapp

import (
	"io"

	"github.com/aripalo/vegas-credentials/internal/config"
	"github.com/aripalo/vegas-credentials/internal/profile"
)

type AssumeAppForTesting struct {
	Flags       config.Flags
	Profile     profile.Profile
	Destination io.Writer
}

func (d *AssumeAppForTesting) GetDestination() io.Writer {
	return d.Destination
}

func (d *AssumeAppForTesting) GetProfile() *profile.Profile {
	return &d.Profile
}

func (d *AssumeAppForTesting) GetFlags() *config.Flags {
	return &d.Flags
}

func New(f config.Flags, p profile.Profile) *AssumeAppForTesting {
	return &AssumeAppForTesting{
		Flags:       f,
		Profile:     p,
		Destination: io.Discard,
	}
}
