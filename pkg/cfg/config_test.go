package cfg

import (
	"reflect"
	"testing"
)

func Test_getFileName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Normal Test",
			want: "/home/vagrant/.rcmd.yml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFileName(); got != tt.want {
				t.Errorf("getFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_configLoad(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want node
	}{
		{
			name: "Normal Test",
			args: args{
				fileName: "/home/vagrant/.rcmd.yml",
			},
			want: node{
				[]nodes{
					{HostName: "master", UserName: "vagrant", Address: "172.16.33.11"},
					{HostName: "node1", UserName: "vagrant", Address: "172.16.33.12"},
					{HostName: "node2", UserName: "vagrant", Address: "172.16.33.13"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := configLoad(tt.args.fileName)
			if err != nil {
				t.Errorf("configLoad() = Error %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("configLoad() = %v, want %v", got, tt.want)
			}
		})
	}
}
