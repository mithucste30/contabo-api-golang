package compute

import "time"

// Instance represents a compute instance (VPS/VDS)
type Instance struct {
	TenantID      string    `json:"tenantId"`
	CustomerID    string    `json:"customerId"`
	InstanceID    int64     `json:"instanceId"`
	Name          string    `json:"name"`
	DisplayName   string    `json:"displayName"`
	Status        string    `json:"status"`
	ImageID       string    `json:"imageId"`
	ImageName     string    `json:"imageName"`
	ProductID     string    `json:"productId"`
	Region        string    `json:"region"`
	DataCenter    string    `json:"dataCenter"`
	CreatedDate   time.Time `json:"createdDate"`
	CancelDate    string    `json:"cancelDate,omitempty"`
	IPConfig      IPConfig  `json:"ipConfig"`
	MACAddress    string    `json:"macAddress"`
	RAMMemoryMB   float64   `json:"ramMb"`
	CPUCores      int       `json:"cpuCores"`
	DiskMB        float64   `json:"diskMb"`
	OSType        string    `json:"osType"`
	SSHKeys       []int64   `json:"sshKeys,omitempty"`
	DefaultUser   string    `json:"defaultUser,omitempty"`
}

// IPConfig represents the IP configuration of an instance
type IPConfig struct {
	V4 IPConfigV4 `json:"v4"`
	V6 IPConfigV6 `json:"v6"`
}

// IPConfigV4 represents IPv4 configuration
type IPConfigV4 struct {
	IP      string `json:"ip"`
	Gateway string `json:"gateway"`
	Netmask string `json:"netmask"`
}

// IPConfigV6 represents IPv6 configuration
type IPConfigV6 struct {
	IP      string `json:"ip"`
	Gateway string `json:"gateway"`
	Netmask string `json:"netmask"`
}

// InstancesResponse represents the response for listing instances
type InstancesResponse struct {
	Pagination struct {
		Size          int   `json:"size"`
		TotalElements int64 `json:"totalElements"`
		TotalPages    int   `json:"totalPages"`
		Number        int   `json:"number"`
	} `json:"_pagination"`
	Links struct {
		Self     string `json:"self"`
		First    string `json:"first,omitempty"`
		Previous string `json:"previous,omitempty"`
		Next     string `json:"next,omitempty"`
		Last     string `json:"last,omitempty"`
	} `json:"_links"`
	Data []Instance `json:"data"`
}

// CreateInstanceRequest represents the request body for creating an instance
type CreateInstanceRequest struct {
	ImageID        string  `json:"imageId"`
	ProductID      string  `json:"productId"`
	Region         string  `json:"region"`
	SSHKeys        []int64 `json:"sshKeys,omitempty"`
	RootPassword   int64   `json:"rootPassword,omitempty"`
	UserData       string  `json:"userData,omitempty"`
	License        string  `json:"license,omitempty"`
	Period         int64   `json:"period"`
	DisplayName    string  `json:"displayName,omitempty"`
	DefaultUser    string  `json:"defaultUser,omitempty"`
	AddOns         *AddOns `json:"addOns,omitempty"`
	ApplicationID  string  `json:"applicationId,omitempty"`
}

// AddOns represents additional services for an instance
type AddOns struct {
	PrivateNetworking *PrivateNetworkingAddOn `json:"privateNetworking,omitempty"`
}

// PrivateNetworkingAddOn represents private networking addon configuration
type PrivateNetworkingAddOn struct {
	Enabled bool `json:"enabled"`
}

// CreateInstanceResponse represents the response when creating an instance
type CreateInstanceResponse struct {
	Data []Instance `json:"data"`
}

// PatchInstanceRequest represents the request body for updating an instance
type PatchInstanceRequest struct {
	DisplayName *string `json:"displayName,omitempty"`
}

// UpgradeInstanceRequest represents the request body for upgrading an instance
type UpgradeInstanceRequest struct {
	ProductID string `json:"productId"`
}

