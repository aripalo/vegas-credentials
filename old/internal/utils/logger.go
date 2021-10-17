package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/mattn/go-colorable"
	"github.com/mattn/go-tty"
)

// OutputToAwsCredentialProcess prints to stdout so aws credential_process can read it
// https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sourcing-external.html
func OutputToAwsCredentialProcess(output json.RawMessage) {
	fmt.Fprintf(os.Stdout, string(output))
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

// GetSafeWriter tries to create a writer into tty, but if not present, use stderr.
// Reason why we can't write to stdout is that it messes up the credential_process
// (as it expects to receive temporary session credentials in stdout)
// and some AWS tools don't print out stderr from credential_process.
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
