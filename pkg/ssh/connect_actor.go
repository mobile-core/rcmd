package ssh

import "golang.org/x/crypto/ssh"

//connectActorStruct returns the SshMethod.
func connectActorStruct() SshMethod {
	return &connectActor{}
}

// Run starts remote shell sessions interactively.
func (c *connectActor) Run(sessions []*ssh.Session) error {
	if err := interactiveShellCalling(sessions[0]); err != nil {
		return err
	}
	return nil
}
