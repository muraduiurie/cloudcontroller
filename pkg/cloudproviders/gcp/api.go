package gcp

import (
	"context"
	"github.com/go-logr/logr"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/container/v1"
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

type API struct {
	Compute   ComputeService
	Container ContainerService
	Config
}

type ClusterUpdates struct {
	DesiredNodeCount int64 `json:"desiredNodeCount"`
}

func NewAPI(ctx context.Context, log logr.Logger, gcpSaFilePath string) (*API, error) {
	log.Info("Creating GCP API client")

	config, err := getConfig(gcpSaFilePath)
	if err != nil {
		return nil, err
	}

	computeService, err := compute.NewService(ctx, option.WithCredentialsFile(gcpSaFilePath))
	if err != nil {
		return nil, err
	}

	containerService, err := container.NewService(ctx, option.WithCredentialsFile(gcpSaFilePath))
	if err != nil {
		return nil, err
	}

	return &API{
		Compute: ComputeService{
			Clients: ComputeClients{
				Instances: &GCPInstances{
					InstancesService: computeService.Instances,
				},
				Networks: &GCPNetworks{
					NetworksService: computeService.Networks,
				},
			},
		},
		Container: ContainerService{
			Clients: ContainerClients{
				Clusters: &GCPKubernetesClusters{
					ClustersService: containerService.Projects.Zones.Clusters,
				},
			},
		},
		Config: config,
	}, nil
}

func (a *API) ListInstances(zone string) (*compute.InstanceList, error) {
	resp, err := a.Compute.Clients.Instances.List(a.ProjectId, zone).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *API) ListNetworks() (*compute.NetworkList, error) {
	resp, err := a.Compute.Clients.Networks.List(a.ProjectId).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *API) GetNetwork(nid string) (*compute.Network, error) {
	resp, err := a.Compute.Clients.Networks.Get(a.ProjectId, nid).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *API) CreateNetwork(network *compute.Network) (*compute.Operation, error) {
	resp, err := a.Compute.Clients.Networks.Insert(a.ProjectId, network).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *API) DeleteNetwork(nid string) (*compute.Operation, error) {
	resp, err := a.Compute.Clients.Networks.Delete(a.ProjectId, nid).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *API) ListClusters(zone string) (*container.ListClustersResponse, error) {
	resp, err := a.Container.Clients.Clusters.List(a.ProjectId, zone).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *API) GetCluster(zone, clusterName string) (*container.Cluster, error) {
	resp, err := a.Container.Clients.Clusters.Get(a.ProjectId, zone, clusterName).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *API) CreateCluster(zone string, cluster *container.Cluster) (*container.Operation, error) {
	resp, err := a.Container.Clients.Clusters.Create(a.ProjectId, zone, &container.CreateClusterRequest{
		Cluster:   cluster,
		Zone:      zone,
		ProjectId: a.ProjectId,
	}).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *API) DeleteCluster(zone, clusterName string) (*container.Operation, error) {
	resp, err := a.Container.Clients.Clusters.Delete(a.ProjectId, zone, clusterName).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *API) UpdateCluster(zone, clusterName string, cu *ClusterUpdates) (*container.Operation, error) {
	updateRequest := container.UpdateClusterRequest{
		ClusterId: clusterName,
		Zone:      zone,
		ProjectId: a.ProjectId,
		Update:    &container.ClusterUpdate{},
	}
	resp, err := a.Container.Clients.Clusters.Update(a.ProjectId, zone, clusterName, &updateRequest).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}
