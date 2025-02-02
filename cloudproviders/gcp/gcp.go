package gcp

import (
	"context"
	"encoding/json"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
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
	ListInstances(ctx context.Context, project string, zone string) (*compute.InstanceList, error)
	ListNetworks(ctx context.Context, project string) (*compute.NetworkList, error)
	GetNetwork(ctx context.Context, nid string) (*compute.Network, error)
	CreateNetwork(ctx context.Context, network *compute.Network) (*compute.Operation, error)
	DeleteNetwork(ctx context.Context, nid string) (*compute.Operation, error)
}
type Compute struct {
	Client *compute.Service
	Config
}

func getConfig(filepath string) (Config, error) {
	// open the json file
	file, err := os.Open(filepath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()
	// decode the json file
	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func NewComputeClient(ctx context.Context, gcpSaFilePath string) (*Compute, error) {
	config, err := getConfig(gcpSaFilePath)
	if err != nil {
		return nil, err
	}

	// Authenticate using the service account file
	computeService, err := compute.NewService(ctx, option.WithCredentialsFile(gcpSaFilePath))
	if err != nil {
		return nil, err
	}
	return &Compute{Client: computeService, Config: config}, nil
}

func (g *Compute) ListInstances(ctx context.Context, project string, zone string) (*compute.InstanceList, error) {
	resp, err := g.Client.Instances.List(project, zone).Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *Compute) ListNetworks(ctx context.Context, project string) (*compute.NetworkList, error) {
	resp, err := g.Client.Networks.List(project).Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *Compute) GetNetwork(ctx context.Context, nid string) (*compute.Network, error) {
	resp, err := g.Client.Networks.Get(g.ProjectId, nid).Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *Compute) CreateNetwork(ctx context.Context, network *compute.Network) (*compute.Operation, error) {
	resp, err := g.Client.Networks.Insert(g.ProjectId, network).Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *Compute) DeleteNetwork(ctx context.Context, nid string) (*compute.Operation, error) {
	resp, err := g.Client.Networks.Delete(g.ProjectId, nid).Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}
