package ssh

import (
	"golang.org/x/crypto/ssh"
)

// sshPasswordAuthentication configures password authentication for SSH client.
func sshPasswordAuthentication(userName string, password string) (*ssh.ClientConfig, error) {
	config := &ssh.ClientConfig{
		User: userName,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return config, nil
}
