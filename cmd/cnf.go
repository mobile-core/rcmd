package cmd

import (
	"github.com/mobile-core/rcmd/pkg/k8s"
	"github.com/spf13/cobra"
)

type Params struct {
	namespace string
	while     bool
}

var cnfCmd = &cobra.Command{}

func init() {
	cnfCmd.Use = "cnf"
	cnfCmd.Short = "rcmd to manipulate k8s resources"
	cnfCmd.Example = rootExample
	cnfCmd.Version = "0.1"

	cnfCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if true {
			return cnfCmd.Help()
		}

		return nil
	}
	rootCmd.AddCommand(cnfCmd)
}

func init() {
	var listCmd = &cobra.Command{}
	listCmd.Use = "list"
	listCmd.Short = "rcmd to manipulate k8s resources"
	listCmd.Example = rootExample
	listCmd.Version = "0.1"

	params := Params{
		namespace: "kube-system",
		while:     false,
	}

	listCmd.Flags().StringVarP(&params.namespace, "namespace", "n", params.namespace, "")
	listCmd.Flags().BoolVar(&params.while, "while", params.while, "")

	listCmd.RunE = func(cmd *cobra.Command, args []string) error {
		k8s.GetPodsList(params.namespace)
		// if true {
		// 	return listCmd.Help()
		// }

		return nil
	}
	cnfCmd.AddCommand(listCmd)
}

func init() {
	var deployCmd = &cobra.Command{}
	deployCmd.Use = "deploy"
	deployCmd.Short = "rcmd to manipulate k8s resources"
	deployCmd.Example = rootExample
	deployCmd.Version = "0.1"

	deployCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if true {
			return deployCmd.Help()
		}

		return nil
	}
	cnfCmd.AddCommand(deployCmd)
}

func init() {
	var deleteCmd = &cobra.Command{}
	deleteCmd.Use = "delete"
	deleteCmd.Short = "rcmd to manipulate k8s resources"
	deleteCmd.Example = rootExample
	deleteCmd.Version = "0.1"

	deleteCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if true {
			return deleteCmd.Help()
		}

		return nil
	}
	cnfCmd.AddCommand(deleteCmd)
}

func init() {
	var logCmd = &cobra.Command{}
	logCmd.Use = "log"
	logCmd.Short = "rcmd to manipulate k8s resources"
	logCmd.Example = rootExample
	logCmd.Version = "0.1"

	logCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if true {
			return logCmd.Help()
		}

		return nil
	}
	cnfCmd.AddCommand(logCmd)
}

func init() {
	var cliCmd = &cobra.Command{}
	cliCmd.Use = "cli"
	cliCmd.Short = "rcmd to manipulate k8s resources"
	cliCmd.Example = rootExample
	cliCmd.Version = "0.1"

	cliCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if true {
			return cliCmd.Help()
		}

		return nil
	}
	cnfCmd.AddCommand(cliCmd)
}

func init() {
	var execCmd = &cobra.Command{}
	execCmd.Use = "exec"
	execCmd.Short = "rcmd to manipulate k8s resources"
	execCmd.Example = rootExample
	execCmd.Version = "0.1"

	execCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if true {
			return execCmd.Help()
		}

		return nil
	}
	cnfCmd.AddCommand(execCmd)
}
