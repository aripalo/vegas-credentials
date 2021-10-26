package logger

import (
	"github.com/aripalo/vegas-credentials/internal/interfaces"
	"github.com/aripalo/vegas-credentials/internal/utils"
)

func DebugJSON(a interfaces.AssumeCredentialProcess, emoji string, prefix string, message interface{}) {
	f := a.GetFlags()
	if f.Debug {
		output, err := utils.PrettyJSON(message)
		if err != nil {
			panic(err)
		}

		PrintRuler(a, "-")
		Debugf(a, emoji, prefix, "\n%s\n", output)
		PrintRuler(a, "-")
	}
}
