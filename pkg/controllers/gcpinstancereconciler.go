package controllers

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	benzaiten "github.com/muraduiurie/cloudcontroller/api/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

type GCPInstanceReconciler struct {
	client.Client
	Scheme        *runtime.Scheme
	eventRecorder record.EventRecorder
	cloud         CloudProviders
	Log           logr.Logger
}

func (cr *GCPInstanceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := cr.Log.WithValues("gcpinstance", req.NamespacedName)

	gk := benzaiten.GCPInstance{}
	err := cr.Get(ctx, req.NamespacedName, &gk)
	if err != nil {
		if kerr.IsNotFound(err) {
			logger.Info("gcpinstance not found")
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	// TODO: Add reconciliation logic here

	logger.Info("gcp instance reconciled")
	cr.eventRecorder.Event(&gk, "Normal", "Reconciled", "GCP Instance reconciled")

	return ctrl.Result{}, nil
}

func (cr *GCPInstanceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&benzaiten.GCPInstance{}).
		Complete(cr)
}

func setupGCPInstanceController(mgr manager.Manager, log logr.Logger, cp CloudProviders) error {
	eventRecorder := mgr.GetEventRecorderFor("gcpinstance")
	cc := GCPInstanceReconciler{
		Client:        mgr.GetClient(),
		Scheme:        mgr.GetScheme(),
		eventRecorder: eventRecorder,
		cloud:         cp,
		Log:           log.WithName("GCPInstanceReconciler"),
	}

	// create GCPInstance controller
	err := cc.SetupWithManager(mgr)
	if err != nil {
		return fmt.Errorf("unable to create GCPInstance controller: %w", err)
	}

	return nil
}
