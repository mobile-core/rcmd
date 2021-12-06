package ssh

import (
	"fmt"
	"os/exec"
)

func single(command string, node string) string {
	out, err := exec.Command("ssh", node, "-i", "${HOME}/.ssh/id_rsa", command).Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(out))
	}
	return string(out)
}
