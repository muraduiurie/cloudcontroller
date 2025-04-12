package controllers

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	"github.com/muraduiurie/cloudcontroller/pkg/cloudproviders/gcp"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

type CloudProviders struct {
	GCP *gcp.API
}

func RunControllers(ctx context.Context, log logr.Logger, mgr manager.Manager, appconfig *AppConfigs) error {
	log.Info("setting up controllers")

	// get GCP api
	gcpApi, err := gcp.NewAPI(ctx, log.WithName("gcp"), appconfig.Controller.CloudProviderConfigs.GCP.GcpSaFilePath)
	if err != nil {
		log.Error(err, "error creating GCP API client, skipping GCP controller")
	}
	cp := CloudProviders{
		GCP: gcpApi,
	}

	// setup GCP controllers
	if gcpApi != nil {
		log.Info("setting up GCP controllers")
		err = setupGCPKubernetesClusterController(mgr, log, cp)
		if err != nil {
			return fmt.Errorf("unable to setup GKECluster controller: %w", err)
		}

		err = setupGCPInstanceController(mgr, log, cp)
		if err != nil {
			return fmt.Errorf("unable to setup GKEInstance controller: %w", err)
		}

		err = setupGCPNetworkController(mgr, log, cp)
		if err != nil {
			return fmt.Errorf("unable to setup GKENetwork controller: %w", err)
		}
	}

	// start manager
	log.Info("starting controller manager")
	if err := mgr.Start(ctx); err != nil {
		return fmt.Errorf("error running controller manager: %w", err)
	}

	return nil
}
