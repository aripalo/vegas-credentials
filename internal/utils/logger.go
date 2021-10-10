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

// SafeLog logs directly to tty (with stderr fallback), since aws credential_process reads stdout
func SafeLog(a ...interface{}) {
	var out io.Writer

	tty, err := tty.Open()
	if err != nil {
		out = os.Stderr
	} else {
		out = colorable.NewColorable(tty.Output())
	}
	defer tty.Close()

	fmt.Fprintln(out, a...)
}
