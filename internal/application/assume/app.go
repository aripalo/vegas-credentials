package assume

import (
	"fmt"
	"io"
	"os"

	"github.com/aripalo/aws-mfa-credential-process/internal/config"
	"github.com/aripalo/aws-mfa-credential-process/internal/profile"
)

type App struct {
	Out     io.Writer
	Config  *config.Config
	Profile *profile.Profile
}

func New() (*App, error) {
	a := &App{
		Out:     os.Stdout,
		Config:  &config.Config{},
		Profile: &profile.Profile{},
	}
	return a, nil
}

func (a *App) Assume() {

	err := a.Profile.Load(a.Config)
	if err != nil {
		panic(err)
	}

	fmt.Println("ASSUME")
}
