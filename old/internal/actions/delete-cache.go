package actions

import (
	"github.com/aripalo/aws-mfa-credential-process/internal/cache"
	"github.com/aripalo/aws-mfa-credential-process/internal/config"
	"github.com/urfave/cli/v2"
)

// DeleteCache performs cache cleaning (eiter the whole cache or items related to specific profile)
func DeleteCache(c *cli.Context) error {
	flags := config.ParseDeleteCacheFlags(c)
	sharedInitialization(c.Command.Name, flags.Verbose, flags.DisableDialog)
	return cache.RemoveAll(flags.ProfileName)
}
