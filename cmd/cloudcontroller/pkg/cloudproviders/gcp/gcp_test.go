package gcp

import (
	"google.golang.org/api/container/v1"
	"google.golang.org/api/googleapi"
	"testing"

	"github.com/golang/mock/gomock"
	compute "google.golang.org/api/compute/v1"
)

func TestListInstances(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mocks
	mockInstancesInterface := NewMockInstancesInterface(ctrl)
	mockListInstancesInterface := NewMockListInstancesInterface(ctrl)

	// Set up expectations
	projectID := "test-project"
	zone := "us-central1-a"
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

func TestListInstances_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mocks
	mockInstancesInterface := NewMockInstancesInterface(ctrl)
	mockListInstancesInterface := NewMockListInstancesInterface(ctrl)

	// Set up expectations
	projectID := "test-project"
	zone := "us-central1-a"
	expectedError := &googleapi.Error{Message: "test error"}

	// Expect the List method to be called with the correct parameters and return the mock ListInstancesInterface
	mockInstancesInterface.EXPECT().
		List(projectID, zone).
		Return(mockListInstancesInterface)

	// Expect the Do method to be called and return an error
	mockListInstancesInterface.EXPECT().
		Do().
		Return(nil, expectedError)

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
	if err == nil {
		t.Fatal("Expected an error, but got nil")
	}

	if instanceList != nil {
		t.Errorf("Expected instance list to be nil, got %v", instanceList)
	}

	if err != expectedError {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}
}

func TestListClusters(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mocks
	mockClustersInterface := NewMockClustersInterface(ctrl)
	mockListClustersInterface := NewMockListClustersInterface(ctrl)

	// Set up expectations
	projectID := "test-project"
	zone := "us-central1-a"
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

//func TestListInstances_Error(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	// Create mocks
//	mockInstancesInterface := NewMockInstancesInterface(ctrl)
//	mockListInstancesInterface := NewMockListInstancesInterface(ctrl)
//
//	// Set up expectations
//	projectID := "test-project"
//	zone := "us-central1-a"
//	expectedError := &googleapi.Error{Message: "test error"}
//
//	// Expect the List method to be called with the correct parameters and return the mock ListInstancesInterface
//	mockInstancesInterface.EXPECT().
//		List(projectID, zone).
//		Return(mockListInstancesInterface)
//
//	// Expect the Do method to be called and return an error
//	mockListInstancesInterface.EXPECT().
//		Do().
//		Return(nil, expectedError)
//
//	// Create the API instance with the mock
//	api := &API{
//		Compute: ComputeService{
//			Clients: ComputeClients{
//				Instances: mockInstancesInterface,
//			},
//		},
//		Config: Config{
//			ProjectId: projectID,
//		},
//	}
//
//	// Call the function under test
//	instanceList, err := api.ListInstances(zone)
//
//	// Verify the results
//	if err == nil {
//		t.Fatal("Expected an error, but got nil")
//	}
//
//	if instanceList != nil {
//		t.Errorf("Expected instance list to be nil, got %v", instanceList)
//	}
//
//	if err != expectedError {
//		t.Errorf("Expected error %v, got %v", expectedError, err)
//	}
//}
