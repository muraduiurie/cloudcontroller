package controllers

import (
	"context"
	"fmt"
	benzaiten "github.com/charmelionag/cloudcontroller/api/v1"
	"github.com/charmelionag/cloudcontroller/pkg/cloudproviders/gcp"
	"github.com/golang/mock/gomock"
	"google.golang.org/api/container/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
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
	k8sScheme *runtime.Scheme
	testEnv   *envtest.Environment
	k8sMgr    manager.Manager
)

const (
	defaultZone      = "us-central1-a"
	defaultProjectID = "test-project"
	defaultGKCName   = "test-gkc"
	defaultNamespace = "default"
	//networkID        = "test-network"
)

func TestMain(m *testing.M) {
	// Set up logger.
	ctrl.SetLogger(zap.New(zap.UseDevMode(true)))

	// Start the envtest environment.
	testEnv = &envtest.Environment{
		CRDDirectoryPaths: []string{
			filepath.Join("..", "..", "crds"),
		},
		BinaryAssetsDirectory: "../../../../envtest/k8s/1.31.0-linux-amd64",
	}

	var err error
	cfg, err = testEnv.Start()
	if err != nil {
		log.Fatalf("Failed to start test environment: %v", err)
	}

	// Create a new manager with the test config.
	mgr, err := ctrl.NewManager(cfg, ctrl.Options{
		Scheme: Scheme,
	})
	if err != nil {
		log.Fatalf("Failed to create manager: %v", err)
	}

	// Use a cancelable context for graceful shutdown.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start the manager in a separate goroutine.
	go func() {
		if err := mgr.Start(ctx); err != nil {
			log.Fatalf("Manager failed to start: %v", err)
		}
	}()

	// Obtain a client from the manager.
	k8sClient = mgr.GetClient()
	k8sScheme = mgr.GetScheme()
	k8sMgr = mgr

	// Run tests.
	code := m.Run()

	// Signal shutdown and clean up.
	cancel() // Gracefully stop the manager.

	if err := testEnv.Stop(); err != nil {
		log.Printf("Failed to stop test environment: %v", err)
	}

	os.Exit(code)
}

func newFakeGKCReconciler() (*GCPKubernetesClusterReconciler, error) {
	er := k8sMgr.GetEventRecorderFor("gcpkubernetescluster")
	return &GCPKubernetesClusterReconciler{
		Client:        k8sClient,
		Scheme:        k8sScheme,
		eventRecorder: er,
	}, nil
}

func createFakeGCPApiGetDefaultClusterClient(ctrl *gomock.Controller) *gcp.API {
	// Create mocks
	mockClustersInterface := gcp.NewMockClustersInterface(ctrl)
	mockGetClustersInterface := gcp.NewMockGetClustersInterface(ctrl)

	// Set up expectations
	expectedCluster := &container.Cluster{
		Name:             defaultGKCName,
		InitialNodeCount: 1,
		Zone:             defaultZone,
	}

	// Expect the Get method to be called with the correct parameters and return the mock GetClustersInterface
	mockClustersInterface.EXPECT().
		Get(defaultProjectID, defaultZone, defaultGKCName).
		Return(mockGetClustersInterface)

	// Expect the Do method to be called and return the expected cluster
	mockGetClustersInterface.EXPECT().
		Do().
		Return(expectedCluster, nil)

	// Create the API cluster with the mock
	api := &gcp.API{
		Container: gcp.ContainerService{
			Clients: gcp.ContainerClients{
				Clusters: mockClustersInterface,
			},
		},
		Config: gcp.Config{
			ProjectId: defaultProjectID,
		},
	}

	return api
}

func createFakeGKC(fakeClient client.Client) (*benzaiten.GCPKubernetesCluster, error) {
	gkcCreate := benzaiten.GCPKubernetesCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      defaultGKCName,
			Namespace: defaultNamespace,
		},
		Spec: benzaiten.GCPKubernetesClusterSpec{
			Zone:             defaultZone,
			Name:             defaultGKCName,
			InitialNodeCount: 3,
			Autopilot:        true,
		},
	}

	err := fakeClient.Create(context.TODO(), &gkcCreate)
	if err != nil {
		return nil, fmt.Errorf("failed to create fake GCPKubernetesCluster: %w", err)
	}

	gk := benzaiten.GCPKubernetesCluster{}
	err = fakeClient.Get(context.TODO(), types.NamespacedName{Name: gkcCreate.Name, Namespace: gkcCreate.Namespace}, &gk)
	if err != nil {
		return nil, fmt.Errorf("failed to get fake GCPKubernetesCluster: %w", err)
	}

	return &gk, nil
}

func TestGKCReconciler_GetClusterNoChanges(t *testing.T) {
	rec, err := newFakeGKCReconciler()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	rec.cloud = CloudProviders{
		GCP: createFakeGCPApiGetDefaultClusterClient(mockCtrl),
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
