package cmd

import (
	"fmt"

	"github.com/aripalo/vegas-credentials/internal/app"
	"github.com/aripalo/vegas-credentials/internal/app/flagparser"
	"github.com/aripalo/vegas-credentials/internal/msg"
	"github.com/spf13/cobra"
)

var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Manage cache",
	PreRun: func(cmd *cobra.Command, args []string) {
		msg.Trace("", fmt.Sprintf("%s cmd init", cmd.Name()))
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		msg.Trace("", fmt.Sprintf("%s cmd done", cmd.Name()))
	},
}

var cacheCleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean (remove) cache items",
	Long:  "By defaults cleans all caches. Use flags to control if only specific caches need to be cleaned.",
	PreRun: func(cmd *cobra.Command, args []string) {
		msg.Trace("", fmt.Sprintf("%s cmd init", cmd.Name()))
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		msg.Trace("", fmt.Sprintf("%s cmd done", cmd.Name()))
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		g, err := flagparser.Parse(app.GlobalFlags{}, cmd)
		if err != nil {
			return err
		}

		f, err := flagparser.Parse(app.CacheFlags{}, cmd)
		if err != nil {
			return err
		}

		a := app.New(g)
		return a.CacheClean(f)
	},
}
