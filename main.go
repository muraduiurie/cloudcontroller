package main

import (
	"context"
	"gitlab.com/charmelionag/cloudcontroller/cloudproviders/gcp"
	"log"
)

const (
	gcpSaFilePath = "./creds/gcp-creds.json"
)

func main() {
	ctx := context.Background()
	//zone := "us-central1-a"

	gcpApi, err := gcp.NewAPI(ctx, gcpSaFilePath)
	if err != nil {
		log.Fatalf("Error creating GCP API: %v", err)
	}

	//List instances
	//resp, err := gcpApi.CreateNetwork(ctx, &compute.Network{
	//	AutoCreateSubnetworks: true,
	//	Name:                  "test-network",
	//})
	//if err != nil {
	//	log.Fatalf("Error creating network: %v", err)
	//}
	//log.Printf("Network created: %v", resp)

	//resp, err := gcpCompute.DeleteNetwork(ctx, "test-network")
	//if err != nil {
	//	log.Fatalf("Error deleting network: %v", err)
	//}
	//log.Printf("Network deleted: %v", resp.Status)

	// List GKE clusters
	lresp, err := gcpApi.ListClusters(ctx, "us-central1-a")
	if err != nil {
		log.Fatalf("Error listing clusters: %v", err)
	}
	log.Printf("Clusters: %v", lresp.Clusters[0].Name)

	// Create a GKE cluster
	//resp, err := gcpApi.CreateCluster(ctx, "us-central1-a", &container.Cluster{
	//	Name:             "test-cluster",
	//	InitialNodeCount: 1,
	//})
	//if err != nil {
	//	log.Fatalf("Error creating cluster: %v", err)
	//}
	//log.Printf("Cluster created: %v", resp)

	// Delete Cluster
	resp, err := gcpApi.DeleteCluster(ctx, "us-central1-a", lresp.Clusters[0].Name)
	if err != nil {
		log.Fatalf("Error deleting cluster: %v", err)
	}
	log.Printf("Cluster deleted: %v", resp.Status)
}
