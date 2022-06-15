package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	err := run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
	}
}

func run(args []string) error {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	var port, kubeconfig string
	flag.StringVar(&port, "port", "8080", "server port")
	flag.StringVar(&kubeconfig, "kubeconfig", filepath.Join(homedir, ".kube", "config"), "path to the kubeconfig file")
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return err
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	// add handlers
	http.Handle("/healthz", &healthzHandler{clientset: clientset})

	return http.ListenAndServe(":"+port, nil)
}

// healthzHandler is an HTTP handler for the healthz API.
type healthzHandler struct {
	clientset *kubernetes.Clientset
}

// ServeHTTP implements http.Handler
func (h *healthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pods, err := h.clientset.CoreV1().Pods("").List(r.Context(), metav1.ListOptions{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	fmt.Fprintf(w, "healthz: OK\nThere are %d pods in the cluster\n", len(pods.Items))
}
