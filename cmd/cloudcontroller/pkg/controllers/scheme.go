package controllers

import (
	benzaiten "github.com/muraduiurie/cloudcontroller/api/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

var (
	Scheme = runtime.NewScheme()
)

// initiate the program by creating the scheme
func init() {
	utilruntime.Must(benzaiten.AddToScheme(Scheme))
}
