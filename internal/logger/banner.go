package logger

import (
	"github.com/aripalo/aws-mfa-credential-process/internal/config"
	"github.com/aripalo/aws-mfa-credential-process/internal/data"
)

// PrintBanner prints informational header if verbose mode enabled
func PrintBanner(d data.Provider, commandName string, version string) {
	c := d.GetConfig()
	if c.Verbose {
		PrintRuler(d, "=")
		Titleln(d, "", config.APP_NAME, commandName)
		Infoln(d, "- ", "Version", version)
		Infoln(d, "- ", "Repository", config.APP_REPO)
		PrintRuler(d, "-")
	}
}
