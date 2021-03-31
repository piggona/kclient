package main

import (
	"fmt"

	"github.com/piggona/kclient/deployments"
)

func main() {
	resource := deployments.InitResource("/Users/haohao/.kube/config", "tekton-pipelines")
	fmt.Println(deployments.GetDashboardVersion(resource))
}
