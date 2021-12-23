package ssh

import (
	"fmt"
	"os/exec"
)

func Single(command string, node string) {
	out, err := exec.Command("ssh", node, "-i", "${HOME}/.ssh/id_rsa", command).Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

func Multiple(command string, node []string) {
	for _, v := range node {
		Single(command, v)
	}
}
