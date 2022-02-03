package ssh

import (
	"errors"

	"github.com/mobile-core/rcmd/pkg/log"
	"golang.org/x/crypto/ssh"
)

// nodeStruct returns the SshMethod.
func execActorStruct() SshMethod {
	return &execActor{}
}

// Authentication returns configuration for authentication of ssh sessions.
func (e *execActor) Authentication() ([]*ssh.ClientConfig, error) {
	var clientConfig []*ssh.ClientConfig

	for i := 0; i < len(e.host); i++ {
		if e.publicKey[i] != "" {
			cfg, err := sshPublicKeyAuthentication(e.user[i], e.publicKey[i], e.password[i])
			if err != nil {
				return clientConfig, err
			}
			clientConfig = append(clientConfig, cfg)
		} else {
			cfg, err := sshPasswordAuthentication(e.user[i], e.password[i])
			if err != nil {
				return clientConfig, err
			}
			clientConfig = append(clientConfig, cfg)
		}
	}
	return clientConfig, nil
}

// Connect creates SSH sessions.
func (e *execActor) Connect(sshConfig []*ssh.ClientConfig) ([]*ssh.Session, error) {
	var sessions []*ssh.Session
	for i := 0; i < len(e.host); i++ {
		session, err := createSshSession(e.host[i], e.port[i], sshConfig[i])
		if err != nil {
			return sessions, err
		}
		sessions = append(sessions, session)
	}
	return sessions, nil
}

// Run starts remote shell sessions non-interactively, and enter the command to hosts.
func (e *execActor) Run(sessions []*ssh.Session) error {
	var cnt = 0
	loggerFactory := log.NewLoggerFactory()

	for _, session := range sessions {
		if e.command != "" {
			logger := loggerFactory.NewLogger(e.host[cnt])
			nonInteractiveShellCalling(session, e.command, logger)
		} else {
			return errors.New("please type any commands")
		}
		cnt++
	}
	return nil
}
