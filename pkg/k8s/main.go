package k8s

import (
	"os"
	"fmt"
	"strings"
	"context"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/api/core/v1"
	// corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/rest"
	// "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func GetPodsList(namespace string) error {
	config, err := getKubeConfig()
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

	DisplayPodList(pods)
	return nil
}

func getKubeConfig() (*rest.Config, error) {
	home := os.Getenv("HOME")
	if home == "" {
		home = os.Getenv("USERPROFILE")
	}

	kubeconfig := filepath.Join(home, ".kube", "config")
	return clientcmd.BuildConfigFromFlags("", kubeconfig)
}

func DisplayPodList(pods *v1.PodList) {
	padingLen := map[string]int{
		"namespace":     0,
		"podName":       0,
		"containerName": 0,
		"status":        0,
		"podsip":        0,
	}

	for _, v := range pods.Items {
		if padingLen["namespace"] < len(v.GetNamespace()) {
			padingLen["namespace"] = len(v.GetNamespace())
		}

		if padingLen["podName"] < len(v.GetName()) {
			padingLen["podName"] = len(v.GetName())
		}

		if padingLen["status"] < len(v.Status.Phase) {
			padingLen["status"] = len(v.Status.Phase)
		}

		for _, w := range v.Spec.Containers {
			if padingLen["containerName"] < len(w.Name) {
				padingLen["containerName"] = len(w.Name)
			}
		}

		if padingLen["podsip"] < len(v.Status.PodIP) {
			padingLen["podsip"] = len(v.Status.PodIP)
		}
	}

	/*
		+-----+-----------+--------------------------+----------------+-------+----------+---------+-----------+
		| No. | NAMESPACE | APP NAME                 | CONTAINER NAME | READY | RESTARTS | STATUS  | PODS IP   |
		+-----+-----------+--------------------------+----------------+-------+----------+---------+-----------+
		|   1 | f5gc      | f5gc-mongodb-0           |                | 1/1   | 0        |         | 10.1.0.15 |
		|     |           |                          | mongodb        |       | 0        | Running |           |
		|   2 | f5gc      | f5gc-ue-7789f69d66-dkmm9 |                | 2/2   | 0        |         | 10.1.0.15 |
		|     |           |                          | f5gc-ue        | 10/10 | 0        | Running |           |
		|     |           |                          | tcpdump        |       | 0        | Running |           |
		+-----+-----------+--------------------------+----------------+-------+----------+---------+-----------+
	*/

	framePrint(padingLen)
	headerPrint(padingLen)
	framePrint(padingLen)

	for i, pod := range pods.Items {
		namespacePading := padingSet("NAMESPACE", len(pod.GetNamespace()), padingLen["namespace"])
		podNamePading   := padingSet("APP NAME", len(pod.GetName()), padingLen["podName"])
		statusPading    := padingSet("STATUS", len(pod.Status.Phase), padingLen["status"])
		podsIpPading    := padingSet("PODS IP", len(pod.Status.PodIP), padingLen["podsip"])

		for j, container := range pod.Spec.Containers {
			containerNamePading := padingSet("CONTAINER NAME", len(container.Name), padingLen["containerName"])

			readyCount := 0
			for _, v := range pod.Status.ContainerStatuses {
				readyCount += readyCounter(v.Ready)
			}

			if j == 0 {
				fmt.Printf("| %3d ", i+1) // No.
				fmt.Printf("| %s%s ", pod.GetNamespace(), strings.Repeat(" ", namespacePading)) // NAMESPACE
				fmt.Printf("| %s%s ", pod.GetName(),      strings.Repeat(" ", podNamePading)) // APP NAME
				fmt.Printf("| %s ", strings.Repeat(" ", len(container.Name) + containerNamePading)) // CONTAINER NAME
				fmt.Printf("| %2d/%-2d ", readyCount, len(pod.Status.ContainerStatuses)) // READY
				fmt.Printf("| %s ", strings.Repeat(" ", 8)) // RESTARTS
				fmt.Printf("| %s%s ", pod.Status.Phase,   strings.Repeat(" ", statusPading)) // STATUS
				fmt.Printf("| %s%s ", pod.Status.PodIP, strings.Repeat(" ", podsIpPading)) // PODS IP
				fmt.Printf("|\n")	
			}

			fmt.Printf("| %s ", strings.Repeat(" ", 3)) // No.
			fmt.Printf("| %s ", strings.Repeat(" ", len(pod.GetNamespace()) + namespacePading)) // NAMESPACE
			fmt.Printf("| %s ", strings.Repeat(" ", len(pod.GetName()) + podNamePading)) // APP NAME
			fmt.Printf("| %s%s ", container.Name, strings.Repeat(" ", containerNamePading)) // CONTAINER NAME
			fmt.Printf("| %2d/%-2d ", readyCounter(pod.Status.ContainerStatuses[j].Ready), 1) // READY
			fmt.Printf("| %s ", strings.Repeat(" ", 8)) // RESTARTS
			fmt.Printf("| %s ", strings.Repeat(" ", len(pod.Status.Phase) + statusPading)) // STATUS
			fmt.Printf("| %s ", strings.Repeat(" ", len(pod.Status.PodIP) + podsIpPading)) // PODS IP
			fmt.Printf("|\n")	
		}
	}
	framePrint(padingLen)
}

