package controllers

import (
	"fmt"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func RunOperator(kubeconfig *rest.Config) error {
	// create new manager with in cluster client and scheme
	ctrl.SetLogger(zap.New())
	mgr, err := ctrl.NewManager(kubeconfig, ctrl.Options{
		Scheme: Scheme,
	})
	if err != nil {
		SetupLog.Error(err, "unable to create controller manager")
	}

	err = runControllers(mgr)
	if err != nil {
		SetupLog.Error(err, "unable to run controllers")
		return err
	}

	return nil
}

func runControllers(mgr manager.Manager) error {
	err := setupGKEClusterController(mgr)
	if err != nil {
		return fmt.Errorf("unable to setup GKECluster controller: %w", err)
	}

	// start manager
	SetupLog.Info("starting controller manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		return fmt.Errorf("error running controller manager: %w", err)
	}

	return nil
}

func setupGKEClusterController(mgr manager.Manager) error {
	cc := GKEClusterReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}

	// create GKECluster controller
	err := cc.SetupWithManager(mgr)
	if err != nil {
		return fmt.Errorf("unable to create GKECluster controller: %w", err)
	}

	return nil
}
