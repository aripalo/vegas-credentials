package debug

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/aripalo/aws-mfa-credential-process/internal/config"
	"github.com/aripalo/aws-mfa-credential-process/internal/profile"
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

func (a *App) Debug() {

	result, err := json.MarshalIndent(a.Config, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(result))

	p, err := profile.Read(a.Config)
	if err != nil {
		panic(err)
	}

	pretty, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(pretty))
}
