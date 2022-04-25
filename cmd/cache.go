package cmd

import (
	"github.com/aripalo/vegas-credentials/internal/application"
	"github.com/aripalo/vegas-credentials/internal/utils"

	"github.com/spf13/cobra"
)

var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Manage cache",
}

var cacheCleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean (remove) cache items",
	Long:  "By defaults cleans all caches. Use flags to control if only specific caches need to be cleaned.",
	RunE: func(cmd *cobra.Command, args []string) error {

		g, err := utils.ParseFlags(application.GlobalFlags{}, cmd)
		if err != nil {
			return err
		}

		f, err := utils.ParseFlags(application.CacheFlags{}, cmd)
		if err != nil {
			return err
		}

		app := application.New(g)
		return app.CacheClean(f)
	},
}