package response

import (
	"github.com/aripalo/aws-mfa-credential-process/internal/data"
	"github.com/aripalo/aws-mfa-credential-process/internal/logger"
	"github.com/aripalo/aws-mfa-credential-process/internal/sts"
)

// Get Temporary Session Credentials response for AWS credential_process
func (r *Response) Get(d data.Provider) error {

	var err error

	value, expiration, err := sts.GetAssumedCredentials(d)
	if err != nil {
		return err
	}

	r.Version = AWS_CREDENTIAL_PROCESS_VERSION
	r.AccessKeyID = value.AccessKeyID
	r.SecretAccessKey = value.SecretAccessKey
	r.SessionToken = value.SessionToken
	r.Expiration = expiration

	logger.DebugJSON(d, "🔧 ", "New Response", r)

	return nil
}
