package utils

import (
	"fmt"
	"io"
	"os"

	"github.com/mattn/go-colorable"
	"github.com/mattn/go-tty"
)

// OutputToAwsCredentialProcess prints to stdout so aws credential_process can read it
func OutputToAwsCredentialProcess(output string) {
	fmt.Fprintf(os.Stdout, output)
}

// SafeLogLn logs with newline directly to tty (with stderr fallback), since aws credential_process reads stdout
func SafeLogLn(a ...interface{}) {
	out := GetSafeWriter()
	fmt.Fprintln(out, a...)
}

// SafeLog logs without newline directly to tty (with stderr fallback), since aws credential_process reads stdout
func SafeLog(a ...interface{}) {
	out := GetSafeWriter()
	fmt.Fprint(out, a...)
}

func GetSafeWriter() io.Writer {
	var out io.Writer

	tty, err := tty.Open()
	defer tty.Close()
	if err != nil {
		out = os.Stderr
	} else {
		out = colorable.NewColorable(tty.Output())
	}

	return out
}
