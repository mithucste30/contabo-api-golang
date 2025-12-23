package storage

import (
	"context"
	"fmt"
)

// Client interface for making API requests
type Client interface {
	Get(ctx context.Context, path string, v interface{}) error
	Post(ctx context.Context, path string, body, v interface{}) error
	Put(ctx context.Context, path string, body, v interface{}) error
	Patch(ctx context.Context, path string, body, v interface{}) error
	Delete(ctx context.Context, path string) error
}

// ListOptions represents common query parameters for list operations
type ListOptions struct {
	Page    int
	Size    int
	OrderBy []string
}

// Service handles object storage-related API operations
type Service struct {
	client Client
}

// NewService creates a new storage service
func NewService(client Client) *Service {
	return &Service{client: client}
}

// ListObjectStorages retrieves a list of object storages
func (s *Service) ListObjectStorages(ctx context.Context, opts *ListOptions) (*ObjectStoragesResponse, error) {
	path := "/v1/object-storages"
	if opts != nil {
		path += buildQueryString(opts, nil)
	}

	var resp ObjectStoragesResponse
	if err := s.client.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetObjectStorage retrieves a specific object storage by ID
func (s *Service) GetObjectStorage(ctx context.Context, objectStorageID string) (*ObjectStorage, error) {
	path := fmt.Sprintf("/v1/object-storages/%s", objectStorageID)

	var resp struct {
		Data []ObjectStorage `json:"data"`
	}
	if err := s.client.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("object storage not found")
	}

	return &resp.Data[0], nil
}

// CreateObjectStorage creates a new object storage
func (s *Service) CreateObjectStorage(ctx context.Context, req *CreateObjectStorageRequest) (*ObjectStorage, error) {
	path := "/v1/object-storages"

	var resp CreateObjectStorageResponse
	if err := s.client.Post(ctx, path, req, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no object storage returned")
	}

	return &resp.Data[0], nil
}

// UpdateObjectStorage updates an object storage
func (s *Service) UpdateObjectStorage(ctx context.Context, objectStorageID string, req *PatchObjectStorageRequest) (*ObjectStorage, error) {
	path := fmt.Sprintf("/v1/object-storages/%s", objectStorageID)

	var resp struct {
		Data []ObjectStorage `json:"data"`
	}
	if err := s.client.Patch(ctx, path, req, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no object storage returned")
	}

	return &resp.Data[0], nil
}

// UpgradeObjectStorage upgrades object storage capacity
func (s *Service) UpgradeObjectStorage(ctx context.Context, objectStorageID string, req *UpgradeObjectStorageRequest) (*ObjectStorage, error) {
	path := fmt.Sprintf("/v1/object-storages/%s/resize", objectStorageID)

	var resp struct {
		Data []ObjectStorage `json:"data"`
	}
	if err := s.client.Post(ctx, path, req, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no object storage returned")
	}

	return &resp.Data[0], nil
}

// CancelObjectStorage cancels an object storage
func (s *Service) CancelObjectStorage(ctx context.Context, objectStorageID string) error {
	path := fmt.Sprintf("/v1/object-storages/%s/cancel", objectStorageID)
	return s.client.Post(ctx, path, nil, nil)
}

// GetObjectStorageStats retrieves usage statistics for object storage
func (s *Service) GetObjectStorageStats(ctx context.Context, objectStorageID string) (*ObjectStorageStats, error) {
	path := fmt.Sprintf("/v1/object-storages/%s/stats", objectStorageID)

	var resp ObjectStorageStatsResponse
	if err := s.client.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("stats not found")
	}

	return &resp.Data[0], nil
}

// GetCredentials retrieves S3 credentials for object storage
func (s *Service) GetCredentials(ctx context.Context, objectStorageID string) (*Credentials, error) {
	path := fmt.Sprintf("/v1/users/object-storage-credentials/%s", objectStorageID)

	var resp CredentialsResponse
	if err := s.client.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("credentials not found")
	}

	return &resp.Data[0], nil
}

// buildQueryString builds a query string from ListOptions and additional parameters
func buildQueryString(opts *ListOptions, params map[string]string) string {
	values := make(map[string][]string)

	if opts != nil {
		if opts.Page > 0 {
			values["page"] = []string{fmt.Sprintf("%d", opts.Page)}
		}
		if opts.Size > 0 {
			values["size"] = []string{fmt.Sprintf("%d", opts.Size)}
		}
		if len(opts.OrderBy) > 0 {
			values["orderBy"] = opts.OrderBy
		}
	}

	for k, v := range params {
		if v != "" {
			values[k] = []string{v}
		}
	}

	if len(values) == 0 {
		return ""
	}

	query := "?"
	first := true
	for k, vlist := range values {
		for _, v := range vlist {
			if !first {
				query += "&"
			}
			query += k + "=" + v
			first = false
		}
	}

	return query
}
