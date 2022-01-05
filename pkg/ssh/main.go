package ssh

import (
	"github.com/mobile-core/rcmd/pkg/log"
	"golang.org/x/crypto/ssh"
)

func SshStruct(host []string) SshMethod {
	return nodeStruct()
}

func (s *sshConfig) Set(
	host []string,
	port []string,
	user []string,
	password []string,
	command string,
) {
	s.host = host
	s.port = port
	s.user = user
	s.password = password
	s.command = command
}

// Authentication returns configuration for authentication of ssh sessions.
func (s *sshConfig) Authentication() ([]*ssh.ClientConfig, error) {
	var (
		clientConfig []*ssh.ClientConfig
		user         = s.user[0]
		password     = s.password[0]
	)

	cfg, err := sshPasswordAuthentication(user, password)
	if err != nil {
		return clientConfig, err
	}
	clientConfig = append(clientConfig, cfg)
	return clientConfig, nil
}

// Connect creates SSH sessions.
func (s *sshConfig) Connect(sshConfig []*ssh.ClientConfig) ([]*ssh.Session, error) {
	var (
		host = s.host[0]
		port = s.port[0]
		cfg  = sshConfig[0]
	)

	var sessions []*ssh.Session
	session, err := createSshSession(host, port, cfg)
	if err != nil {
		return sessions, err
	}
	sessions = append(sessions, session)
	return sessions, nil
}

// Run starts remote shell sessions interactively or non-interactively, and enter the command to hosts.
func (s *sshConfig) Run(sessions []*ssh.Session) error {
	var host = s.host[0]
	loggerFactory := log.NewLoggerFactory()

	for _, session := range sessions {
		if s.command == "" {
			if err := interactiveShellCalling(session); err != nil {
				return err
			}
		} else {
			logger := loggerFactory.NewLogger(host)
			nonInteractiveShellCalling(session, s.command, logger)
		}
	}
	return nil
}
