package assume

import (
	"github.com/aripalo/vegas-credentials/internal/interfaces"
	"github.com/aripalo/vegas-credentials/internal/logger"
	"github.com/aripalo/vegas-credentials/internal/response"
)

// getCredentials handles everything in regards of getting Temporary Session Credentials (both from cache and STS)
func getCredentials(a interfaces.AssumeCredentialProcess) error {

	r := response.New()
	var err error

	defer func() {
		err := r.Teardown()
		if err != nil {
			panic(err)
		}
	}()

	p := a.GetProfile()
	logger.Debugln(a, "👷", "Role", p.Target.RoleArn)

	err = r.GetCachedCredentials(a)
	if err != nil {
		logger.Debugf(a, "ℹ️ ", "Credentials", "Cached: %s\n", err.Error())
		err = r.GetNewCredentials(a)
		if err != nil {
			logger.Errorln(a, "ℹ️ ", "Credentials", err.Error())
			return err
		} else {
			logger.Debugln(a, "ℹ️ ", "Credentials", "Received from STS")
			logger.PrintRuler(a, "=")
			err = r.Output()
			return err

		}
	} else {
		logger.Debugln(a, "ℹ️ ", "Credentials", "Received from Cache")
		logger.PrintRuler(a, "=")
		err = r.Output()
		return err
	}

}
