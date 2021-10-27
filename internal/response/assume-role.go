package response

import (
	"github.com/aripalo/vegas-credentials/internal/interfaces"
	"github.com/aripalo/vegas-credentials/internal/logger"
	"github.com/aripalo/vegas-credentials/internal/sts"
)

var getAssumedCredentials = sts.GetAssumedCredentials

// Assume IAM Role and fetch temporary session credentials to be used in credential_process
func (r *Response) AssumeRole(a interfaces.AssumeCredentialProcess) error {

	var err error

	value, expiration, err := getAssumedCredentials(a)
	if err != nil {
		return err
	}

	r.Version = AWS_CREDENTIAL_PROCESS_VERSION
	r.AccessKeyID = value.AccessKeyID
	r.SecretAccessKey = value.SecretAccessKey
	r.SessionToken = value.SessionToken
	r.Expiration = expiration

	logger.DebugJSON(a, "🔧 ", "New Response", r)

	return nil
}
