package cmd

import (
	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{}

func init() {
	repoCmd.Use = "repo"
	repoCmd.Short = "rcmd to manipulate k8s resources"
	repoCmd.Version = "0.1"

	repoCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if true {
			return repoCmd.Help()
		}

		return nil
	}
	rootCmd.AddCommand(repoCmd)
}
