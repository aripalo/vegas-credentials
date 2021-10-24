package logger

import (
	"github.com/aripalo/vegas-credentials/internal/data"
	"github.com/aripalo/vegas-credentials/internal/utils"
)

func DebugJSON(d data.Provider, emoji string, prefix string, message interface{}) {
	c := d.GetConfig()
	if c.Debug {
		output, err := utils.PrettyJSON(message)
		if err != nil {
			panic(err)
		}

		PrintRuler(d, "-")
		Debugf(d, emoji, prefix, "\n%s\n", output)
		PrintRuler(d, "-")
	}
}
