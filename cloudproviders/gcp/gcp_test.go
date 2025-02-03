package gcp

import (
	"context"
	"github.com/golang/mock/gomock"
	compute "google.golang.org/api/compute/v1"
	"testing"
)

const (
	zone = "us-central1-a"
)

func TestCompute_ListInstances(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create the mock.
	mockClient := NewMockServiceInterface(ctrl)

	// Define a fake response.
	fakeInstances := &compute.InstanceList{
		Items: []*compute.Instance{
			{
				Name: "fake-instance-1",
			},
		},
	}

	// Set up the mock to expect a call and return fakeInstances.
	mockClient.
		EXPECT().
		ListInstances(gomock.Any(), zone).
		Return(fakeInstances, nil)

	// Call the mock
	resp, err := mockClient.ListInstances(context.Background(), zone)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(resp.Items) != 1 || resp.Items[0].Name != "fake-instance-1" {
		t.Fatalf("expected 1 instance named fake-instance-1, got %v", resp.Items)
	}
}

func TestCompute_ListNetworks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create the mock.
	mockClient := NewMockServiceInterface(ctrl)

	// Define a fake response.
	fakeNetworks := &compute.NetworkList{
		Items: []*compute.Network{
			{
				Name: "fake-network-1",
			},
		},
	}

	// Set up the mock to expect a call and return fakeNetworks.
	mockClient.
		EXPECT().
		ListNetworks(gomock.Any()).
		Return(fakeNetworks, nil)

	// Call the mock
	resp, err := mockClient.ListNetworks(context.Background())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(resp.Items) != 1 || resp.Items[0].Name != "fake-network-1" {
		t.Fatalf("expected 1 network named fake-network-1, got %v", resp.Items)
	}
}

func TestCompute_GetNetwork(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create the mock.
	mockClient := NewMockServiceInterface(ctrl)

	// Define a fake response.
	fakeNetwork := &compute.Network{
		Name: "fake-network-1",
	}

	// Set up the mock to expect a call and return fakeNetwork.
	mockClient.
		EXPECT().
		GetNetwork(gomock.Any(), "fake-network-id").
		Return(fakeNetwork, nil)

	// Call the mock
	resp, err := mockClient.GetNetwork(context.Background(), "fake-network-id")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if resp.Name != "fake-network-1" {
		t.Fatalf("expected network named fake-network-1, got %v", resp)
	}
}

func TestCompute_CreateNetwork(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create the mock.
	mockClient := NewMockServiceInterface(ctrl)

	// Define a fake response.
	fakeOperation := &compute.Operation{
		Name: "fake-operation-1",
	}

	// Set up the mock to expect a call and return fakeOperation.
	mockClient.
		EXPECT().
		CreateNetwork(gomock.Any(), &compute.Network{Name: "fake-network-1"}).
		Return(fakeOperation, nil)

	// Call the mock
	resp, err := mockClient.CreateNetwork(context.Background(), &compute.Network{Name: "fake-network-1"})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if resp.Name != "fake-operation-1" {
		t.Fatalf("expected operation named fake-operation-1, got %v", resp)
	}
}

func TestCompute_DeleteNetwork(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create the mock.
	mockClient := NewMockServiceInterface(ctrl)

	// Define a fake response.
	fakeOperation := &compute.Operation{
		Name: "fake-operation-1",
	}

	// Set up the mock to expect a call and return fakeOperation.
	mockClient.
		EXPECT().
		DeleteNetwork(gomock.Any(), "fake-network-id").
		Return(fakeOperation, nil)

	// Call the mock
	resp, err := mockClient.DeleteNetwork(context.Background(), "fake-network-id")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if resp.Name != "fake-operation-1" {
		t.Fatalf("expected operation named fake-operation-1, got %v", resp)
	}
}