func framePrint(padingLen map[string]int) {
	var (
		numberPading        = 5
		namespacePading     = 0
		podNamePading       = 0
		containerNamePading = 0
		readyPading         = 7
		restartsPading      = 10
		statusPading        = 0
		podsIpPading        = 0
	)

	if padingLen["namespace"] > 9 {
		namespacePading = padingLen["namespace"] + 2
	} else {
		namespacePading = 11
	}

	if padingLen["podName"] > 8 {
		podNamePading = padingLen["podName"] + 2
	} else {
		podNamePading = 10
	}

	if padingLen["status"] > 6 {
		statusPading = padingLen["status"] + 2
	} else {
		statusPading = 8
	}

	if padingLen["containerName"] > 16 {
		containerNamePading = padingLen["containerName"] + 2
	} else {
		containerNamePading = 16
	}

	if padingLen["podsip"] > 7 {
		podsIpPading = padingLen["podsip"] + 2
	} else {
		podsIpPading = 9
	}

	fmt.Printf("+%s", strings.Repeat("-", numberPading))
	fmt.Printf("+%s", strings.Repeat("-", namespacePading))
	fmt.Printf("+%s", strings.Repeat("-", podNamePading))
	fmt.Printf("+%s", strings.Repeat("-", containerNamePading))
	fmt.Printf("+%s", strings.Repeat("-", readyPading))
	fmt.Printf("+%s", strings.Repeat("-", restartsPading))
	fmt.Printf("+%s", strings.Repeat("-", statusPading))
	fmt.Printf("+%s", strings.Repeat("-", podsIpPading))
	fmt.Printf("+\n")
}

func headerPrint(padingLen map[string]int) {
	var (
		namespacePading     = 0
		podNamePading       = 0
		containerNamePading = 0
		statusPading        = 0
		podsIpPading        = 0
	)

	if padingLen["namespace"] > 9 {
		namespacePading = padingLen["namespace"] - len("NAMESPACE")
	}

	if padingLen["podName"] > 8 {
		podNamePading = padingLen["podName"] - len("APP NAME")
	}

	if padingLen["status"] > 6 {
		statusPading = padingLen["status"] - len("STATUS")
	}

	if padingLen["containerName"] > 16 {
		containerNamePading = padingLen["containerName"] - len("CONTAINER NAME")
	}

	if padingLen["podsip"] > 7 {
		podsIpPading = padingLen["podsip"] - len("PODS IP")
	}

	fmt.Printf("| %3s ", "No.")
	fmt.Printf("| %s%s ", "NAMESPACE",      strings.Repeat(" ", namespacePading))
	fmt.Printf("| %s%s ", "APP NAME",       strings.Repeat(" ", podNamePading))
	fmt.Printf("| %s%s ", "CONTAINER NAME", strings.Repeat(" ", containerNamePading))
	fmt.Printf("| %5s ", "READY")
	fmt.Printf("| %s ", "RESTARTS")
	fmt.Printf("| %s%s ", "STATUS",         strings.Repeat(" ", statusPading))
	fmt.Printf("| %s%s ", "PODS IP",        strings.Repeat(" ", podsIpPading))
	fmt.Printf("|\n")
}

func padingSet(name string, length int, max int) int {
	switch (name) {
		case "CONTAINER NAME":
			if max > len(name) {
				return max - length
			}
			return max
	}
	if max > len(name) {
		return max - length
	}
	return max + 1
}

func readyCounter(ready bool) int {
	if ready {
		return 1
	}
	return 0
}
