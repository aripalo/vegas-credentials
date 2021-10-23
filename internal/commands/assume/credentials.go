package assume

import (
	"github.com/aripalo/aws-mfa-credential-process/internal/data"
	"github.com/aripalo/aws-mfa-credential-process/internal/logger"
	"github.com/aripalo/aws-mfa-credential-process/internal/response"
)

// getCredentials handles everything in regards of getting Temporary Session Credentials (both from cache and STS)
func getCredentials(d data.Provider) error {

	r := response.New()
	var err error

	defer func() error {
		err := r.Teardown()
		if err != nil {
			return err
		}
		return nil
	}()

	p := d.GetProfile()
	logger.Debugln(d, "👷", "Role", p.RoleArn)

	err = r.GetCachedCredentials(d)
	if err != nil {
		logger.Debugf(d, "ℹ️ ", "Credentials", "Cached: %s\n", err.Error())
		err = r.GetNewCredentials(d)
		if err != nil {
			logger.Errorln(d, "ℹ️ ", "Credentials", err.Error())
			return err
		} else {
			logger.Debugln(d, "ℹ️ ", "Credentials", "Received from STS")
			logger.PrintRuler(d, "=")
			err = r.Output()
			return err

		}
	} else {
		logger.Debugln(d, "ℹ️ ", "Credentials", "Received from Cache")
		logger.PrintRuler(d, "=")
		err = r.Output()
		return err
	}

}
