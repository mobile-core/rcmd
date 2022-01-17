package ssh

import (
	"io/ioutil"

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

// sshPublicKeyAuthention configures public-key authentication for SSH client.
func sshPublicKeyAuthentication(userName string, publicKey string, passphrase string) (*ssh.ClientConfig, error) {
	var sshconfig *ssh.ClientConfig
	buf, err := ioutil.ReadFile(publicKey)
	if err != nil {
		return sshconfig, err
	}

	// This function doesn't work properly with passphrase.
	// If you have a passphrase, please use "ParsePrivateKeyWithPassphrase".
	var key ssh.Signer
	key, err = ssh.ParsePrivateKey(buf)
	if err != nil {
		key, err = sshPublicKeyAuthenticationWithPassphrase(buf, []byte(passphrase))
		if err != nil {
			return sshconfig, err
		}
	}

	config := &ssh.ClientConfig{
		User: userName,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return config, nil
}

// sshPublicKeyAuthenticationWithPassphrase configures public-key authentication with a passphrase for SSH client.
func sshPublicKeyAuthenticationWithPassphrase(buf []byte, passphrase []byte) (ssh.Signer, error) {
	var signer ssh.Signer
	signer, err := ssh.ParsePrivateKeyWithPassphrase(buf, passphrase)
	if err != nil {
		return signer, err
	}
	return signer, nil
}
