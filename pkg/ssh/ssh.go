package ssh

import (
	"fmt"
	"os/exec"
)

func single(command string, node string) error {
	out, err := exec.Command("ssh", node, "-i", "${HOME}/.ssh/id_rsa", command).Output()
	if err != nil {
		return err
	}
	fmt.Println(string(out))
	return nil
}

func multiple(command string, node []string) error {
	for _, v := range node {
		single(command, v)
	}
	return nil
}
