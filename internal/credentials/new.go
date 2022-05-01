package credentials

import (
	"os"

	"github.com/aripalo/vegas-credentials/internal/assumable"
)

// AWS_CREDENTIAL_PROCESS_VERSION defines the supported AWS credential_process version
const AWS_CREDENTIAL_PROCESS_VERSION int = 1

// New defines a response waiting to be fulfilled
func New(opts assumable.Opts) *Credentials {
	r := &Credentials{
		opts:    opts,
		output:  os.Stdout,
		cache:   NewCredentialCache(),
		Version: AWS_CREDENTIAL_PROCESS_VERSION,
	}

	return r
}

// Teardown operations for response, use with defer
func (r *Credentials) Teardown() error {
	return r.cache.Disconnect()
}
