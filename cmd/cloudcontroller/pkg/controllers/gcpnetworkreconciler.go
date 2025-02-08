package controllers

import (
	"context"
	"fmt"
	cloudv1 "github.com/charmelionag/cloudcontroller/api/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

type GCPNetworkReconciler struct {
	client.Client
	Scheme        *runtime.Scheme
	eventRecorder record.EventRecorder
	cloud         CloudProviders
}

func (cr *GCPNetworkReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx).WithValues("gcpnetwork", req.NamespacedName)

	gk := cloudv1.GCPNetwork{}

	err := cr.Get(ctx, req.NamespacedName, &gk)
	if err != nil {
		if kerr.IsNotFound(err) {
			logger.Info("gcpnetwork not found")
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	// TODO: Add reconciliation logic here

	logger.Info("gcp network reconciled")
	cr.eventRecorder.Event(&gk, "Normal", "Reconciled", "GCP Network reconciled")

	return ctrl.Result{}, nil
}

func (cr *GCPNetworkReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&cloudv1.GCPNetwork{}).
		Complete(cr)
}

func setupGCPNetworkController(mgr manager.Manager, cp CloudProviders) error {
	eventRecorder := mgr.GetEventRecorderFor("gcpnetwork")
	cc := GCPNetworkReconciler{
		Client:        mgr.GetClient(),
		Scheme:        mgr.GetScheme(),
		eventRecorder: eventRecorder,
		cloud:         cp,
	}

	// create GCPNetwork controller
	err := cc.SetupWithManager(mgr)
	if err != nil {
		return fmt.Errorf("unable to create GCPNetwork controller: %w", err)
	}

	return nil
}
