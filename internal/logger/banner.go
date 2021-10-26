package logger

import (
	"os"

	"github.com/aripalo/vegas-credentials/internal/config"
	"github.com/aripalo/vegas-credentials/internal/interfaces"
)

// PrintBanner prints informational header if verbose mode enabled
func PrintBanner(a interfaces.AssumeCredentialProcess, commandName string, version string) {
	f := a.GetFlags()
	if f.Verbose {
		PrintRuler(a, "=")
		Titleln(a, "", config.APP_NAME, commandName)
		Infoln(a, "- ", "Version", version)
		Infoln(a, "- ", "Repository", config.APP_REPO)

		wrapper := os.Getenv("AWS_MFA_CREDENTIAL_PROCESS_WRAPPER")

		if wrapper != "" {
			Infoln(a, "- ", "Called via Wrapper", wrapper)
		}

		PrintRuler(a, "-")
	}
}
