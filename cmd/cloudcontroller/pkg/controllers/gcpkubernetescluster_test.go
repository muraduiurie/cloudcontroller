package controllers

import (
	"context"
	"fmt"
	benzaiten "github.com/charmelionag/cloudcontroller/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"log"
	"os"
	"path/filepath"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"testing"
)

var (
	cfg       *rest.Config
	k8sClient client.Client
	testEnv   *envtest.Environment
	fakeMgr   manager.Manager
)

func TestMain(m *testing.M) {
	// Set up logger.
	ctrl.SetLogger(zap.New(zap.UseDevMode(true)))

	// Start the envtest environment.
	testEnv = &envtest.Environment{
		CRDDirectoryPaths: []string{
			filepath.Join("..", "..", "crds"),
		},
		BinaryAssetsDirectory: "/home/imuradu/.local/share/kubebuilder-envtest/k8s/1.31.0-linux-amd64",
	}

	var err error
	cfg, err = testEnv.Start()
	if err != nil {
		log.Fatal(err)
	}

	// Create a new manager with the test config.
	mgr, err := ctrl.NewManager(cfg, ctrl.Options{
		Scheme: Scheme,
	})
	if err != nil {
		log.Fatal(err)
	}
	fakeMgr = mgr

	// Start the manager in a separate goroutine.
	go func() {
		if err := mgr.Start(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	// Obtain a client from the manager.
	k8sClient = mgr.GetClient()

	// Run tests.
	code := m.Run()

	// Teardown the test environment.
	if err := testEnv.Stop(); err != nil {
		log.Fatal(err)
	}

	os.Exit(code)
}

func newFakeGKCReconciler() (*GCPKubernetesClusterReconciler, error) {
	er := fakeMgr.GetEventRecorderFor("gcpkubernetescluster")
	return &GCPKubernetesClusterReconciler{
		Client:        k8sClient,
		Scheme:        Scheme,
		eventRecorder: er,
	}, nil
}

func createFakeGKC(fakeClient client.Client) (*benzaiten.GCPKubernetesCluster, error) {
	gkc := benzaiten.GCPKubernetesCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-gkc",
			Namespace: "default",
		},
		Spec: benzaiten.GCPKubernetesClusterSpec{
			Zone:             "us-central1-a",
			Name:             "test-gke",
			InitialNodeCount: 1,
		},
	}

	err := fakeClient.Create(context.TODO(), &gkc)
	if err != nil {
		return nil, fmt.Errorf("failed to create fake GCPKubernetesCluster: %w", err)
	}

	return &gkc, nil
}

func TestGKCReconciler(t *testing.T) {
	rec, err := newFakeGKCReconciler()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	gkc, err := createFakeGKC(rec.Client)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = rec.Reconcile(context.TODO(), ctrl.Request{NamespacedName: types.NamespacedName{Name: gkc.Name, Namespace: gkc.Namespace}})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
