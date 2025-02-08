package gcp

import (
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
	expectedInstanceList := &compute.InstanceList{}

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
