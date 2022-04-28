package msg

import (
	"github.com/aripalo/go-delightful"
	"github.com/aripalo/vegas-credentials/internal/config"
)

// Initialize
var Message delightful.Message

type Options struct {
	VerboseMode bool
	ColorMode   bool
	EmojiMode   bool
}

func Init(options Options) {
	Message = delightful.New(config.AppName)
	Message.SetSilentMode(!options.VerboseMode)
	Message.SetVerboseMode(options.VerboseMode)
	Message.SetColorMode(options.ColorMode)
	Message.SetEmojiMode(options.EmojiMode)
}
