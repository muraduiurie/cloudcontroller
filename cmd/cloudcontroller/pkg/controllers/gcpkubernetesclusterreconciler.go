package controllers

import (
	"context"
	"fmt"
	benzaiten "github.com/charmelionag/cloudcontroller/api/v1"
	"google.golang.org/api/container/v1"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"strings"
	"time"
)

type GCPKubernetesClusterReconciler struct {
	client.Client
	Scheme        *runtime.Scheme
	eventRecorder record.EventRecorder
	cloud         CloudProviders
}

func (cr *GCPKubernetesClusterReconciler) updateStatus(ctx context.Context, cluster *benzaiten.GCPKubernetesCluster, cs benzaiten.ClusterStatus, msg, rsn, et string) error {
	cr.eventRecorder.Event(cluster, et, rsn, msg)
	cluster.Status = benzaiten.GCPKubernetesClusterStatus{
		Phase: cs,
	}

	err := cr.Status().Update(ctx, cluster)
	if err != nil {
		return err
	}

	return nil
}

func (cr *GCPKubernetesClusterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx).WithValues("gcpkubernetescluster", req.NamespacedName)
	gkcCR := benzaiten.GCPKubernetesCluster{}
	err := cr.Get(ctx, req.NamespacedName, &gkcCR)
	if err != nil {
		if kerr.IsNotFound(err) {
			logger.Info("gcpkubernetescluster not found")
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}
	// does cluster exist in GCP?
	gkc, err := cr.cloud.GCP.GetCluster(gkcCR.Spec.Zone, gkcCR.Spec.ClusterName)
	if err != nil && notFoundGCPResource(err) {
		// cluster does not exist in GCP
		logger.Info("gcpkubernetescluster not found, creating cluster...")
		_, err := cr.cloud.GCP.CreateCluster(gkcCR.Spec.Zone, &container.Cluster{
			Name:             gkcCR.Spec.ClusterName,
			InitialNodeCount: gkcCR.Spec.InitialNodeCount,
		})
		if err != nil {
			logger.Error(err, "error creating gcpkubernetescluster")
			return ctrl.Result{}, err
		}
		// create cluster
		err = cr.updateStatus(ctx, &gkcCR, benzaiten.ClusterStatusProvisioning, "GCP Kubernetes Cluster provisioning", "ClusterProvisioning", "Normal")
		if err != nil {
			logger.Error(err, "error updating gcpkubernetescluster status")
			return ctrl.Result{}, err
		}
		// wait for cluster running state
		ticker := time.NewTicker(3 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				gkcCreated, err := cr.cloud.GCP.GetCluster(gkcCR.Spec.Zone, gkcCR.Spec.ClusterName)
				if err != nil {
					logger.Error(err, "error verifying cluster status", "gcpkubernetescluster", gkcCR)
					return ctrl.Result{}, err
				}

				if gkcCreated.Status == string(benzaiten.ClusterStatusRunning) {
					// cluster is running
					err = cr.updateStatus(ctx, &gkcCR, benzaiten.ClusterStatusRunning, "GCP Kubernetes Cluster running", "ClusterRunning", "Normal")
					if err != nil {
						logger.Error(err, "error updating gcpkubernetescluster status")
						return ctrl.Result{}, err
					}
					return ctrl.Result{}, nil
				} else if gkcCreated.Status == string(benzaiten.ClusterStatusError) || gkcCreated.Status == string(benzaiten.ClusterStatusDegraded) || gkcCreated.Status == string(benzaiten.ClusterStatusUnspecified) {
					// cluster is in error
					err = cr.updateStatus(ctx, &gkcCR, benzaiten.ClusterStatusError, "GCP Kubernetes Cluster in failed state", "ClusterFailedState", "Warning")
					if err != nil {
						logger.Error(err, "error updating gcpkubernetescluster status")
						return ctrl.Result{}, err
					}
					return ctrl.Result{}, nil
				}
			case <-ctx.Done():
				return ctrl.Result{}, ctx.Err()
			}
		}
	} else if err != nil {
		logger.Error(err, "error getting gcpkubernetescluster")
		return ctrl.Result{}, err
	}
	// synchronize changes if exists
	logger.Info("gcpkubernetescluster found, synchronizing...", "gcpkubernetescluster", gkc.Name)

	logger.Info("gcp kubernetes cluster reconciled")
	return ctrl.Result{RequeueAfter: time.Second * 60}, nil
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

func notFoundGCPResource(err error) bool {
	return strings.Split(fmt.Sprintf("%v", err), ":")[1] == " Error 404"
}
