package controllers

import (
	"context"
	"fmt"
	benzaiten "github.com/charmelionag/cloudcontroller/api/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

type GCPKubernetesClusterReconciler struct {
	client.Client
	Scheme        *runtime.Scheme
	eventRecorder record.EventRecorder
	cloud         CloudProviders
}

func (cr *GCPKubernetesClusterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx).WithValues("gcpkubernetescluster", req.NamespacedName)

	gk := benzaiten.GCPKubernetesCluster{}

	err := cr.Get(ctx, req.NamespacedName, &gk)
	if err != nil {
		if kerr.IsNotFound(err) {
			logger.Info("gcpkubernetescluster not found")
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	// Check if GCP kubernetes cluster exists
	gke, err := cr.cloud.GCP.GetCluster(gk.Spec.Zone, gk.Spec.Name)
	if err != nil {
		logger.Error(err, "error getting GCP Kubernetes Cluster: %v", err)
		fmt.Printf("%+v\n", gke)
		return ctrl.Result{}, err
	}

	logger.Info("gcp kubernetes cluster reconciled")
	cr.eventRecorder.Event(&gk, "Normal", "Reconciled", "GCP Kubernetes Cluster reconciled")

	return ctrl.Result{}, nil
}

func (cr *GCPKubernetesClusterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&benzaiten.GCPKubernetesCluster{}).
		Complete(cr)
}

func setupGCPKubernetesClusterController(mgr manager.Manager, cp CloudProviders) error {
	eventRecorder := mgr.GetEventRecorderFor("gcpkubernetescluster")
	cc := GCPKubernetesClusterReconciler{
		Client:        mgr.GetClient(),
		Scheme:        mgr.GetScheme(),
		eventRecorder: eventRecorder,
		cloud:         cp,
	}

	// create GCPKubernetesCluster controller
	err := cc.SetupWithManager(mgr)
	if err != nil {
		return fmt.Errorf("unable to create GCPKubernetesCluster controller: %w", err)
	}

	return nil
}
