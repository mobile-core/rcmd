package ssh

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	host := "target.host.local"
	port := "22"
	user := "user"
	pass := "password"

	// Create sshClientConfig
	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// SSH connect.
	client, err := ssh.Dial("tcp", host+":"+port, sshConfig)

	// Create Session
	session, err := client.NewSession()
	defer session.Close()

	// キー入力を接続先が認識できる形式に変換する(ここがキモ)
	fd := int(os.Stdin.Fd())
	state, err := terminal.MakeRaw(fd)
	if err != nil {
		fmt.Println(err)
	}
	defer terminal.Restore(fd, state)

	// ターミナルサイズの取得
	w, h, err := terminal.GetSize(fd)
	if err != nil {
		fmt.Println(err)
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	err = session.RequestPty("xterm", h, w, modes)
	if err != nil {
		fmt.Println(err)
	}

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	err = session.Shell()
	if err != nil {
		fmt.Println(err)
	}

	err = session.Wait()
	if err != nil {
		fmt.Println(err)
	}
}
