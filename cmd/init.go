package cmd

import (
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{}

func init() {
	initCmd.Use = "init"
	initCmd.Short = "rcmd to manipulate k8s resources"
	initCmd.Version = "0.1"

	initCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if true {
			return initCmd.Help()
		}

		return nil
	}
	rootCmd.AddCommand(initCmd)
}
