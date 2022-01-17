package ssh

import (
	"testing"

	"golang.org/x/crypto/ssh"
)

func Test_sshPasswordAuthentication(t *testing.T) {
	var (
		username = "root"
		password = "r00t"
	)

	type args struct {
		userName string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    *ssh.ClientConfig
		wantErr bool
	}{
		{
			name: "Normal Test",
			args: args{
				userName: username,
				password: password,
			},
			want: &ssh.ClientConfig{
				User: username,
				Auth: []ssh.AuthMethod{
					ssh.Password(password),
				},
				HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := sshPasswordAuthentication(tt.args.userName, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("sshPasswordAuthentication() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_sshPublicKeyAuthentication(t *testing.T) {
	var (
		username  = "root"
		publicKey = "/home/vagrant/.ssh/id_rsa"
	)

	type args struct {
		userName   string
		publicKey  string
		passphrase string
	}
	tests := []struct {
		name    string
		args    args
		want    *ssh.ClientConfig
		wantErr bool
	}{
		{
			name: "Normal Test",
			args: args{
				userName:   username,
				publicKey:  publicKey,
				passphrase: "",
			},
			wantErr: false,
		},
		{
			name: "Failure Test",
			args: args{
				userName:   username,
				publicKey:  "/root/.ssh/id_abc",
				passphrase: "",
			},
			wantErr: true,
		},
		{
			name: "Failure passphrase Test",
			args: args{
				userName:   username,
				publicKey:  "/root/.ssh/id_test",
				passphrase: "testt",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := sshPublicKeyAuthentication(tt.args.userName, tt.args.publicKey, tt.args.passphrase)
			if (err != nil) != tt.wantErr {
				t.Errorf("sshPublicKeyAuthentication() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
