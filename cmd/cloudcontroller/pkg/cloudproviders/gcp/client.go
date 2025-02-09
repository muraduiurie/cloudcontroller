package gcp

import (
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/container/v1"
	"google.golang.org/api/googleapi"
)

//===============================================================================================
// TYPES AND INTERFACES
//===============================================================================================

// Services
type (
	ContainerService struct {
		Clients ContainerClients
	}
	ComputeService struct {
		Clients ComputeClients
	}
)

// Clients
type (
	ComputeClients struct {
		Instances InstancesInterface
		Networks  NetworksInterface
	}
	ContainerClients struct {
		Clusters ClustersInterface
	}
)

// Resources
type (
	// compute resources
	GCPInstances struct {
		InstancesService *compute.InstancesService
	}
	GCPNetworks struct {
		NetworksService *compute.NetworksService
	}

	// container resources
	GCPKubernetesClusters struct {
		ClustersService *container.ProjectsZonesClustersService
	}
)

// Interfaces
type (
	// compute interfaces
	//// instances
	InstancesInterface interface {
		List(project, zone string) ListInstancesInterface
	}
	//// networks
	NetworksInterface interface {
		List(project string) ListNetworksInterface
		Get(project, network string) GetNetworksInterface
		Insert(project string, network *compute.Network) CreateNetworksInterface
		Delete(project, network string) DeleteNetworksInterface
	}

	// container interfaces
	//// kubernetes clusters
	ClustersInterface interface {
		List(project, zone string) ListClustersInterface
	}
)

// Requests
type (
	// compute do interfaces
	//// instances
	ListInstancesInterface interface {
		Do(opts ...googleapi.CallOption) (*compute.InstanceList, error)
	}
	//// networks
	ListNetworksInterface interface {
		Do(opts ...googleapi.CallOption) (*compute.NetworkList, error)
	}
	GetNetworksInterface interface {
		Do(opts ...googleapi.CallOption) (*compute.Network, error)
	}
	CreateNetworksInterface interface {
		Do(opts ...googleapi.CallOption) (*compute.Operation, error)
	}
	DeleteNetworksInterface interface {
		Do(opts ...googleapi.CallOption) (*compute.Operation, error)
	}

	// container interfaces
	//// kubernetes clusters
	ListClustersInterface interface {
		Do(opts ...googleapi.CallOption) (*container.ListClustersResponse, error)
	}
)

// Executor requests
type (
	// compute google calls
	//// instances
	ListInstancesRequest struct {
		googleCall *compute.InstancesListCall
	}
	//// networks
	ListNetworksRequest struct {
		googleCall *compute.NetworksListCall
	}
	GetNetworksRequest struct {
		googleCall *compute.NetworksGetCall
	}
	CreateNetworksRequest struct {
		googleCall *compute.NetworksInsertCall
	}
	DeleteNetworksRequest struct {
		googleCall *compute.NetworksDeleteCall
	}

	// container google calls
	//// kubernetes clusters
	ListClustersRequest struct {
		googleCall *container.ProjectsZonesClustersListCall
	}
)

// ===============================================================================================
// FUNCTIONS
// ===============================================================================================
// Verbs
// // Compute
// //// Instances
func (i *GCPInstances) List(projectID, zone string) ListInstancesInterface {
	return &ListInstancesRequest{
		googleCall: i.InstancesService.List(projectID, zone),
	}
}

// //// Networks
func (n *GCPNetworks) List(projectID string) ListNetworksInterface {
	return &ListNetworksRequest{
		googleCall: n.NetworksService.List(projectID),
	}
}
func (n *GCPNetworks) Get(projectID, network string) GetNetworksInterface {
	return &GetNetworksRequest{
		googleCall: n.NetworksService.Get(projectID, network),
	}
}
func (n *GCPNetworks) Insert(projectID string, network *compute.Network) CreateNetworksInterface {
	return &CreateNetworksRequest{
		googleCall: n.NetworksService.Insert(projectID, network),
	}
}
func (n *GCPNetworks) Delete(projectID, network string) DeleteNetworksInterface {
	return &DeleteNetworksRequest{
		googleCall: n.NetworksService.Delete(projectID, network),
	}
}

// // Container
// ///// Clusters
func (g *GCPKubernetesClusters) List(projectID, zone string) ListClustersInterface {
	return &ListClustersRequest{
		googleCall: g.ClustersService.List(projectID, zone),
	}
}

// Execs
// // Compute
// //// Instances
func (lc *ListInstancesRequest) Do(opts ...googleapi.CallOption) (*compute.InstanceList, error) {
	return lc.googleCall.Do(opts...)
}

// //// Networks
func (lc *ListNetworksRequest) Do(opts ...googleapi.CallOption) (*compute.NetworkList, error) {
	return lc.googleCall.Do(opts...)
}
func (lc *GetNetworksRequest) Do(opts ...googleapi.CallOption) (*compute.Network, error) {
	return lc.googleCall.Do(opts...)
}
func (lc *CreateNetworksRequest) Do(opts ...googleapi.CallOption) (*compute.Operation, error) {
	return lc.googleCall.Do(opts...)
}
func (lc *DeleteNetworksRequest) Do(opts ...googleapi.CallOption) (*compute.Operation, error) {
	return lc.googleCall.Do(opts...)
}

// // Container
// //// Clusters
func (lc *ListClustersRequest) Do(opts ...googleapi.CallOption) (*container.ListClustersResponse, error) {
	return lc.googleCall.Do(opts...)
}
