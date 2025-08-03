package gcp

import (
	"google.golang.org/api/container/v1"
	"testing"

	"github.com/golang/mock/gomock"
	compute "google.golang.org/api/compute/v1"
)

const (
	zone      = "us-central1-a"
	projectID = "test-project"
	networkID = "test-network"
)

func TestListInstances(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mocks
	mockInstancesInterface := NewMockInstancesInterface(ctrl)
	mockListInstancesInterface := NewMockListInstancesInterface(ctrl)

	// Set up expectations
	expectedInstanceList := &compute.InstanceList{
		Items: []*compute.Instance{
			{
				Name: "test-instance",
			},
		},
	}

	// Expect the List method to be called with the correct parameters and return the mock ListInstancesInterface
	mockInstancesInterface.EXPECT().
		List(projectID, zone).
		Return(mockListInstancesInterface)

	// Expect the Do method to be called and return the expected instance list
	mockListInstancesInterface.EXPECT().
		Do().
		Return(expectedInstanceList, nil)

	// Create the API instance with the mock
	api := &API{
		Compute: ComputeService{
			Clients: ComputeClients{
				Instances: mockInstancesInterface,
			},
		},
		Config: Config{
			ProjectId: projectID,
		},
	}

	// Call the function under test
	instanceList, err := api.ListInstances(zone)

	// Verify the results
	if err != nil {
		t.Fatalf("ListInstances returned an error: %v", err)
	}

	if instanceList != expectedInstanceList {
		t.Errorf("Expected instance list %v, got %v", expectedInstanceList, instanceList)
	}
}

func TestListClusters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mocks
	mockClustersInterface := NewMockClustersInterface(ctrl)
	mockListClustersInterface := NewMockListClustersInterface(ctrl)

	// Set up expectations
	expectedClustersList := &container.ListClustersResponse{
		Clusters: []*container.Cluster{
			{
				Name:             "test-cluster",
				InitialNodeCount: 1,
			},
		},
	}

	// Expect the List method to be called with the correct parameters and return the mock ListClustersInterface
	mockClustersInterface.EXPECT().
		List(projectID, zone).
		Return(mockListClustersInterface)

	// Expect the Do method to be called and return the expected cluster list
	mockListClustersInterface.EXPECT().
		Do().
		Return(expectedClustersList, nil)

	// Create the API cluster with the mock
	api := &API{
		Container: ContainerService{
			Clients: ContainerClients{
				Clusters: mockClustersInterface,
			},
		},
		Config: Config{
			ProjectId: projectID,
		},
	}

	// Call the function under test
	clusterList, err := api.ListClusters(zone)
	// Verify the results
	if err != nil {
		t.Fatalf("ListClusters returned an error: %v", err)
	}

	if clusterList != expectedClustersList {
		t.Errorf("Expected cluster list %v, got %v", expectedClustersList, clusterList)
	}
}

func TestListNetworks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mocks
	mockNetworksInterface := NewMockNetworksInterface(ctrl)
	mockListNetworksInterface := NewMockListNetworksInterface(ctrl)

	// Set up expectations
	expectedNetworkList := &compute.NetworkList{
		Items: []*compute.Network{
			{
				Name: "test-network",
			},
		},
	}

	// Expect the List method to be called with the correct parameters and return the mock ListNetworksInterface
	mockNetworksInterface.EXPECT().
		List(projectID).
		Return(mockListNetworksInterface)

	// Expect the Do method to be called and return the expected network list
	mockListNetworksInterface.EXPECT().
		Do().
		Return(expectedNetworkList, nil)

	// Create the API network with the mock
	api := &API{
		Compute: ComputeService{
			Clients: ComputeClients{
				Networks: mockNetworksInterface,
			},
		},
		Config: Config{
			ProjectId: projectID,
		},
	}

	// Call the function under test
	networkList, err := api.ListNetworks()

	// Verify the results
	if err != nil {
		t.Fatalf("ListNetworks returned an error: %v", err)
	}

	if networkList != expectedNetworkList {
		t.Errorf("Expected network list %v, got %v", expectedNetworkList, networkList)
	}
}

