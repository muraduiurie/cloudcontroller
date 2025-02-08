package gcp

import (
	"context"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
)

type Config struct {
	Type                    string `json:"type"`
	ProjectId               string `json:"project_id"`
	PrivateKeyId            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientId                string `json:"client_id"`
	AuthUri                 string `json:"auth_uri"`
	TokenUri                string `json:"token_uri"`
	AuthProviderX509CertUrl string `json:"auth_provider_x509_cert_url"`
	ClientX509CertUrl       string `json:"client_x509_cert_url"`
	UniverseDomain          string `json:"universe_domain"`
}

//type ContainerService struct {
//	Clients
//}

type API struct {
	Compute ComputeService
	//Container ContainerService
	Config
}

type ComputeService struct {
	Clients ComputeClients
}

type ComputeClients struct {
	Instances InstancesInterface
}

func NewAPI(ctx context.Context, gcpSaFilePath string) (*API, error) {
	config, err := getConfig(gcpSaFilePath)
	if err != nil {
		return nil, err
	}

	computeService, err := compute.NewService(ctx, option.WithCredentialsFile(gcpSaFilePath))
	if err != nil {
		return nil, err
	}

	//containerService, err := container.NewService(ctx, option.WithCredentialsFile(gcpSaFilePath))
	//if err != nil {
	//	return nil, err
	//}

	return &API{
		Compute: ComputeService{
			Clients: ComputeClients{
				Instances: &GCPInstances{
					InstancesService: computeService.Instances,
				},
			},
		},
		//Container: &Container{Client: containerService},
		Config: config,
	}, nil
}

// First-level instances interface
type InstancesInterface interface {
	List(project, zone string) ListInstancesInterface
}

// Instances struct implements InstancesInterface
// Wrap the underlying compute service.
type GCPInstances struct {
	InstancesService *compute.InstancesService
}

func (i *GCPInstances) List(projectID, zone string) ListInstancesInterface {
	return &ListInstancesRequest{
		googleCall: i.InstancesService.List(projectID, zone),
	}
}

type ListInstancesInterface interface {
	Do(opts ...googleapi.CallOption) (*compute.InstanceList, error)
}

type ListInstancesRequest struct {
	googleCall *compute.InstancesListCall
}

func (lc *ListInstancesRequest) Do(opts ...googleapi.CallOption) (*compute.InstanceList, error) {
	return lc.googleCall.Do(opts...)
}

func (a *API) ListInstances(zone string) (*compute.InstanceList, error) {
	resp, err := a.Compute.Clients.Instances.List(a.ProjectId, zone).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//func (a *API) ListNetworks(ctx context.Context) (*compute.NetworkList, error) {
//	resp, err := a.Compute.Client.Networks.List(a.ProjectId).Context(ctx).Do()
//	if err != nil {
//		return nil, err
//	}
//	return resp, nil
//}
//
//func (a *API) GetNetwork(ctx context.Context, nid string) (*compute.Network, error) {
//	resp, err := a.Compute.Client.Networks.Get(a.ProjectId, nid).Context(ctx).Do()
//	if err != nil {
//		return nil, err
//	}
//	return resp, nil
//}
//
//func (a *API) CreateNetwork(ctx context.Context, network *compute.Network) (*compute.Operation, error) {
//	resp, err := a.Compute.Client.Networks.Insert(a.ProjectId, network).Context(ctx).Do()
//	if err != nil {
//		return nil, err
//	}
//	return resp, nil
//}
//
//func (a *API) DeleteNetwork(ctx context.Context, nid string) (*compute.Operation, error) {
//	resp, err := a.Compute.Client.Networks.Delete(a.ProjectId, nid).Context(ctx).Do()
//	if err != nil {
//		return nil, err
//	}
//	return resp, nil
//}
//
//func (a *API) ListClusters(ctx context.Context, zone string) (*container.ListClustersResponse, error) {
//	resp, err := a.Container.Client.Projects.Zones.Clusters.List(a.ProjectId, zone).Context(ctx).Do()
//	if err != nil {
//		return nil, err
//	}
//	return resp, nil
//}
//
//func (a *API) GetCluster(ctx context.Context, zone, clusterName string) (*container.Cluster, error) {
//	resp, err := a.Container.Client.Projects.Zones.Clusters.Get(a.ProjectId, zone, clusterName).Context(ctx).Do()
//	if err != nil {
//		return nil, err
//	}
//	return resp, nil
//}
//
//func (a *API) CreateCluster(ctx context.Context, zone string, cluster *container.Cluster) (*container.Operation, error) {
//	resp, err := a.Container.Client.Projects.Zones.Clusters.Create(a.ProjectId, zone, &container.CreateClusterRequest{
//		Cluster:   cluster,
//		Zone:      zone,
//		ProjectId: a.ProjectId,
//	}).Context(ctx).Do()
//	if err != nil {
//		return nil, err
//	}
//	return resp, nil
//}
//
//func (a *API) DeleteCluster(ctx context.Context, zone, clusterName string) (*container.Operation, error) {
//	resp, err := a.Container.Client.Projects.Zones.Clusters.Delete(a.ProjectId, zone, clusterName).Context(ctx).Do()
//	if err != nil {
//		return nil, err
//	}
//	return resp, nil
//}
