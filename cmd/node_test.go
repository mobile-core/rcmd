package cmd

import (
	"reflect"
	"testing"
)

func Test_setDefaultVars(t *testing.T) {
	type args struct {
		host []string
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 []string
	}{
		{
			name: "Normal Test",
			args: args{
				host: []string{"master", "node1"},
			},
			want:  []string{"vagrant", "vagrant"},
			want1: []string{"/home/vagrant/.ssh/id_rsa", "/home/vagrant/.ssh/id_rsa"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := setDefaultVars(tt.args.host)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setDefaultVars() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("setDefaultVars() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
