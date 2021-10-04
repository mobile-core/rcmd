package cmd

import (
	"github.com/spf13/cobra"
)

var nodeCmd = &cobra.Command{}

func init() {
	nodeCmd.Use = "node"
	nodeCmd.Short = "rcmd to manipulate k8s resources"
	nodeCmd.Version = "0.1"

	nodeCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if true {
			return nodeCmd.Help()
		}

		return nil
	}
	rootCmd.AddCommand(nodeCmd)
}

func init() {
	var listCmd = &cobra.Command{}
	listCmd.Use = "list"
	listCmd.Short = "list command"
	listCmd.Example = rootExample
	listCmd.Version = "0.1"

	listCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if true {
			return listCmd.Help()
		}

		return nil
	}
	nodeCmd.AddCommand(listCmd)
}

func init() {
	var cliCmd = &cobra.Command{}
	cliCmd.Use = "cli"
	cliCmd.Short = "cli command"
	cliCmd.Example = rootExample
	cliCmd.Version = "0.1"

	cliCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if true {
			return cliCmd.Help()
		}

		return nil
	}
	nodeCmd.AddCommand(cliCmd)
}
