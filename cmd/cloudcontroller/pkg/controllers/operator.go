package controllers

import (
	"context"
	"fmt"
	"github.com/charmelionag/cloudcontroller/pkg/cloudproviders/gcp"
	"gopkg.in/yaml.v3"
	"k8s.io/client-go/rest"
	"log"
	"os"
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

	// initiate the context
	ctx := ctrl.SetupSignalHandler()
	err = runControllers(ctx, mgr)
	if err != nil {
		SetupLog.Error(err, "unable to run controllers")
		return err
	}

	return nil
}

type CloudProviders struct {
	GCP *gcp.API
}

func runControllers(ctx context.Context, mgr manager.Manager) error {
	configs, err := loadConfigs()
	if err != nil {
		return fmt.Errorf("unable to load configs: %w", err)
	}
	if configs.Controller.GcpSaFilePath == "" {
		return fmt.Errorf("GCP service account file path not set")
	}

	gcpApi, err := gcp.NewAPI(ctx, configs.Controller.GcpSaFilePath)
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

func loadConfigs() (*Configs, error) {
	var configs Configs
	// Read YAML file
	data, err := os.ReadFile(os.Getenv("CONFIG_PATH"))
	if err != nil {
		return nil, fmt.Errorf("failed reading file: %w", err)
	}

	// Unmarshal YAML into struct
	if err := yaml.Unmarshal(data, &configs); err != nil {
		return nil, fmt.Errorf("failed unmarshalling YAML: %w", err)
	}
	return &configs, nil
}
