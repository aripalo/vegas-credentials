package awscreds

import (
	"github.com/aripalo/aws-mfa-credential-process/internal/application/assume/awscreds/response"
	"github.com/aripalo/aws-mfa-credential-process/internal/data"
	"github.com/aripalo/aws-mfa-credential-process/internal/logger"
)

// GetCredentials handles everything in regards of getting Temporary Session Credentials (both from cache and STS)
func GetCredentials(d data.Provider) error {

	var r *response.Response
	var err error

	p := d.GetProfile()
	logger.Infoln(d, "👷", "Role", p.AssumeRoleArn)

	r, err = getCachedCredentials(d)
	if err != nil {
		logger.Infof(d, "ℹ️ ", "Credentials", "Cached: %s", err.Error())
		r, err = getNewCredentials(d)
		if err != nil {
			logger.Errorln(d, "ℹ️ ", "Credentials", err.Error())
			return err
		} else {
			logger.Successln(d, "ℹ️ ", "Credentials", "Received from STS")
			logger.PrintRuler(d, "=")
			err = r.Output()
			return err

		}
	} else {
		logger.Successln(d, "ℹ️ ", "Credentials", "Received from Cache")
		logger.PrintRuler(d, "=")
		err = r.Output()
		return err
	}

}
