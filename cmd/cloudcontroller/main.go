package main

import (
	"github.com/muraduiurie/cloudcontroller/pkg/controllers"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

func main() {
	// create main logger
	logger := zap.New()
	ctrl.SetLogger(logger)
	log := ctrl.Log.WithName("main")
	log.Info("set up manager")

	// create manager
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme: controllers.Scheme,
	})
	if err != nil {
		log.Error(err, "unable to create controller manager")
	}

	// load appconfigs
	ac, err := controllers.LoadAppConfigs(log.WithName("appconfigs"))

	// run controllers
	ctx := ctrl.SetupSignalHandler()
	err = controllers.RunControllers(ctx, log.WithName("controllers"), mgr, ac)
	if err != nil {
		log.Error(err, "unable to run controllers")
	}
}
