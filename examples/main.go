package main

import (
	"context"
	"fmt"
	"log"
	"os"

	contabo "github.com/mithucste30/contabo-api-golang"
	"github.com/mithucste30/contabo-api-golang/compute"
)

func main() {
	// Load credentials from environment variables
	config := contabo.NewConfig(
		os.Getenv("CONTABO_CLIENT_ID"),
		os.Getenv("CONTABO_CLIENT_SECRET"),
		os.Getenv("CONTABO_API_USER"),
		os.Getenv("CONTABO_API_PASSWORD"),
	)

	// Create SDK instance
	sdk, err := contabo.NewSDK(config)
	if err != nil {
		log.Fatalf("Failed to create SDK: %v", err)
	}

	ctx := context.Background()

	// Example 1: List all compute instances
	fmt.Println("=== Listing Compute Instances ===")
	instances, err := sdk.Compute.ListInstances(ctx, &compute.ListOptions{
		Page: 1,
		Size: 10,
	})
	if err != nil {
		log.Fatalf("Failed to list instances: %v", err)
	}

	fmt.Printf("Found %d instances:\n", instances.Pagination.TotalElements)
	for _, instance := range instances.Data {
		fmt.Printf("  - ID: %d, Name: %s, Status: %s, Region: %s\n",
			instance.InstanceID, instance.DisplayName, instance.Status, instance.Region)
	}

	// Example 2: List all images
	fmt.Println("\n=== Listing Images ===")
	standardImage := true
	images, err := sdk.Compute.ListImages(ctx, &compute.ListOptions{
		Page: 1,
		Size: 5,
	}, &standardImage)
	if err != nil {
		log.Fatalf("Failed to list images: %v", err)
	}

	fmt.Printf("Found %d images:\n", images.Pagination.TotalElements)
	for _, image := range images.Data {
		fmt.Printf("  - ID: %s, Name: %s, OS: %s\n",
			image.ImageID, image.Name, image.OSType)
	}

	// Example 3: List object storages
	fmt.Println("\n=== Listing Object Storages ===")
	storages, err := sdk.Storage.ListObjectStorages(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to list object storages: %v", err)
	}

	fmt.Printf("Found %d object storages:\n", storages.Pagination.TotalElements)
	for _, storage := range storages.Data {
		fmt.Printf("  - ID: %s, Region: %s, Space: %.2f TB, Status: %s\n",
			storage.ObjectStorageID, storage.Region, storage.TotalPurchasedSpaceTB, storage.Status)
	}

	// Example 4: List private networks
	fmt.Println("\n=== Listing Private Networks ===")
	networks, err := sdk.Network.ListPrivateNetworks(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to list private networks: %v", err)
	}

	fmt.Printf("Found %d private networks:\n", networks.Pagination.TotalElements)
	for _, network := range networks.Data {
		fmt.Printf("  - ID: %d, Name: %s, Region: %s, CIDR: %s\n",
			network.PrivateNetworkID, network.Name, network.Region, network.CIDR)
	}

	// Example 5: Create a snapshot
	if len(instances.Data) > 0 {
		fmt.Println("\n=== Creating Snapshot ===")
		instanceID := instances.Data[0].InstanceID

		snapshot, err := sdk.Compute.CreateSnapshot(ctx, instanceID, &compute.CreateSnapshotRequest{
			Name:        "example-snapshot",
			Description: "Created via Go SDK example",
		})
		if err != nil {
			log.Printf("Failed to create snapshot: %v", err)
		} else {
			fmt.Printf("Created snapshot: ID=%s, Name=%s\n", snapshot.SnapshotID, snapshot.Name)
		}
	}

	// Example 6: List DNS zones
	fmt.Println("\n=== Listing DNS Zones ===")
	zones, err := sdk.DNS.ListZones(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to list DNS zones: %v", err)
	}

	fmt.Printf("Found %d DNS zones:\n", zones.Pagination.TotalElements)
	for _, zone := range zones.Data {
		fmt.Printf("  - ID: %s, Name: %s\n", zone.ZoneID, zone.Name)
	}

	// Example 7: List tags
	fmt.Println("\n=== Listing Tags ===")
	tags, err := sdk.Tag.ListTags(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to list tags: %v", err)
	}

	fmt.Printf("Found %d tags:\n", tags.Pagination.TotalElements)
	for _, tag := range tags.Data {
		fmt.Printf("  - ID: %d, Name: %s, Color: %s\n", tag.TagID, tag.Name, tag.Color)
	}

	fmt.Println("\n=== Examples completed successfully ===")
}
