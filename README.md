# Contabo Go SDK

Official Go SDK for the [Contabo API](https://api.contabo.com/). This SDK provides a comprehensive and easy-to-use interface for managing Contabo cloud resources including compute instances, object storage, private networks, DNS, and more.

## Features

- **Full API Coverage**: Supports all Contabo API endpoints including:
  - Compute Instances (VPS/VDS)
  - Object Storage (S3-compatible)
  - Private Networks (VPC)
  - DNS Zones and Records
  - Secrets Management
  - Tags
  - Users and Roles
- **OAuth2 Authentication**: Automatic token management and refresh
- **Type-Safe**: Strongly typed requests and responses
- **Context Support**: All methods support Go context for cancellation and timeouts
- **Error Handling**: Comprehensive error types with detailed API error information
- **Pagination Support**: Easy handling of paginated responses

## Installation

```bash
go get github.com/mithucste30/contabo-sdk-golang
```

## Quick Start

### 1. Get Your API Credentials

Obtain your API credentials from the [Contabo Customer Control Panel](https://my.contabo.com/api/details):

1. Client ID
2. Client Secret
3. API User (your email)
4. API Password

### 2. Basic Usage

```go
package main

import (
	"context"
	"fmt"
	"log"

	contabo "github.com/mithucste30/contabo-sdk-golang"
)

func main() {
	// Configure the SDK
	config := contabo.NewConfig(
		"your-client-id",
		"your-client-secret",
		"your-api-user@example.com",
		"your-api-password",
	)

	// Create SDK instance
	sdk, err := contabo.NewSDK(config)
	if err != nil {
		log.Fatalf("Failed to create SDK: %v", err)
	}

	ctx := context.Background()

	// List compute instances
	instances, err := sdk.Compute.ListInstances(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to list instances: %v", err)
	}

	for _, instance := range instances.Data {
		fmt.Printf("Instance: %s (%d) - Status: %s\n",
			instance.DisplayName, instance.InstanceID, instance.Status)
	}
}
```

### 3. Using Environment Variables

```go
package main

import (
	"context"
	"os"

	contabo "github.com/mithucste30/contabo-sdk-golang"
)

func main() {
	config := contabo.NewConfig(
		os.Getenv("CONTABO_CLIENT_ID"),
		os.Getenv("CONTABO_CLIENT_SECRET"),
		os.Getenv("CONTABO_API_USER"),
		os.Getenv("CONTABO_API_PASSWORD"),
	)

	sdk, _ := contabo.NewSDK(config)
	ctx := context.Background()

	// Use the SDK
	instances, _ := sdk.Compute.ListInstances(ctx, nil)
	// ...
}
```

## Service Overview

### Compute Service

Manage compute instances, images, and snapshots:

```go
// List instances
instances, err := sdk.Compute.ListInstances(ctx, &contabo.ListOptions{
	Page: 1,
	Size: 10,
	OrderBy: []string{"name:asc"},
})

// Get specific instance
instance, err := sdk.Compute.GetInstance(ctx, 12345)

// Create instance
newInstance, err := sdk.Compute.CreateInstance(ctx, &compute.CreateInstanceRequest{
	ImageID:   "ubuntu-22.04",
	ProductID: "V1",
	Region:    "EU",
	Period:    1,
	DisplayName: "my-server",
})

// Instance actions
err = sdk.Compute.StartInstance(ctx, instanceID)
err = sdk.Compute.StopInstance(ctx, instanceID)
err = sdk.Compute.RestartInstance(ctx, instanceID)
err = sdk.Compute.ShutdownInstance(ctx, instanceID)

// Snapshots
snapshot, err := sdk.Compute.CreateSnapshot(ctx, instanceID, &compute.CreateSnapshotRequest{
	Name: "backup-2024",
	Description: "Monthly backup",
})

snapshots, err := sdk.Compute.ListSnapshots(ctx, instanceID, nil)
err = sdk.Compute.DeleteSnapshot(ctx, instanceID, snapshotID)
instance, err := sdk.Compute.RollbackSnapshot(ctx, instanceID, snapshotID)

// Images
images, err := sdk.Compute.ListImages(ctx, nil, nil)
customImage, err := sdk.Compute.CreateImage(ctx, &compute.CreateImageRequest{
	Name:    "my-custom-image",
	URL:     "https://example.com/image.qcow2",
	OSType:  "Linux",
	Version: "1.0",
})
```

### Storage Service

Manage S3-compatible object storage:

```go
// List object storages
storages, err := sdk.Storage.ListObjectStorages(ctx, nil)

// Create object storage
storage, err := sdk.Storage.CreateObjectStorage(ctx, &storage.CreateObjectStorageRequest{
	Region: "EU",
	TotalPurchasedSpaceTB: 1.0,
	DisplayName: "my-storage",
})

// Get storage statistics
stats, err := sdk.Storage.GetObjectStorageStats(ctx, storageID)

// Get S3 credentials
creds, err := sdk.Storage.GetCredentials(ctx, storageID)
fmt.Printf("Access Key: %s\nSecret Key: %s\n", creds.AccessKey, creds.SecretKey)

// Upgrade storage
upgraded, err := sdk.Storage.UpgradeObjectStorage(ctx, storageID, &storage.UpgradeObjectStorageRequest{
	TotalPurchasedSpaceTB: 5.0,
})
```

### Network Service

Manage private networks (VPC):

```go
// List private networks
networks, err := sdk.Network.ListPrivateNetworks(ctx, nil)

// Create private network
network, err := sdk.Network.CreatePrivateNetwork(ctx, &network.CreatePrivateNetworkRequest{
	Name:        "my-vpc",
	Region:      "EU",
	Description: "Production VPC",
	InstanceIDs: []int64{12345, 67890},
})

// Assign instances to network
err = sdk.Network.AssignInstances(ctx, networkID, &network.AssignInstanceRequest{
	InstanceIDs: []int64{11111},
})

// Unassign instances
err = sdk.Network.UnassignInstances(ctx, networkID, &network.UnassignInstanceRequest{
	InstanceIDs: []int64{11111},
})
```

### DNS Service

Manage DNS zones and records:

```go
// Create DNS zone
zone, err := sdk.DNS.CreateZone(ctx, &dns.CreateZoneRequest{
	Name: "example.com",
})

// List zones
zones, err := sdk.DNS.ListZones(ctx, nil)

// Create DNS record
record, err := sdk.DNS.CreateRecord(ctx, "example.com", &dns.CreateRecordRequest{
	Name:    "www",
	Type:    "A",
	Content: "192.0.2.1",
	TTL:     3600,
})

// List records
records, err := sdk.DNS.ListRecords(ctx, "example.com", nil)

// Update PTR record (reverse DNS)
ptr, err := sdk.DNS.UpdatePTRRecord(ctx, "192.0.2.1", &dns.PatchPTRRequest{
	PTR: "server.example.com",
})
```

### Secret Service

Manage SSH keys and passwords:

```go
// Create SSH key secret
secret, err := sdk.Secret.CreateSecret(ctx, &secret.CreateSecretRequest{
	Name:  "my-ssh-key",
	Type:  "ssh",
	Value: "ssh-rsa AAAAB3...",
})

// List secrets
secrets, err := sdk.Secret.ListSecrets(ctx, nil)

// Update secret
updated, err := sdk.Secret.UpdateSecret(ctx, secretID, &secret.PatchSecretRequest{
	Name: stringPtr("updated-name"),
})

// Delete secret
err = sdk.Secret.DeleteSecret(ctx, secretID)
```

### Tag Service

Organize resources with tags:

```go
// Create tag
tag, err := sdk.Tag.CreateTag(ctx, &tag.CreateTagRequest{
	Name:  "production",
	Color: "#FF0000",
})

// List tags
tags, err := sdk.Tag.ListTags(ctx, nil)

// Assign tag to resource
err = sdk.Tag.AssignTag(ctx, tagID, "instance", "12345")

// Unassign tag
err = sdk.Tag.UnassignTag(ctx, tagID, "instance", "12345")
```

### User Service

Manage users and roles:

```go
// Create user
user, err := sdk.User.CreateUser(ctx, &user.CreateUserRequest{
	FirstName: "John",
	LastName:  "Doe",
	Email:     "john.doe@example.com",
	Enabled:   true,
	Admin:     false,
})

// List users
users, err := sdk.User.ListUsers(ctx, nil)

// Create role
role, err := sdk.User.CreateRole(ctx, &user.CreateRoleRequest{
	Name:               "read-only",
	Admin:              false,
	AccessAllResources: false,
	Type:               "apiPermission",
	Permissions: map[string]string{
		"GET /v1/compute/instances": "allow",
	},
})

// List roles
roles, err := sdk.User.ListRoles(ctx, nil)
```

## Pagination

Handle paginated responses easily:

```go
opts := &contabo.ListOptions{
	Page:    1,
	Size:    25,
	OrderBy: []string{"createdDate:desc"},
}

for {
	instances, err := sdk.Compute.ListInstances(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	// Process instances
	for _, instance := range instances.Data {
		fmt.Printf("Instance: %s\n", instance.DisplayName)
	}

	// Check if there are more pages
	if opts.Page >= instances.Pagination.TotalPages {
		break
	}

	opts.Page++
}
```

## Error Handling

The SDK provides detailed error information:

```go
instance, err := sdk.Compute.GetInstance(ctx, 12345)
if err != nil {
	// Check if it's an API error
	if apiErr, ok := err.(*contabo.APIError); ok {
		fmt.Printf("API Error: %d - %s\n", apiErr.StatusCode, apiErr.Message)
		fmt.Printf("Request ID: %s\n", apiErr.RequestID)
		fmt.Printf("Trace ID: %s\n", apiErr.TraceID)
	} else {
		fmt.Printf("Error: %v\n", err)
	}
}
```

## Context and Tracing

Add trace IDs for request grouping:

```go
// Add trace ID to context
ctx := context.WithValue(context.Background(), "x-trace-id", "my-trace-id")

// All requests in this context will include the trace ID
instances, err := sdk.Compute.ListInstances(ctx, nil)
```

## Examples

See the [examples](./examples) directory for complete working examples:

```bash
cd examples
export CONTABO_CLIENT_ID="your-client-id"
export CONTABO_CLIENT_SECRET="your-client-secret"
export CONTABO_API_USER="your-api-user@example.com"
export CONTABO_API_PASSWORD="your-api-password"
go run main.go
```

## API Documentation

For detailed API documentation, refer to:
- [Contabo API Documentation](https://api.contabo.com/)
- [OpenAPI Specification](https://api.contabo.com/api-docs)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This SDK is provided as-is for use with Contabo services.

## Support

For API-related issues, contact [Contabo Support](https://contabo.com/en/support/).

For SDK-specific issues, please open an issue on GitHub.
