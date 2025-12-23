package storage

import "time"

// ObjectStorage represents an S3-compatible object storage
type ObjectStorage struct {
	TenantID        string    `json:"tenantId"`
	CustomerID      string    `json:"customerId"`
	ObjectStorageID string    `json:"objectStorageId"`
	CreatedDate     time.Time `json:"createdDate"`
	CancelDate      string    `json:"cancelDate,omitempty"`
	AutoScaling     AutoScaling `json:"autoScaling"`
	DataCenter      string    `json:"dataCenter"`
	TotalPurchasedSpaceTB float64 `json:"totalPurchasedSpaceTB"`
	S3URL           string    `json:"s3Url"`
	S3TenantID      string    `json:"s3TenantId"`
	Status          string    `json:"status"`
	Region          string    `json:"region"`
	DisplayName     string    `json:"displayName,omitempty"`
}

// AutoScaling represents auto-scaling configuration
type AutoScaling struct {
	State          string  `json:"state"`
	SizeLimitTB    float64 `json:"sizeLimitTB"`
	ErrorMessage   string  `json:"errorMessage,omitempty"`
}

// ObjectStoragesResponse represents the response for listing object storages
type ObjectStoragesResponse struct {
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
	Data []ObjectStorage `json:"data"`
}

// CreateObjectStorageRequest represents the request body for creating object storage
type CreateObjectStorageRequest struct {
	Region      string  `json:"region"`
	TotalPurchasedSpaceTB float64 `json:"totalPurchasedSpaceTB"`
	AutoScaling *AutoScalingRequest `json:"autoScaling,omitempty"`
	DisplayName string  `json:"displayName,omitempty"`
}

// AutoScalingRequest represents auto-scaling configuration for creation
type AutoScalingRequest struct {
	State       string  `json:"state"`
	SizeLimitTB float64 `json:"sizeLimitTB"`
}

// CreateObjectStorageResponse represents the response when creating object storage
type CreateObjectStorageResponse struct {
	Data []ObjectStorage `json:"data"`
}

// PatchObjectStorageRequest represents the request body for updating object storage
type PatchObjectStorageRequest struct {
	DisplayName *string             `json:"displayName,omitempty"`
	AutoScaling *AutoScalingRequest `json:"autoScaling,omitempty"`
}

// UpgradeObjectStorageRequest represents the request body for upgrading object storage
type UpgradeObjectStorageRequest struct {
	TotalPurchasedSpaceTB float64 `json:"totalPurchasedSpaceTB"`
	AutoScaling           *AutoScalingRequest `json:"autoScaling,omitempty"`
}

// ObjectStorageStats represents usage statistics
type ObjectStorageStats struct {
	ObjectStorageID string  `json:"objectStorageId"`
	UsedSpaceTB     float64 `json:"usedSpaceTB"`
	UsedSpacePercentage float64 `json:"usedSpacePercentage"`
}

// ObjectStorageStatsResponse represents the response for storage statistics
type ObjectStorageStatsResponse struct {
	Data []ObjectStorageStats `json:"data"`
}

// Credentials represents S3 access credentials
type Credentials struct {
	TenantID    string `json:"tenantId"`
	CustomerID  string `json:"customerId"`
	AccessKey   string `json:"accessKey"`
	SecretKey   string `json:"secretKey"`
	DisplayName string `json:"displayName,omitempty"`
}

// CredentialsResponse represents the response for S3 credentials
type CredentialsResponse struct {
	Data []Credentials `json:"data"`
}
