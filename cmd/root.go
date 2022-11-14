package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	configPath string

	rootCmd = &cobra.Command{
		Use: "config",
		Run: func(cmd *cobra.Command, args []string) {},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&configPath, "path", "p", "config.yaml", "config file (default is config.yaml)")
}

func initConfig() {
	if _, err := NewConfig(configPath); err != nil {
		panic(fmt.Sprintf("new config error: %v", err))
	}
}