func TestGetNetwork(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mocks
	mockNetworksInterface := NewMockNetworksInterface(ctrl)
	mockGetNetworksInterface := NewMockGetNetworksInterface(ctrl)

	// Set up expectations
	expectedNetwork := &compute.Network{
		Name:                  "test-network",
		AutoCreateSubnetworks: true,
	}

	// Expect the Get method to be called with the correct parameters and return the mock GetNetworksInterface
	mockNetworksInterface.EXPECT().
		Get(projectID, networkID).
		Return(mockGetNetworksInterface)

	// Expect the Do method to be called and return the expected network
	mockGetNetworksInterface.EXPECT().
		Do().
		Return(expectedNetwork, nil)

	// Create the API network with the mock
	api := &API{
		Compute: ComputeService{
			Clients: ComputeClients{
				Networks: mockNetworksInterface,
			},
		},
		Config: Config{
			ProjectId: projectID,
		},
	}

	// Call the function under test
	network, err := api.GetNetwork(networkID)

	// Verify the results
	if err != nil {
		t.Fatalf("GetNetwork returned an error: %v", err)
	}

	if network != expectedNetwork {
		t.Errorf("Expected network %v, got %v", expectedNetwork, network)
	}
}

func TestCreateNetwork(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mocks
	mockNetworksInterface := NewMockNetworksInterface(ctrl)
	mockCreateNetworksInterface := NewMockCreateNetworksInterface(ctrl)

	// Set up expectations
	expectedOperation := &compute.Operation{
		Name: "test-operation",
	}

	// Expect the Insert method to be called with the correct parameters and return the mock CreateNetworksInterface
	mockNetworksInterface.EXPECT().
		Insert(projectID, gomock.Any()).
		Return(mockCreateNetworksInterface)

	// Expect the Do method to be called and return the expected operation
	mockCreateNetworksInterface.EXPECT().
		Do().
		Return(expectedOperation, nil)

	// Create the API network with the mock
	api := &API{
		Compute: ComputeService{
			Clients: ComputeClients{
				Networks: mockNetworksInterface,
			},
		},
		Config: Config{
			ProjectId: projectID,
		},
	}

	// Call the function under test
	operation, err := api.CreateNetwork(&compute.Network{
		Name:                  "test-network",
		AutoCreateSubnetworks: true,
	})

	// Verify the results
	if err != nil {
		t.Fatalf("CreateNetwork returned an error: %v", err)
	}

	if operation != expectedOperation {
		t.Errorf("Expected operation %v, got %v", expectedOperation, operation)
	}
}

func TestDeleteNetwork(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mocks
	mockNetworksInterface := NewMockNetworksInterface(ctrl)
	mockDeleteNetworksInterface := NewMockDeleteNetworksInterface(ctrl)

	// Set up expectations
	expectedOperation := &compute.Operation{
		Name: "test-operation",
	}

	// Expect the Delete method to be called with the correct parameters and return the mock DeleteNetworksInterface
	mockNetworksInterface.EXPECT().
		Delete(projectID, networkID).
		Return(mockDeleteNetworksInterface)

	// Expect the Do method to be called and return the expected operation
	mockDeleteNetworksInterface.EXPECT().
		Do().
		Return(expectedOperation, nil)

	// Create the API network with the mock
	api := &API{
		Compute: ComputeService{
			Clients: ComputeClients{
				Networks: mockNetworksInterface,
			},
		},
		Config: Config{
			ProjectId: projectID,
		},
	}

	// Call the function under test
	operation, err := api.DeleteNetwork(networkID)

	// Verify the results
	if err != nil {
		t.Fatalf("DeleteNetwork returned an error: %v", err)
	}

	if operation != expectedOperation {
		t.Errorf("Expected operation %v, got %v", expectedOperation, operation)
	}
}

