package response

import (
	"github.com/aripalo/vegas-credentials/internal/data"
	"github.com/aripalo/vegas-credentials/internal/logger"
	"github.com/aripalo/vegas-credentials/internal/sts"
)

// Assume IAM Role and fetch temporary session credentials to be used in credential_process
func (r *Response) AssumeRole(d data.Provider) error {

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
