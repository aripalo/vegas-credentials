package utils

import (
	"os"
	"vegas3/internal/msg"
)

func Bail(message string) {
	msg.Message.Failureln("❌", message)
	os.Exit(1)
}
