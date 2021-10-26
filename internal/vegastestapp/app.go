package vegastestapp

import (
	"io"

	"github.com/aripalo/vegas-credentials/internal/config"
	"github.com/aripalo/vegas-credentials/internal/newprofile"
)

type AssumeAppForTesting struct {
	Flags       config.Flags
	Profile     newprofile.NewProfile
	Destination io.Writer
}

func (d *AssumeAppForTesting) GetDestination() io.Writer {
	return d.Destination
}

func (d *AssumeAppForTesting) GetProfile() *newprofile.NewProfile {
	return &d.Profile
}

func (d *AssumeAppForTesting) GetFlags() *config.Flags {
	return &d.Flags
}

func New(f config.Flags, p newprofile.NewProfile) *AssumeAppForTesting {
	return &AssumeAppForTesting{
		Flags:       f,
		Profile:     p,
		Destination: io.Discard,
	}
}
