package cmd

import (
	"fmt"

	"github.com/aripalo/vegas-credentials/internal/application"
	"github.com/aripalo/vegas-credentials/internal/flag"
	"github.com/aripalo/vegas-credentials/internal/logger"
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
	PreRun: func(cmd *cobra.Command, args []string) {
		logger.Trace(fmt.Sprintf("%s cmd init", cmd.Name()))
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		logger.Trace(fmt.Sprintf("%s cmd done", cmd.Name()))
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		g, err := flag.Parse(application.GlobalFlags{}, cmd)
		if err != nil {
			return err
		}

		f, err := flag.Parse(application.CacheFlags{}, cmd)
		if err != nil {
			return err
		}

		app := application.New(g)
		return app.CacheClean(f)
	},
}
