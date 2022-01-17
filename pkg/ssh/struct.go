package ssh

import "golang.org/x/crypto/ssh"

// SshMethod represents the procedure to establish ssh connections.
type SshMethod interface {
	Set([]string, []string, []string, []string, []string, string)
	Authentication() ([]*ssh.ClientConfig, error)
	Connect([]*ssh.ClientConfig) ([]*ssh.Session, error)
	Run([]*ssh.Session) error
}

// baseActor defines parameters for connecting to the host via ssh.
type baseActor struct {
	host      []string
	port      []string
	user      []string
	password  []string
	publicKey []string
	command   string
}

// execActor defines the configuration file.
type execActor struct {
	baseActor
}

//connectActor defines the confisuration file.
type connectActor struct {
	baseActor
}
