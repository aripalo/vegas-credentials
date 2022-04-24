package utils

import (
	"os"

	"github.com/aripalo/vegas-credentials/internal/msg"
)

func Bail(message string) {
	msg.Message.Failureln("❌", message)
	os.Exit(1)
}
