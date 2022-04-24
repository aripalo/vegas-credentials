package application

import "fmt"

type AssumeFlags struct {
	Profile string `mapstructure:"profile"`
}

func (a *App) Assume(flags AssumeFlags) error {

	fmt.Println("TODO!", flags)

	return nil
}
