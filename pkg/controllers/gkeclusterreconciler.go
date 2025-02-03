package controllers

import (
	"context"
	cloudv1 "github.com/charmelionag/cloudcontroller/api/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

type GKEClusterReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (cr *GKEClusterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx).WithValues("webpage", req.NamespacedName)

	gk := cloudv1.GKECluster{}

	err := cr.Client.Get(ctx, req.NamespacedName, &gk)
	if err != nil {
		if kerr.IsNotFound(err) {
			logger.Info("webpage not found")
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	// TODO: Add reconciliation logic here

	logger.Info("gke cluster reconciled")

	return ctrl.Result{}, nil
}

func (cr *GKEClusterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&cloudv1.GKECluster{}).
		Complete(cr)
}
