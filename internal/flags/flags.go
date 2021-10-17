package flags

import (
	"github.com/aripalo/aws-mfa-credential-process/internal/config"
	"github.com/spf13/cobra"
)

func DefineAssumeFlags(cmd *cobra.Command, c *config.Config, required ...string) {
	cmd.Flags().StringVar(&c.Profile, "profile", c.Profile, "[Required] Which AWS Profile to use")

	for _, v := range required {
		cmd.MarkFlagRequired(v)
	}

	cmd.Flags().BoolVar(&c.Verbose, "verbose", c.Verbose, "Verbose output")
	cmd.Flags().BoolVar(&c.HideArns, "hide-arns", c.HideArns, "HideArns")
	cmd.Flags().BoolVar(&c.DisableDialog, "disable-dialog", c.DisableDialog, "DisableDialog")
	cmd.Flags().BoolVar(&c.DisableRefresh, "disable-refresh", c.DisableRefresh, "DisableRefresh")

}

func DefineDeleteCacheFlags(cmd *cobra.Command, c *config.Config) {
	cmd.Flags().StringVar(&c.Profile, "profile", c.Profile, "Which AWS Profile to use")
}
