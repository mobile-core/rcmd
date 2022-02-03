package cmd

import (
	"github.com/mobile-core/rcmd/pkg/cfg"
	"github.com/mobile-core/rcmd/pkg/ssh"
	"github.com/spf13/cobra"
)

type params struct {
	host      []string
	port      []string
	user      []string
	password  []string
	publicKey []string
	command   string
}

var nodeCmd = &cobra.Command{}
var branch int

func init() {
	nodeCmd.Use = "node"
	nodeCmd.Short = "node command"

	nodeCmd.RunE = func(cmd *cobra.Command, args []string) error {
		return nodeCmd.Help()
	}
	rootCmd.AddCommand(nodeCmd)
}

func init() {
	var execCmd = &cobra.Command{}
	params := params{
		host:      []string{},
		port:      []string{},
		user:      []string{},
		password:  []string{},
		publicKey: []string{},
		command:   "",
	}

	execCmd.Use = "exec"
	execCmd.Short = "Run the command to some hosts with ssh connections."
	execCmd.SilenceUsage = true
	execCmd.Flags().StringArrayVarP(&params.host, "host", "H", params.host, "")
	execCmd.Flags().StringArrayVarP(&params.port, "port", "p", params.port, "")
	execCmd.Flags().StringArrayVarP(&params.user, "user", "u", params.user, "")
	execCmd.Flags().StringArrayVarP(&params.password, "password", "P", params.password, "")
	execCmd.Flags().StringArrayVarP(&params.publicKey, "identity-file", "i", params.publicKey, "")
	execCmd.Flags().StringVarP(&params.command, "command", "c", params.command, "")

	execCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if len(params.host) == 0 && len(params.user) == 0 {
			return execCmd.Help()
		}

		defaultUser, defaultIdentityKey := setDefaultVars(params.host)

		params.host = add(params.host, "")
		params.port = addDiff(params.port, params.host, "22")
		params.password = addDiff(params.password, params.host, "")

		for i := 0; i < len(params.host); i++ {
			params.user = addDiff(params.user, params.host, defaultUser[i])
			params.publicKey = addDiff(params.publicKey, params.host, defaultIdentityKey[i])
		}

		branch = 1
		actor := ssh.SshStruct(branch)
		actor.Set(
			params.host,
			params.port,
			params.user,
			params.password,
			params.publicKey,
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

func init() {
	var connectCmd = &cobra.Command{}
	params := params{
		host:      []string{},
		port:      []string{},
		user:      []string{},
		password:  []string{},
		publicKey: []string{},
		command:   "",
	}

	connectCmd.Use = "connect"
	connectCmd.Short = "Login to the host via ssh connections."
	connectCmd.SilenceUsage = true
	connectCmd.Flags().StringArrayVarP(&params.host, "host", "H", params.host, "")
	connectCmd.Flags().StringArrayVarP(&params.port, "port", "p", params.port, "")
	connectCmd.Flags().StringArrayVarP(&params.user, "user", "u", params.user, "")
	connectCmd.Flags().StringArrayVarP(&params.password, "password", "P", params.password, "")
	connectCmd.Flags().StringArrayVarP(&params.publicKey, "identity-file", "i", params.publicKey, "")

	connectCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if len(params.host) == 0 && len(params.user) == 0 {
			return connectCmd.Help()
		}

		defaultUser, defaultIdentityKey := setDefaultVars(params.host)

		params.host = add(params.host, "")
		params.port = addDiff(params.port, params.host, "22")
		params.user = addDiff(params.user, params.host, defaultUser[0])
		params.password = addDiff(params.password, params.host, "")
		params.publicKey = addDiff(params.publicKey, params.host, defaultIdentityKey[0])

		branch = 2
		actor := ssh.SshStruct(branch)
		actor.Set(
			params.host,
			params.port,
			params.user,
			params.password,
			params.publicKey,
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
	nodeCmd.AddCommand(connectCmd)
}

func add(array []string, inStr string) []string {
	if len(array) == 0 {
		array = append(array, inStr)
	}
	return array
}

func addDiff(array, comparison []string, inStr string) []string {
	if len(array) == len(comparison) {
		return array
	}

	for i := len(array); i < len(comparison); i++ {
		array = append(array, inStr)
	}
	return array
}

func setDefaultVars(host []string) ([]string, []string) {
	var (
		defaultUser        []string
		defaultIdentityKey []string
	)

	yaml, _ := cfg.Load()

	for i := 0; i < len(host); i++ {
		for j := 0; j < len(yaml.Nodes); j++ {
			if host[i] == yaml.Nodes[j].Name {

				user := yaml.Nodes[j].User
				identityKey := yaml.Nodes[j].IdentityFile

				defaultUser = append(defaultUser, user)
				defaultIdentityKey = append(defaultIdentityKey, identityKey)
				break
			}
		}
	}
	return defaultUser, defaultIdentityKey
}
