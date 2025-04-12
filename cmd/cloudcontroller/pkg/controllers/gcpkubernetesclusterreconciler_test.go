package controllers

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	"github.com/golang/mock/gomock"
	benzaiten "github.com/muraduiurie/cloudcontroller/api/v1"
	"github.com/muraduiurie/cloudcontroller/pkg/cloudproviders/gcp"
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
	defaultZone      = "test-zone"
	defaultProjectID = "test-project"
	defaultGKCName   = "test-gkc"
	defaultNamespace = "default"
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

	//  Create a new manager with the test configmap.
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

func newFakeReconciler(log logr.Logger) (*GCPKubernetesClusterReconciler, error) {
	er := k8sMgr.GetEventRecorderFor("gcpkubernetescluster")
	return &GCPKubernetesClusterReconciler{
		Client:        k8sClient,
		Scheme:        k8sScheme,
		eventRecorder: er,
		Log:           log,
	}, nil
}

func fakeApiGetExistingCluster(ctrl *gomock.Controller) *gcp.API {
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

func fakeApiCreateNewCluster(ctrl *gomock.Controller) *gcp.API {
	mockClustersInterface := gcp.NewMockClustersInterface(ctrl)
	mockGetFailedClustersInterface := gcp.NewMockGetClustersInterface(ctrl)
	mockCreateClustersInterface := gcp.NewMockCreateClustersInterface(ctrl)
	mockGetProvisioningClustersInterface := gcp.NewMockGetClustersInterface(ctrl)
	mockGetRunningClustersInterface := gcp.NewMockGetClustersInterface(ctrl)

	// Verify if cluster exists
	mockClustersInterface.EXPECT().
		Get(defaultProjectID, defaultZone, defaultGKCName).
		Return(mockGetFailedClustersInterface)
	failedCall := mockGetFailedClustersInterface.EXPECT().
		Do().
		Return(nil, fmt.Errorf("googleapi: Error 404: Not found")).Times(1)

	// Create cluster
	expectedOperation := &container.Operation{Name: defaultGKCName}
	mockClustersInterface.EXPECT().
		Create(defaultProjectID, defaultZone, gomock.Any()).
		Return(mockCreateClustersInterface)
	mockCreateClustersInterface.EXPECT().
		Do().
		Return(expectedOperation, nil)

	// Verify provisioning state of cluster
	mockClustersInterface.EXPECT().
		Get(defaultProjectID, defaultZone, defaultGKCName).
		Return(mockGetProvisioningClustersInterface).After(failedCall)
	provisionedCall := mockGetProvisioningClustersInterface.EXPECT().
		Do().
		Return(&container.Cluster{
			Name:             defaultGKCName,
			InitialNodeCount: 1,
			Zone:             defaultZone,
			Status:           string(benzaiten.ClusterStatusProvisioning),
		}, nil).Times(1)

	// Verify running state of cluster
	mockClustersInterface.EXPECT().
		Get(defaultProjectID, defaultZone, defaultGKCName).
		Return(mockGetRunningClustersInterface).After(provisionedCall)
	mockGetRunningClustersInterface.EXPECT().
		Do().
		Return(&container.Cluster{
			Name:             defaultGKCName,
			InitialNodeCount: 1,
			Zone:             defaultZone,
			Status:           string(benzaiten.ClusterStatusRunning),
		}, nil).Times(1)

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

func createFakeGKC(ctx context.Context, fakeClient client.Client, nodeCount int64, name, namespace, zone string) (*benzaiten.GCPKubernetesCluster, error) {
	gkcCreate := benzaiten.GCPKubernetesCluster{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: benzaiten.GCPKubernetesClusterSpec{
			Zone:             zone,
			ClusterName:      name,
			InitialNodeCount: nodeCount,
		},
	}

	err := fakeClient.Create(ctx, &gkcCreate)
	if err != nil {
		return nil, fmt.Errorf("failed to create fake GCPKubernetesCluster: %w", err)
	}

	gk := benzaiten.GCPKubernetesCluster{}
	err = fakeClient.Get(ctx, types.NamespacedName{Name: gkcCreate.Name, Namespace: gkcCreate.Namespace}, &gk)
	if err != nil {
		return nil, fmt.Errorf("failed to get fake GCPKubernetesCluster: %w", err)
	}

	return &gk, nil
}

////////////////////////////////////////////////////
// TESTS
////////////////////////////////////////////////////

func TestGKCReconciler_GetClusterNoChanges(t *testing.T) {
	logger := testLogger()
	logger.WithValues("name", "creation of existing cluster").Info("starting test")

	ctx := context.Background()

	rec, err := newFakeReconciler(logger)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	rec.cloud = CloudProviders{
		GCP: fakeApiGetExistingCluster(mockCtrl),
	}

	gkc, err := createFakeGKC(ctx, rec.Client, 1, defaultGKCName, defaultNamespace, defaultZone)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = rec.Reconcile(ctx, ctrl.Request{NamespacedName: types.NamespacedName{Name: gkc.Name, Namespace: gkc.Namespace}})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = rec.Client.Delete(ctx, gkc)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestGKCReconciler_CreateNewCluster(t *testing.T) {
	logger := testLogger()
	logger.WithValues("name", "creation of a new cluster").Info("starting test")

	ctx := context.Background()

	rec, err := newFakeReconciler(logger)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	rec.cloud = CloudProviders{
		GCP: fakeApiCreateNewCluster(mockCtrl),
	}

	gkc, err := createFakeGKC(ctx, rec.Client, 1, defaultGKCName, defaultNamespace, defaultZone)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = rec.Reconcile(ctx, ctrl.Request{NamespacedName: types.NamespacedName{Name: gkc.Name, Namespace: gkc.Namespace}})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var gkcCreated benzaiten.GCPKubernetesCluster
	err = rec.Get(ctx, types.NamespacedName{Name: gkc.Name, Namespace: gkc.Namespace}, &gkcCreated)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if gkcCreated.Status.Phase != benzaiten.ClusterStatusRunning {
		t.Fatalf("expected cluster status ClusterStatusRunning, got %v", gkcCreated.Status.Phase)
	}

	err = rec.Client.Delete(ctx, gkc)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
