package controllers

import (
	"context"
	"fmt"
	"github.com/charmelionag/cloudcontroller/pkg/cloudproviders/gcp"
	"github.com/charmelionag/cloudcontroller/pkg/configmap"
	"log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

type CloudProviders struct {
	GCP *gcp.API
}

func RunControllers(ctx context.Context, mgr manager.Manager, configmap *configmap.ConfigMap) error {
	gcpApi, err := gcp.NewAPI(ctx, configmap.Controller.GcpSaFilePath)
	if err != nil {
		log.Fatalf("Error creating GCP API: %v", err)
	}
	cp := CloudProviders{
		GCP: gcpApi,
	}

	err = setupGCPKubernetesClusterController(mgr, cp)
	if err != nil {
		return fmt.Errorf("unable to setup GKECluster controller: %w", err)
	}

	err = setupGCPInstanceController(mgr, cp)
	if err != nil {
		return fmt.Errorf("unable to setup GKEInstance controller: %w", err)
	}

	err = setupGCPNetworkController(mgr, cp)
	if err != nil {
		return fmt.Errorf("unable to setup GKENetwork controller: %w", err)
	}

	// start manager
	SetupLog.Info("starting controller manager")
	if err := mgr.Start(ctx); err != nil {
		return fmt.Errorf("error running controller manager: %w", err)
	}

	return nil
}
