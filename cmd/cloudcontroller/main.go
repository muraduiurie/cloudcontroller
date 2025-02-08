package main

import (
	"github.com/charmelionag/cloudcontroller/pkg/controllers"
	"k8s.io/client-go/rest"
)

func main() {
	// initiate config for in cluster client
	kubeconfig, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	err = controllers.RunOperator(kubeconfig)
	if err != nil {
		panic(err.Error())
	}
}
