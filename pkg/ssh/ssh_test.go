package ssh

import (
	"testing"
)

func Test_Single(t *testing.T) {
	type args struct {
		command string
		node    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal test",
			args: args{
				command: "echo hello",
				node:    "node1",
			},
			wantErr: false,
		},
		{
			name: "invalid node name",
			args: args{
				command: "echo hello",
				node:    "node3",
			},
			wantErr: true,
		},
		{
			name: "invalid command",
			args: args{
				command: "hello",
				node:    "node1",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Single(tt.args.command, tt.args.node)
		})
	}
}

func Test_Multiple(t *testing.T) {
	type args struct {
		command string
		node    []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal test",
			args: args{
				command: "echo hello",
				node:    []string{"master", "node1"},
			},
			wantErr: false,
		},
		{
			name: "invalid node name",
			args: args{
				command: "echo hello",
				node:    []string{"node1", "node3"},
			},
			wantErr: true,
		},
		{
			name: "invalid command",
			args: args{
				command: "hello",
				node:    []string{"master", "node1"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Multiple(tt.args.command, tt.args.node)
		})
	}
}
