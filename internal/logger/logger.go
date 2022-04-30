package logger

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/aripalo/vegas-credentials/internal/config/locations"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger func(v ...any)

var output io.Writer = func() io.Writer {
	if flag.Lookup("test.v") == nil {
		return &lumberjack.Logger{
			Filename:   filepath.Join(locations.StateDir, "application.log"),
			MaxSize:    1, // megabytes
			MaxBackups: 3,
			MaxAge:     28, // days
			LocalTime:  true,
		}
	}
	return os.Stderr
}()

var logFlag = log.Ldate | log.Ltime | log.Lshortfile

var (
	Trace Logger = defineLogger("TRACE", output)
	Debug Logger = defineLogger("DEBUG", output)
	Info  Logger = defineLogger("INFO", output)
	Warn  Logger = defineLogger("WARN", output)
	Error Logger = defineLogger("ERROR", output)
	Fatal Logger = defineLogger("FATAL", output)
)

func defineLogger(level string, out io.Writer) Logger {
	return log.New(out, fmt.Sprintf("%s: ", strings.ToUpper(level)), logFlag).Println
}
