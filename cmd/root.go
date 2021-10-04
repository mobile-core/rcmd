package cmd

import (
	"os"
	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{}

const rootExample = ``

func Execute() {
	err := rootCmd.Execute()
    if err != nil {
        os.Exit(0)
    }
}

func init() {
	rootCmd.Use = "rcmd"
	rootCmd.Short = "rcmd to manipulate k8s resources"
	rootCmd.Example = rootExample
	rootCmd.Version = "0.1"

	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		return rootCmd.Help()
	}
}