func TestCreateCluster(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mocks
	mockClustersInterface := NewMockClustersInterface(ctrl)
	mockCreateClustersInterface := NewMockCreateClustersInterface(ctrl)

	// Set up expectations
	expectedOperation := &container.Operation{
		Name: "test-operation",
	}

	// Expect the Create method to be called with the correct parameters and return the mock CreateClustersInterface
	mockClustersInterface.EXPECT().
		Create(projectID, zone, gomock.Any()).
		Return(mockCreateClustersInterface)

	// Expect the Do method to be called and return the expected operation
	mockCreateClustersInterface.EXPECT().
		Do().
		Return(expectedOperation, nil)

	// Create the API cluster with the mock
	api := &API{
		Container: ContainerService{
			Clients: ContainerClients{
				Clusters: mockClustersInterface,
			},
		},
		Config: Config{
			ProjectId: projectID,
		},
	}

	// Call the function under test
	operation, err := api.CreateCluster(zone, &container.Cluster{
		Name:             "test-cluster",
		InitialNodeCount: 1,
	})

	// Verify the results
	if err != nil {
		t.Fatalf("CreateCluster returned an error: %v", err)
	}

	if operation != expectedOperation {
		t.Errorf("Expected operation %v, got %v", expectedOperation, operation)
	}
}

func TestDeleteCluster(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mocks
	mockClustersInterface := NewMockClustersInterface(ctrl)
	mockDeleteClustersInterface := NewMockDeleteClustersInterface(ctrl)

	// Set up expectations
	expectedOperation := &container.Operation{
		Name: "test-operation",
	}

	// Expect the Delete method to be called with the correct parameters and return the mock DeleteClustersInterface
	mockClustersInterface.EXPECT().
		Delete(projectID, zone, "test-cluster").
		Return(mockDeleteClustersInterface)

	// Expect the Do method to be called and return the expected operation
	mockDeleteClustersInterface.EXPECT().
		Do().
		Return(expectedOperation, nil)

	// Create the API cluster with the mock
	api := &API{
		Container: ContainerService{
			Clients: ContainerClients{
				Clusters: mockClustersInterface,
			},
		},
		Config: Config{
			ProjectId: projectID,
		},
	}

	// Call the function under test
	operation, err := api.DeleteCluster(zone, "test-cluster")

	// Verify the results
	if err != nil {
		t.Fatalf("DeleteCluster returned an error: %v", err)
	}

	if operation != expectedOperation {
		t.Errorf("Expected operation %v, got %v", expectedOperation, operation)
	}
}

func TestGetCluster(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mocks
	mockClustersInterface := NewMockClustersInterface(ctrl)
	mockGetClustersInterface := NewMockGetClustersInterface(ctrl)

	// Set up expectations
	expectedCluster := &container.Cluster{
		Name:             "test-cluster",
		InitialNodeCount: 1,
	}

	// Expect the Get method to be called with the correct parameters and return the mock GetClustersInterface
	mockClustersInterface.EXPECT().
		Get(projectID, zone, "test-cluster").
		Return(mockGetClustersInterface)

	// Expect the Do method to be called and return the expected cluster
	mockGetClustersInterface.EXPECT().
		Do().
		Return(expectedCluster, nil)

	// Create the API cluster with the mock
	api := &API{
		Container: ContainerService{
			Clients: ContainerClients{
				Clusters: mockClustersInterface,
			},
		},
		Config: Config{
			ProjectId: projectID,
		},
	}

	// Call the function under test
	cluster, err := api.GetCluster(zone, "test-cluster")

	// Verify the results
	if err != nil {
		t.Fatalf("GetCluster returned an error: %v", err)
	}

	if cluster != expectedCluster {
		t.Errorf("Expected cluster %v, got %v", expectedCluster, cluster)
	}
}
