package ssh

import (
	"github.com/kevinburke/ssh_config"
	"github.com/mobile-core/rcmd/pkg/log"
	"golang.org/x/crypto/ssh"
)

func SshStruct(branch int) SshMethod {
	if branch == 2 {
		return connectActorStruct()
	}
	return execActorStruct()
}

func (b *baseActor) Set(
	host []string,
	port []string,
	user []string,
	password []string,
	publicKey []string,
	command string,
) {
	b.host = host
	b.port = port
	b.user = user
	b.password = password
	b.publicKey = publicKey
	b.command = command
}

// Authentication returns configuration for authentication of ssh sessions.
func (b *baseActor) Authentication() ([]*ssh.ClientConfig, error) {
	var (
		clientConfig []*ssh.ClientConfig
		host         = b.host[0]
		user         = b.user[0]
		publicKey    = b.publicKey[0]
		password     = b.password[0]
	)

	if publicKey != "" {
		if user == "" {
			user = ssh_config.Get(host, "User")
		}
		cfg, err := sshPublicKeyAuthentication(user, publicKey, password)
		if err != nil {
			return clientConfig, err
		}
		clientConfig = append(clientConfig, cfg)
	} else {
		if user == "" {
			user = ssh_config.Get(host, "User")
		}
		cfg, err := sshPasswordAuthentication(user, password)
		if err != nil {
			return clientConfig, err
		}
		clientConfig = append(clientConfig, cfg)
	}
	return clientConfig, nil
}

// Connect creates SSH sessions.
func (b *baseActor) Connect(sshConfig []*ssh.ClientConfig) ([]*ssh.Session, error) {
	var (
		host = b.host[0]
		port = b.port[0]
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
func (b *baseActor) Run(sessions []*ssh.Session) error {
	var host = b.host[0]
	loggerFactory := log.NewLoggerFactory()

	for _, session := range sessions {
		if b.command == "" {
			if err := interactiveShellCalling(session); err != nil {
				return err
			}
		} else {
			logger := loggerFactory.NewLogger(host)
			nonInteractiveShellCalling(session, b.command, logger)
		}
	}
	return nil
}
