package deployments

import (
	"net/http"

	dashboardclientset "github.com/tektoncd/dashboard/pkg/client/clientset/versioned"
	"k8s.io/client-go/dynamic"
	k8sclientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Options struct {
	WorkNamespace string
}

type Resource struct {
	Config          *rest.Config
	HttpClient      *http.Client
	DashboardClient dashboardclientset.Interface
	DynamicClient   dynamic.Interface
	K8sClient       k8sclientset.Interface
	Options         Options
}
