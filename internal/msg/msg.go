package msg

import (
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/aripalo/go-delightful"
	"github.com/aripalo/vegas-credentials/internal/config"
	"github.com/aripalo/vegas-credentials/internal/logger"
	"github.com/enescakir/emoji"
)

// Initialize
var d delightful.Message

type Options struct {
	VerboseMode bool
	ColorMode   bool
	EmojiMode   bool
}

func Init(options Options) {
	d = delightful.New(config.AppName)
	d.SetSilentMode(!options.VerboseMode)
	d.SetVerboseMode(options.VerboseMode)
	d.SetColorMode(options.ColorMode)
	d.SetEmojiMode(options.EmojiMode)
}

func SetSilentMode(silentMode bool) {
	d.SetSilentMode(silentMode)
}

func Trace(prefix emoji.Emoji, message string) {
	logger.Trace(getFilename(), message)
}

func Debug(prefix emoji.Emoji, message string) {
	logger.Debug(getFilename(), message)
	d.Debugln(prefix, message)
}

func DebugNoLog(prefix emoji.Emoji, message string) {
	d.Debugln(prefix, message)
}

func Info(prefix emoji.Emoji, message string) {
	logger.Info(getFilename(), message)
	d.Infoln(prefix, message)
}

func Success(prefix emoji.Emoji, message string) {
	logger.Info(getFilename(), message)
	d.Successln(prefix, message)
}

func Warn(prefix emoji.Emoji, message string) {
	logger.Warn(getFilename(), message)
	d.Warningln(prefix, message)
}

func Error(prefix emoji.Emoji, message string) {
	logger.Error(getFilename(), message)
	d.Failureln(prefix, message)
}

func Prompt(prefix emoji.Emoji, message string) {
	d.Promptln(prefix, message)
}

func Fatal(message string) {
	logger.Fatal(getFilename(), message)
	d.Failureln("❌", message)
	os.Exit(1)
}

func HorizontalRuler() {
	d.HorizontalRuler()
}

// Get the original caller filename.
// Based on https://stackoverflow.com/a/58260501/11266464.
func getFilename() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		panic("Could not get context info for logger!")
	}
	filename := file[strings.LastIndex(file, "/")+1:] + ":" + strconv.Itoa(line)

	return filename
}
