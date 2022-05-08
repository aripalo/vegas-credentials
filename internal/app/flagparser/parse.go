package flagparser

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Parse[T any](f T, cmd *cobra.Command) (T, error) {
	var err error

	// New viper instance
	v := viper.New()

	// Read CLI flags
	// https://github.com/spf13/viper#working-with-flags
	err = v.BindPFlags(cmd.Flags())
	if err != nil {
		return f, err
	}

	// Unmarshal viper configuration into config.Config
	err = v.Unmarshal(&f, func(cfg *mapstructure.DecoderConfig) { cfg.ErrorUnset = true })

	return f, err

}
