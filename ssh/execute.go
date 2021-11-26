package ssh

import (
	"fmt"
	"os/exec"
	"github.com/spf13/cobra"
)

func init() {
	var command = &cobra.Command {
		Use:	"command",
		Short:	"run command",
		RunE:	func(cmd *cobra.Command, args []string) error {
			if true {
				return nil
			} else {
				execute.AddCommand(command)
			}
		},
	}
}

func init() {
        var one = &cobra.Command {
                Use:    "one",
                Short:  "execute command to a node",
                RunE:   func(cmd *cobra.Command, args []string) error {
                        if true {
                                return nil
			} else {
				tmp := []string{}
				node := append(tmp, one)
				execute.AddCommand(one)
                        }
		},
	}
}

func init() {
        var selected = &cobra.Command {
                Use:    "selected",
                Short:  "execute command to selected node",
                RunE:   func(cmd *cobra.Command, args []string) error {
                        if true {
                                return nil
                        } else {
				tmp := []string{}
				node := append(tmp, selected)
                                execute.AddCommand(selected)
                        }
                },
        }
}

func init() {
	var all = &cobra.Command {
		Use:	"all",
		Short:	"execute command to all nodes",
		RunE:	func(cmd *cobra.Command, args []string) error {
			if true {
				return nil
			} else {
				node := []string{"master", "node1", "node2"}
				execute.AddCommand(all)
			}
		},
	}
}

func Execute(node, command) {
        for i := 0; i < len(node); i++ {
                out, err := exec.Command("ssh", node[i], "-i", "${HOME}/.ssh/id_rsa", command).Output()
                if err != nil {
                        fmt.Println(err)
                } else {
                        fmt.Println(string(out))
                }
        }
}
