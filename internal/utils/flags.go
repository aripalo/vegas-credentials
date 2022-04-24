package utils

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func ParseFlags[T any](f T, cmd *cobra.Command) (T, error) {
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
	err = v.Unmarshal(&f)

	return f, err

}
