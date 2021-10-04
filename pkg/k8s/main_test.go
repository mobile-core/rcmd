package k8s

import (
	"testing"
	"local.packages/k8s"
)

func TestK8sMain(t *testing.T) {
	type Tests struct {
		name string
		args string
		want error
	}

	tests := Tests {
		name: "Normal Case",
		args: "kube-system",
		want: nil,
	}

	t.Run(tests.name, func(t *testing.T) {
		res := k8s.GetPodsList(tests.args)
		if res != tests.want {
			t.Errorf("The result is not the expected value.: %s", res)
		}
	})
}
