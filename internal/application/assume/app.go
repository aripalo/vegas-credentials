package assume

import (
	"fmt"
	"io"
	"os"

	"github.com/aripalo/aws-mfa-credential-process/internal/config"
)

type App struct {
	Out    io.Writer
	Config *config.Config
}

func New() (*App, error) {
	a := &App{
		Out:    os.Stdout,
		Config: &config.Config{},
	}
	return a, nil
}

func (a *App) Assume() {

	fmt.Println("ASSUME")
}
