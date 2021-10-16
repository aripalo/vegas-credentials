package actions

import (
	"github.com/aripalo/aws-mfa-credential-process/internal/config"
	"github.com/aripalo/aws-mfa-credential-process/internal/securestorage"
	"github.com/aripalo/aws-mfa-credential-process/internal/utils"
)

// sharedInitialization performs initialization logic required by all actions
func sharedInitialization(commandName string, verbose bool, disableDialog bool) {
	if verbose {
		utils.PrintBanner(commandName)
	}
	config.Init()
	securestorage.Init(disableDialog)
}
