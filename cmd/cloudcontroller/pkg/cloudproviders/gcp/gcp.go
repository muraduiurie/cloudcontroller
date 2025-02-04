package gcp

import (
	"context"
	"encoding/json"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/container/v1"
	"google.golang.org/api/option"
	"log"
	"os"
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

type ServiceInterface interface {
	ListInstances(ctx context.Context, zone string) (*compute.InstanceList, error)
	ListNetworks(ctx context.Context) (*compute.NetworkList, error)
	GetNetwork(ctx context.Context, nid string) (*compute.Network, error)
	CreateNetwork(ctx context.Context, network *compute.Network) (*compute.Operation, error)
	DeleteNetwork(ctx context.Context, nid string) (*compute.Operation, error)
	ListClusters(ctx context.Context, zone string) (*container.ListClustersResponse, error)
	CreateCluster(ctx context.Context, zone string, cluster *container.Cluster) (*container.Operation, error)
	DeleteCluster(ctx context.Context, zone, clusterName string) (*container.Operation, error)
}
type Compute struct {
	Client *compute.Service
}

type Container struct {
	Client *container.Service
}

type API struct {
	Compute   *Compute
	Container *Container
	Config
}

func getConfig(filepath string) (Config, error) {
	// open the json file
	file, err := os.Open(filepath)
	if err != nil {
		return Config{}, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("Failed to close file: %v", err)
			return
		}
	}(file)
	// decode the json file
	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
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

	containerService, err := container.NewService(ctx, option.WithCredentialsFile(gcpSaFilePath))
	if err != nil {
		return nil, err
	}

	return &API{
		Compute:   &Compute{Client: computeService},
		Container: &Container{Client: containerService},
		Config:    config,
	}, nil
}

func (a *API) ListInstances(ctx context.Context, zone string) (*compute.InstanceList, error) {
	resp, err := a.Compute.Client.Instances.List(a.ProjectId, zone).Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *API) ListNetworks(ctx context.Context) (*compute.NetworkList, error) {
	resp, err := a.Compute.Client.Networks.List(a.ProjectId).Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *API) GetNetwork(ctx context.Context, nid string) (*compute.Network, error) {
	resp, err := a.Compute.Client.Networks.Get(a.ProjectId, nid).Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *API) CreateNetwork(ctx context.Context, network *compute.Network) (*compute.Operation, error) {
	resp, err := a.Compute.Client.Networks.Insert(a.ProjectId, network).Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *API) DeleteNetwork(ctx context.Context, nid string) (*compute.Operation, error) {
	resp, err := a.Compute.Client.Networks.Delete(a.ProjectId, nid).Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *API) ListClusters(ctx context.Context, zone string) (*container.ListClustersResponse, error) {
	resp, err := a.Container.Client.Projects.Zones.Clusters.List(a.ProjectId, zone).Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *API) CreateCluster(ctx context.Context, zone string, cluster *container.Cluster) (*container.Operation, error) {
	resp, err := a.Container.Client.Projects.Zones.Clusters.Create(a.ProjectId, zone, &container.CreateClusterRequest{
		Cluster:   cluster,
		Zone:      zone,
		ProjectId: a.ProjectId,
	}).Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *API) DeleteCluster(ctx context.Context, zone, clusterName string) (*container.Operation, error) {
	resp, err := a.Container.Client.Projects.Zones.Clusters.Delete(a.ProjectId, zone, clusterName).Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}
