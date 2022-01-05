package ssh

import (
	"errors"

	"github.com/mobile-core/rcmd/pkg/log"
	"golang.org/x/crypto/ssh"
)

// nodeStruct  returns the SshMethod.
func nodeStruct() SshMethod {
	return &node{}
}

// Authentication returns configuration for authentication of ssh sessions.
func (n *node) Authentication() ([]*ssh.ClientConfig, error) {
	var clientConfig []*ssh.ClientConfig

	for i := 0; i < len(n.host); i++ {
		cfg, err := sshPasswordAuthentication(n.user[i], n.password[i])
		if err != nil {
			return clientConfig, err
		}
		clientConfig = append(clientConfig, cfg)
	}
	return clientConfig, nil
}

// Connect creates SSH sessions.
func (n *node) Connect(sshConfig []*ssh.ClientConfig) ([]*ssh.Session, error) {
	var sessions []*ssh.Session
	for i := 0; i < len(n.host); i++ {
		session, err := createSshSession(n.host[i], n.port[i], sshConfig[i])
		if err != nil {
			return sessions, err
		}
		sessions = append(sessions, session)
	}
	return sessions, nil
}

// Run starts remote shell sessions interactively or non-interactively, and enter the command to hosts.
func (n *node) Run(sessions []*ssh.Session) error {
	var cnt = 0
	loggerFactory := log.NewLoggerFactory()

	for _, session := range sessions {
		if n.command != "" {
			logger := loggerFactory.NewLogger(n.host[cnt])
			nonInteractiveShellCalling(session, n.command, logger)
		} else {
			return errors.New("please type any commands")
		}
		cnt++
	}
	return nil
}
