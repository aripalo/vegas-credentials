package credentials

import (
	"os"
)

// AWS_CREDENTIAL_PROCESS_VERSION defines the supported AWS credential_process version
const AWS_CREDENTIAL_PROCESS_VERSION int = 1

// New defines a response waiting to be fulfilled
func New(cache StsCache, options CredentialOptions) *Credentials {
	r := &Credentials{
		options: options,
		output:  os.Stdout,
		cache:   cache,
		Version: AWS_CREDENTIAL_PROCESS_VERSION,
	}
	return r
}

// Teardown operations for response, use with defer
func (r *Credentials) Teardown() error {
	return r.cache.Disconnect()
}
