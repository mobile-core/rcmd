package cfg

import (
	"reflect"
	"testing"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name string
		want node
	}{
		{
			name: "Normal Test",
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
			got, err := Load()
			if err != nil {
				t.Errorf("Load() = Error %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Load() = %v, want %v", got, tt.want)
			}
		})
	}
}
