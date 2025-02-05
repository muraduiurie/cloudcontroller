package controllers

import (
	"fmt"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func RunOperator(kubeconfig *rest.Config) error {
	// create new manager with client and scheme
	ctrl.SetLogger(zap.New())
	ctrl.Log.Info("setting up manager")
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
	err := setupGCPKubernetesClusterController(mgr)
	if err != nil {
		return fmt.Errorf("unable to setup GKECluster controller: %w", err)
	}

	err = setupGCPInstanceController(mgr)
	if err != nil {
		return fmt.Errorf("unable to setup GKEInstance controller: %w", err)
	}

	err = setupGCPNetworkController(mgr)
	if err != nil {
		return fmt.Errorf("unable to setup GKENetwork controller: %w", err)
	}

	// start manager
	SetupLog.Info("starting controller manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		return fmt.Errorf("error running controller manager: %w", err)
	}

	return nil
}
