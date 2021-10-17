package assume

import (
	"fmt"
	"io"
	"os"

	"github.com/aripalo/aws-mfa-credential-process/internal/application/assume/awscreds/mfa"
	"github.com/aripalo/aws-mfa-credential-process/internal/config"
	"github.com/aripalo/aws-mfa-credential-process/internal/profile"
)

// App declaration
type App struct {
	WriteStream io.Writer
	Config      *config.Config
	Profile     *profile.Profile
}

// GetWriteStream implements data.Provider method
func (a *App) GetWriteStream() io.Writer {
	return a.WriteStream
}

// GetConfig implements data.Provider method
func (a *App) GetConfig() *config.Config {
	return a.Config
}

// GetProfile implements data.Provider method
func (a *App) GetProfile() *profile.Profile {
	return a.Profile
}

// New instantiates the App
func New() (*App, error) {
	a := &App{
		WriteStream: os.Stdout,
		Config:      &config.Config{},
		Profile:     &profile.Profile{},
	}
	return a, nil
}

// Assume defines the command attached to cobra
func (a *App) Assume() {

	err := a.Profile.Load(a.Config)
	if err != nil {
		panic(err)
	}

	fmt.Println("ASSUME")

	_, err = mfa.GetToken(a)
	if err != nil {
		panic(err)
	}

	//fmt.Println(fmt.Sprintf("Got token %s via %s", token.Value, token.Provider))

	/*
		creds, err := awscreds.Get(a)
		if err != nil {
			panic(err)
		}

		var foo awscreds.CredentialProcessOutput

		err = foo.GetOutput(creds)
		if err != nil {
			panic(err)
		}

		pretty, err := json.MarshalIndent(foo, "", "    ")
		if err != nil {
			panic(err)
		}

		fmt.Println(string(pretty))
	*/
}
