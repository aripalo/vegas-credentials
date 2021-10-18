package response

import (
	"time"
)

// Response defines the output format expected by AWS credential_process
// https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-sourcing-external.html
type Response struct {
	Version         int       `json:"Version"`
	AccessKeyID     string    `json:"AccessKeyId"`
	SecretAccessKey string    `json:"SecretAccessKey"`
	SessionToken    string    `json:"SessionToken"`
	Expiration      time.Time `json:"Expiration"`
}

// AWS_CREDENTIAL_PROCESS_VERSION defines the supported AWS credential_process version
const AWS_CREDENTIAL_PROCESS_VERSION int = 1

// New defines a response waiting to be fulfilled
func New() *Response {
	r := &Response{
		Version: AWS_CREDENTIAL_PROCESS_VERSION,
	}
	return r
}
