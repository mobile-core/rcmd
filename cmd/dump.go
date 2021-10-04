package cmd

import (
	"github.com/spf13/cobra"
)

var dumpCmd = &cobra.Command{}

func init() {
	dumpCmd.Use = "dump"
	dumpCmd.Short = "rcmd to manipulate k8s resources"
	dumpCmd.Version = "0.1"

	dumpCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if true {
			return dumpCmd.Help()
		}

		return nil
	}
	rootCmd.AddCommand(dumpCmd)
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
	dumpCmd.AddCommand(listCmd)
}

func init() {
	var startCmd = &cobra.Command{}
	startCmd.Use = "start"
	startCmd.Short = "cli command"
	startCmd.Example = rootExample
	startCmd.Version = "0.1"

	startCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if true {
			return startCmd.Help()
		}

		return nil
	}
	dumpCmd.AddCommand(startCmd)
}

func init() {
	var stopCmd = &cobra.Command{}
	stopCmd.Use = "stop"
	stopCmd.Short = "cli command"
	stopCmd.Example = rootExample
	stopCmd.Version = "0.1"

	stopCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if true {
			return stopCmd.Help()
		}

		return nil
	}
	dumpCmd.AddCommand(stopCmd)
}
