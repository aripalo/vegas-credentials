package logger

import (
	"io"
	"os"

	"github.com/mattn/go-colorable"
	"github.com/mattn/go-tty"
)

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
