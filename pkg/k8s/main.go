package k8s

import (
	"os"
	"fmt"
	"strings"
	"context"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GetPodsList(namespace string) error {
	home := homeDir()
	kubeconfig := filepath.Join(home, ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("failed to get pods: %s", err)
	}

	padingLen := map[string]int{
		"namespace": 0,
		"podname":   0,
		"status":    0,
	}

	for _, v := range pods.Items {
		if padingLen["namespace"] < len(v.GetNamespace()) {
			padingLen["namespace"] = len(v.GetNamespace())
		}

		if padingLen["podname"] < len(v.GetName()) {
			padingLen["podname"] = len(v.GetName())
		}

		if padingLen["status"] < len(v.Status.Phase) {
			padingLen["status"] = len(v.Status.Phase)
		}
	}

	framePrint(padingLen)
	headerPrint(padingLen)
	framePrint(padingLen)

	for i, pod := range pods.Items {
		namespacePading := padingNamespace(len(pod.GetNamespace()), padingLen["namespace"])
		podnamePading   := padingPodName(len(pod.GetName()), padingLen["podname"])
		statusPading    := padingStatus(len(pod.Status.Phase), padingLen["status"])

		fmt.Printf("| %3d ", i+1)
		fmt.Printf("| %s%s ", pod.GetNamespace(), strings.Repeat(" ", namespacePading))
		fmt.Printf("| %s%s ", pod.GetName(), strings.Repeat(" ", podnamePading))
		fmt.Printf("| %s%s ", pod.Status.Phase, strings.Repeat(" ", statusPading))
		fmt.Printf("|\n")
	}
	framePrint(padingLen)
	fmt.Println(pods.Items[1].Spec.Containers[1].Name)
	return nil
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE")
}

func framePrint(padingLen map[string]int) {
	var (
		namespacePading = 0
		podnamePading   = 0
		statusPading    = 0
	)

	if padingLen["namespace"] > 9 {
		namespacePading = padingLen["namespace"] + 2

	} else {
		namespacePading = 11
	}

	if padingLen["podname"] > 4 {
		podnamePading = padingLen["podname"] + 2

	} else {
		podnamePading = 6
	}

	if padingLen["status"] > 6 {
		statusPading = padingLen["status"] + 2

	} else {
		statusPading = 8
	}

	fmt.Printf("+%s", strings.Repeat("-", 5))
	fmt.Printf("+%s", strings.Repeat("-", namespacePading))
	fmt.Printf("+%s", strings.Repeat("-", podnamePading))
	fmt.Printf("+%s", strings.Repeat("-", statusPading))

	fmt.Printf("+\n")
}

func headerPrint(padingLen map[string]int) {
	var (
		namespacePading = 0
		podnamePading   = 0
		statusPading    = 0
	)

	if padingLen["namespace"] > 9 {
		namespacePading = padingLen["namespace"] - len("Namespace")
	}

	if padingLen["podname"] > 4 {
		podnamePading = padingLen["podname"] - len("Name")
	}

	if padingLen["status"] > 6 {
		statusPading = padingLen["status"] - len("Status")
	}

	fmt.Printf("| %3s ", "No.")
	fmt.Printf("| %s%s ", "Namespace", strings.Repeat(" ", namespacePading))
	fmt.Printf("| %s%s ", "Name",      strings.Repeat(" ", podnamePading))
	fmt.Printf("| %s%s ", "Status",    strings.Repeat(" ", statusPading))
	fmt.Printf("|\n")
}

func padingNamespace(namespaceLen int, max int) int {
	if max > len("Namespace") {
		return max - namespaceLen
	}
	return max + 1
}

func padingPodName(podnameLen int, max int) int {
	if max > len("Name") {
		return max - podnameLen
	}
	return max + 1
}

func padingStatus(statusLen int, max int) int {
	if max > len("Status") {
		return max - statusLen
	}
	return max + 1
}
