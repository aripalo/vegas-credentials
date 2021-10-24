package logger

import (
	"os"

	"github.com/aripalo/vegas-credentials/internal/config"
	"github.com/aripalo/vegas-credentials/internal/data"
)

// PrintBanner prints informational header if verbose mode enabled
func PrintBanner(d data.Provider, commandName string, version string) {
	c := d.GetConfig()
	if c.Verbose {
		PrintRuler(d, "=")
		Titleln(d, "", config.APP_NAME, commandName)
		Infoln(d, "- ", "Version", version)
		Infoln(d, "- ", "Repository", config.APP_REPO)

		wrapper := os.Getenv("AWS_MFA_CREDENTIAL_PROCESS_WRAPPER")

		if wrapper != "" {
			Infoln(d, "- ", "Called via Wrapper", wrapper)
		}

		PrintRuler(d, "-")
	}
}
