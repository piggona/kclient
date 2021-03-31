package deployments

import (
	"fmt"
	"net/http"

	dashboardclientset "github.com/tektoncd/dashboard/pkg/client/clientset/versioned"
	"k8s.io/client-go/dynamic"
	k8sclientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func InitResource(configPath string, namespace string) Resource {
	var cfg *rest.Config
	var err error
	cfg, err = clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		fmt.Errorf("error when creating client config: %s", err)
	}
	dashboardClient, _ := dashboardclientset.NewForConfig(cfg)
	k8sClient, _ := k8sclientset.NewForConfig(cfg)
	dynamicClient, _ := dynamic.NewForConfig(cfg)
	transport, err := rest.TransportFor(cfg)

	options := Options{
		WorkNamespace: namespace,
	}
	resource := Resource{
		Config:          cfg,
		HttpClient:      &http.Client{Transport: transport},
		DashboardClient: dashboardClient,
		DynamicClient:   dynamicClient,
		K8sClient:       k8sClient,
		Options:         options,
	}
	return resource
}
