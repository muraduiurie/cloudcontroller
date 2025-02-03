package main

import (
	"context"
	"gitlab.com/charmelionag/cloudcontroller/cloudproviders/gcp"
	"google.golang.org/api/compute/v1"
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
	resp, err := gcpApi.CreateNetwork(ctx, &compute.Network{
		AutoCreateSubnetworks: true,
		Name:                  "test-network",
	})
	if err != nil {
		log.Fatalf("Error creating network: %v", err)
	}
	log.Printf("Network created: %v", resp)

	//resp, err := gcpCompute.DeleteNetwork(ctx, "test-network")
	//if err != nil {
	//	log.Fatalf("Error deleting network: %v", err)
	//}
	//log.Printf("Network deleted: %v", resp.Status)
}
