package ssh

import "testing"

func Test_single(t *testing.T) {
	type args struct {
		command string
		node    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal test",
			args: args{
				command: "echo hello",
				node:    "node1",
			},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := single(tt.args.command, tt.args.node); got != tt.want {
				t.Errorf("single() = %v, want %v", got, tt.want)
			}
		})
	}
}
