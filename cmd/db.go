package cmd

import (
	"github.com/spf13/cobra"
)

var dbCmd = &cobra.Command{}

func init() {
	dbCmd.Use = "db"
	dbCmd.Short = "rcmd to manipulate k8s resources"
	dbCmd.Version = "0.1"

	dbCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if true {
			return dbCmd.Help()
		}

		return nil
	}
	rootCmd.AddCommand(dbCmd)
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
	dbCmd.AddCommand(listCmd)
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
	dbCmd.AddCommand(cliCmd)
}
