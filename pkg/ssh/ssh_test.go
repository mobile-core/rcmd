package ssh

import (
	"testing"
)

func Test_single(t *testing.T) {
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
			name: "fail test1",
			args: args{
				command: "echo hello",
				node:    "node3",
			},
			wantErr: true,
		},
		{
			name: "fail test2",
			args: args{
				command: "hello",
				node:    "node1",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := single(tt.args.command, tt.args.node); (err != nil) != tt.wantErr {
				t.Errorf("single() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_multiple(t *testing.T) {
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
			name: "fail test1",
			args: args{
				command: "echo hello",
				node:    []string{"aaa", "node1"},
			},
			wantErr: true,
		},
		{
			name: "fail test2",
			args: args{
				command: "echo hello",
				node:    []string{"aaa", "bbb"},
			},
			wantErr: true,
		},
		{
			name: "fail test3",
			args: args{
				command: "hello",
				node:    []string{"master", "node1"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := multiple(tt.args.command, tt.args.node); (err != nil) != tt.wantErr {
				t.Errorf("multiple() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
