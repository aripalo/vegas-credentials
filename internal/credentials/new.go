package credentials

import (
	"os"

	"github.com/aripalo/vegas-credentials/internal/assumecfg"
)

// AWS_CREDENTIAL_PROCESS_VERSION defines the supported AWS credential_process version
const AWS_CREDENTIAL_PROCESS_VERSION int = 1

// New defines a response waiting to be fulfilled
func New(cfg assumecfg.AssumeCfg) *Credentials {
	r := &Credentials{
		cfg:     cfg,
		output:  os.Stdout,
		repo:    NewCredentialCache(),
		Version: AWS_CREDENTIAL_PROCESS_VERSION,
	}

	return r
}

// Teardown operations for response, use with defer
func (r *Credentials) Teardown() error {
	return r.repo.Close()
}
