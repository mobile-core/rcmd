package cmd

import (
	"github.com/mobile-core/rcmd/pkg/ssh"
	"github.com/spf13/cobra"
)

type params struct {
	host     []string
	port     []string
	user     []string
	password []string
	command  string
}

var nodeCmd = &cobra.Command{}

func init() {
	nodeCmd.Use = "node"
	nodeCmd.Short = "node command"
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
	var execCmd = &cobra.Command{}
	params := params{
		host:     []string{},
		port:     []string{},
		user:     []string{},
		password: []string{},
		command:  "",
	}

	execCmd.Use = "exec"
	execCmd.Short = "Run the command to some hosts with ssh connections."
	execCmd.Version = "0.1"
	execCmd.SilenceUsage = true
	execCmd.Flags().StringArrayVarP(&params.host, "host", "H", params.host, "")
	execCmd.Flags().StringArrayVarP(&params.port, "port", "p", params.port, "")
	execCmd.Flags().StringArrayVarP(&params.user, "user", "u", params.user, "")
	execCmd.Flags().StringArrayVarP(&params.password, "password", "P", params.password, "")
	execCmd.Flags().StringVarP(&params.command, "command", "c", params.command, "")

	execCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if len(params.host) == 0 && len(params.user) == 0 {
			return execCmd.Help()
		}

		add := func(array []string, inStr string) []string {
			if len(array) == 0 {
				array = append(array, inStr)
			}
			return array
		}

		addDiff := func(array, comparison []string, inStr string) []string {
			if len(array) == len(comparison) {
				return array
			}

			for i := len(array); i < len(comparison); i++ {
				array = append(array, inStr)
			}
			return array
		}

		params.host = add(params.host, "")
		params.port = addDiff(params.port, params.host, "22")
		params.user = addDiff(params.user, params.host, "vagrant")
		params.password = addDiff(params.password, params.host, "vagrant")

		actor := ssh.SshStruct(params.host)
		actor.Set(
			params.host,
			params.port,
			params.user,
			params.password,
			params.command,
		)

		config, err := actor.Authentication()
		if err != nil {
			return err
		}

		sessions, err := actor.Connect(config)
		if err != nil {
			return err
		}

		for _, session := range sessions {
			defer session.Close()
		}

		if err := actor.Run(sessions); err != nil {
			return err
		}
		return nil
	}
	nodeCmd.AddCommand(execCmd)
}
