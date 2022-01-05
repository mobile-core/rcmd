package ssh

import "golang.org/x/crypto/ssh"

// SshMethod represents the procedure to establish ssh connections.
type SshMethod interface {
	Set([]string, []string, []string, []string, string)
	Authentication() ([]*ssh.ClientConfig, error)
	Connect([]*ssh.ClientConfig) ([]*ssh.Session, error)
	Run([]*ssh.Session) error
}

// sshConfig defines parameters for connecting to the host via ssh.
type sshConfig struct {
	host     []string
	port     []string
	user     []string
	password []string
	command  string
}

// node defines the configuration file.
type node struct {
	sshConfig
}
