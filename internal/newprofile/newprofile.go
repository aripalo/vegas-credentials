package newprofile

import (
	"github.com/aripalo/vegas-credentials/internal/newprofile/source"
	"github.com/aripalo/vegas-credentials/internal/newprofile/target"
)

type NewProfile struct {
	Source *source.SourceProfile
	Target *target.TargetProfile
}

func New(targetName string) (*NewProfile, error) {
	n := new(NewProfile)
	var err error

	t, err := target.New(targetName)
	if err != nil {
		return n, err
	}

	n.Target = t

	s, err := source.New(n.Target.SourceProfile)
	if err != nil {
		return n, err
	}

	n.Source = s

	// Set region from source if not given for target
	if n.Target.Region == "" {
		n.Target.Region = n.Source.Region
	}

	return n, err

}
