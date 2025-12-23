package network

import "time"

// PrivateNetwork represents a private network (VPC)
type PrivateNetwork struct {
	PrivateNetworkID int64     `json:"privateNetworkId"`
	TenantID         string    `json:"tenantId"`
	CustomerID       string    `json:"customerId"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	Region           string    `json:"region"`
	RegionName       string    `json:"regionName"`
	DataCenter       string    `json:"dataCenter"`
	AvailableIPs     int64     `json:"availableIps"`
	CIDR             string    `json:"cidr"`
	CreatedDate      time.Time `json:"createdDate"`
	Instances        []Instance `json:"instances"`
}

// Instance represents an instance in a private network
type Instance struct {
	InstanceID int64  `json:"instanceId"`
	PrivateIP  string `json:"privateIp"`
}

// PrivateNetworksResponse represents the response for listing private networks
type PrivateNetworksResponse struct {
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
	Data []PrivateNetwork `json:"data"`
}

// CreatePrivateNetworkRequest represents the request body for creating a private network
type CreatePrivateNetworkRequest struct {
	Region      string `json:"region"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	InstanceIDs []int64 `json:"instanceIds,omitempty"`
}

// CreatePrivateNetworkResponse represents the response when creating a private network
type CreatePrivateNetworkResponse struct {
	Data []PrivateNetwork `json:"data"`
}

// PatchPrivateNetworkRequest represents the request body for updating a private network
type PatchPrivateNetworkRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// AssignInstanceRequest represents the request body for assigning instances to a private network
type AssignInstanceRequest struct {
	InstanceIDs []int64 `json:"instanceIds"`
}

// UnassignInstanceRequest represents the request body for unassigning instances from a private network
type UnassignInstanceRequest struct {
	InstanceIDs []int64 `json:"instanceIds"`
}
