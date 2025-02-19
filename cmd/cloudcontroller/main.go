package main

import (
	"github.com/charmelionag/cloudcontroller/pkg/configmap"
	"github.com/charmelionag/cloudcontroller/pkg/controllers"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

func main() {
	// create new manager with client and scheme
	logger := zap.New()
	ctrl.SetLogger(logger)
	ctrl.Log.Info("setting up manager")
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme: controllers.Scheme,
	})
	if err != nil {
		controllers.SetupLog.Error(err, "unable to create controller manager")
	}
	// load configmap
	configmap, err := configmap.LoadConfigMap()
	// initiate the context and run controllers
	ctx := ctrl.SetupSignalHandler()
	err = controllers.RunControllers(ctx, mgr, configmap)
	if err != nil {
		controllers.SetupLog.Error(err, "unable to run controllers")
	}
}