// ReinstallInstanceRequest represents the request body for reinstalling an instance
type ReinstallInstanceRequest struct {
	ImageID      string  `json:"imageId"`
	SSHKeys      []int64 `json:"sshKeys,omitempty"`
	RootPassword int64   `json:"rootPassword,omitempty"`
	UserData     string  `json:"userData,omitempty"`
	DefaultUser  string  `json:"defaultUser,omitempty"`
}

// RescueInstanceRequest represents the request body for rescue mode
type RescueInstanceRequest struct {
	RootPassword int64   `json:"rootPassword,omitempty"`
	SSHKeys      []int64 `json:"sshKeys,omitempty"`
	UserData     string  `json:"userData,omitempty"`
}

// ResetPasswordRequest represents the request body for password reset
type ResetPasswordRequest struct {
	RootPassword int64 `json:"rootPassword,omitempty"`
}

// Snapshot represents an instance snapshot
type Snapshot struct {
	TenantID    string    `json:"tenantId"`
	CustomerID  string    `json:"customerId"`
	SnapshotID  string    `json:"snapshotId"`
	InstanceID  int64     `json:"instanceId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedDate time.Time `json:"createdDate"`
	AutoDelete  bool      `json:"autoDelete"`
}

// SnapshotsResponse represents the response for listing snapshots
type SnapshotsResponse struct {
	Pagination struct {
		Size          int   `json:"size"`
		TotalElements int64 `json:"totalElements"`
		TotalPages    int   `json:"totalPages"`
		Number        int   `json:"number"`
	} `json:"_pagination"`
	Links struct {
		Self     string `json:"self"`
		First    string `json:"first,omitempty"`
		Previous string `json:"previous,omitempty"`
		Next     string `json:"next,omitempty"`
		Last     string `json:"last,omitempty"`
	} `json:"_links"`
	Data []Snapshot `json:"data"`
}

// CreateSnapshotRequest represents the request body for creating a snapshot
type CreateSnapshotRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// CreateSnapshotResponse represents the response when creating a snapshot
type CreateSnapshotResponse struct {
	Data []Snapshot `json:"data"`
}

// PatchSnapshotRequest represents the request body for updating a snapshot
type PatchSnapshotRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// RollbackSnapshotResponse represents the response when rolling back to a snapshot
type RollbackSnapshotResponse struct {
	Data []Instance `json:"data"`
}

// Image represents a compute image
type Image struct {
	ImageID      string   `json:"imageId"`
	TenantID     string   `json:"tenantId"`
	CustomerID   string   `json:"customerId"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	URL          string   `json:"url"`
	SizeMB       float64  `json:"sizeMb"`
	UploadedSizeMB float64 `json:"uploadedSizeMb"`
	OSType       string   `json:"osType"`
	Version      string   `json:"version"`
	Format       string   `json:"format"`
	Status       string   `json:"status"`
	ErrorMessage string   `json:"errorMessage,omitempty"`
	StandardImage bool    `json:"standardImage"`
	CreatedDate  time.Time `json:"createdDate"`
	LastModifiedDate time.Time `json:"lastModifiedDate"`
}

// ImagesResponse represents the response for listing images
type ImagesResponse struct {
	Pagination struct {
		Size          int   `json:"size"`
		TotalElements int64 `json:"totalElements"`
		TotalPages    int   `json:"totalPages"`
		Number        int   `json:"number"`
	} `json:"_pagination"`
	Links struct {
		Self     string `json:"self"`
		First    string `json:"first,omitempty"`
		Previous string `json:"previous,omitempty"`
		Next     string `json:"next,omitempty"`
		Last     string `json:"last,omitempty"`
	} `json:"_links"`
	Data []Image `json:"data"`
}

// CreateImageRequest represents the request body for creating a custom image
type CreateImageRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url"`
	OSType      string `json:"osType"`
	Version     string `json:"version"`
}

// CreateImageResponse represents the response when creating an image
type CreateImageResponse struct {
	Data []Image `json:"data"`
}

// PatchImageRequest represents the request body for updating an image
type PatchImageRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}
