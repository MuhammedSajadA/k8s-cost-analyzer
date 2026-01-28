package k8s

import (
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

// ClientProvider is the minimal interface ListNamespaces needs.
// kubernetes.Clientset satisfies this via CoreV1().
type ClientProvider interface {
	CoreV1() corev1.CoreV1Interface
}

type K8sClients struct {
	CoreV1 corev1.CoreV1Interface
}
