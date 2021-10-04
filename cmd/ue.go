package cmd

import (
	"github.com/spf13/cobra"
)

var ueCmd = &cobra.Command{}

func init() {
	ueCmd.Use = "ue"
	ueCmd.Short = "rcmd to manipulate k8s resources"
	ueCmd.Version = "0.1"

	ueCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if true {
			return ueCmd.Help()
		}

		return nil
	}
	rootCmd.AddCommand(ueCmd)
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
	ueCmd.AddCommand(listCmd)
}

func init() {
	var attachCmd = &cobra.Command{}
	attachCmd.Use = "attach"
	attachCmd.Short = "cli command"
	attachCmd.Example = rootExample
	attachCmd.Version = "0.1"

	attachCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if true {
			return attachCmd.Help()
		}

		return nil
	}
	ueCmd.AddCommand(attachCmd)
}

func init() {
	var detachCmd = &cobra.Command{}
	detachCmd.Use = "detach"
	detachCmd.Short = "cli command"
	detachCmd.Example = rootExample
	detachCmd.Version = "0.1"

	detachCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if true {
			return detachCmd.Help()
		}

		return nil
	}
	ueCmd.AddCommand(detachCmd)
}

func init() {
	var dumpCmd = &cobra.Command{}
	dumpCmd.Use = "dump"
	dumpCmd.Short = "cli command"
	dumpCmd.Example = rootExample
	dumpCmd.Version = "0.1"

	dumpCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if true {
			return dumpCmd.Help()
		}

		return nil
	}
	ueCmd.AddCommand(dumpCmd)
}

func init() {
	var infoCmd = &cobra.Command{}
	infoCmd.Use = "info"
	infoCmd.Short = "cli command"
	infoCmd.Example = rootExample
	infoCmd.Version = "0.1"

	infoCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if true {
			return infoCmd.Help()
		}

		return nil
	}
	ueCmd.AddCommand(infoCmd)
}

func init() {
	var statusCmd = &cobra.Command{}
	statusCmd.Use = "status"
	statusCmd.Short = "cli command"
	statusCmd.Example = rootExample
	statusCmd.Version = "0.1"

	statusCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if true {
			return statusCmd.Help()
		}

		return nil
	}
	ueCmd.AddCommand(statusCmd)
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
	ueCmd.AddCommand(cliCmd)
}
