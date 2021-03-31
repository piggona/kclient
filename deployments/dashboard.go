package deployments

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetDashboardVersion(r Resource) string {
	listOptions := metav1.ListOptions{
		LabelSelector: "app=tekton-dashboard",
	}
	deployments, _ := r.K8sClient.AppsV1().Deployments(r.Options.WorkNamespace).List(context.TODO(), listOptions)
	for _, deployment := range deployments.Items {
		deploymentLabels := deployment.GetLabels()
		version, _ := deploymentLabels["version"]
		return version
	}
	return ""
}
