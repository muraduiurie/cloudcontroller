package controllers

import (
	cloudv1 "github.com/charmelionag/cloudcontroller/api/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	Scheme   = runtime.NewScheme()
	SetupLog = ctrl.Log.WithName("setup")
)

// initiate the program by creating the scheme
func init() {
	utilruntime.Must(cloudv1.AddToScheme(Scheme))
}
