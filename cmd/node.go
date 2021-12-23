package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/mobile-core/rcmd/pkg/ssh"
	"github.com/spf13/cobra"
)

var params struct {
	command string
	node []string
}

var nodeCmd = &cobra.Command{
	Use:	"node [NODE]",
	Args:  cobra.MinimumNArgs(1),
	Short: "node command",
	Long: "node command",
	RunE: func(cmd *cobra.Command, args []string) error {
		ssh.Multiple(command strinargs, node []string)
		if true {
			return nodeCmd.Help()
		}
		fmt.Println(params.command)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(nodeCmd)
	nodeCmd.Flags().StringVarP(&params.command, "command", "c", "echo hello", "type anything")
}
