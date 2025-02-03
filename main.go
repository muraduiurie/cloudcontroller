package main

import (
	"github.com/charmelionag/cloudcontroller/pkg/controllers"
	"k8s.io/client-go/rest"
)

//const (
//	gcpSaFilePath = "./creds/gcp-creds.json"
//)

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
