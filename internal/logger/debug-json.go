package logger

import (
	"github.com/aripalo/aws-mfa-credential-process/internal/data"
	"github.com/aripalo/aws-mfa-credential-process/internal/utils"
)

func DebugJSON(d data.Provider, emoji string, prefix string, message interface{}) {
	c := d.GetConfig()
	if c.Debug {
		output, err := utils.PrettyJSON(message)
		if err != nil {
			panic(err)
		}

		Debugf(d, emoji, prefix, "\n%s\n", output)
	}
}
