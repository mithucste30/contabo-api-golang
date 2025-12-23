package compute

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

// Service handles compute-related API operations
type Service struct {
	client Client
}

// NewService creates a new compute service
func NewService(client Client) *Service {
	return &Service{client: client}
}

// Instances

// ListInstances retrieves a list of compute instances
func (s *Service) ListInstances(ctx context.Context, opts *ListOptions) (*InstancesResponse, error) {
	path := "/v1/compute/instances"
	if opts != nil {
		path += buildQueryString(opts, nil)
	}

	var resp InstancesResponse
	if err := s.client.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetInstance retrieves a specific instance by ID
func (s *Service) GetInstance(ctx context.Context, instanceID int64) (*Instance, error) {
	path := fmt.Sprintf("/v1/compute/instances/%d", instanceID)

	var resp struct {
		Data []Instance `json:"data"`
	}
	if err := s.client.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("instance not found")
	}

	return &resp.Data[0], nil
}

// CreateInstance creates a new compute instance
func (s *Service) CreateInstance(ctx context.Context, req *CreateInstanceRequest) (*Instance, error) {
	path := "/v1/compute/instances"

	var resp CreateInstanceResponse
	if err := s.client.Post(ctx, path, req, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no instance returned")
	}

	return &resp.Data[0], nil
}

// UpdateInstance updates an instance (PATCH)
func (s *Service) UpdateInstance(ctx context.Context, instanceID int64, req *PatchInstanceRequest) (*Instance, error) {
	path := fmt.Sprintf("/v1/compute/instances/%d", instanceID)

	var resp struct {
		Data []Instance `json:"data"`
	}
	if err := s.client.Patch(ctx, path, req, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no instance returned")
	}

	return &resp.Data[0], nil
}

// ReinstallInstance reinstalls an instance with a new image
func (s *Service) ReinstallInstance(ctx context.Context, instanceID int64, req *ReinstallInstanceRequest) (*Instance, error) {
	path := fmt.Sprintf("/v1/compute/instances/%d", instanceID)

	var resp struct {
		Data []Instance `json:"data"`
	}
	if err := s.client.Put(ctx, path, req, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no instance returned")
	}

	return &resp.Data[0], nil
}

// CancelInstance cancels an instance
func (s *Service) CancelInstance(ctx context.Context, instanceID int64) error {
	path := fmt.Sprintf("/v1/compute/instances/%d/cancel", instanceID)
	return s.client.Post(ctx, path, nil, nil)
}

// UpgradeInstance upgrades an instance to a different product
func (s *Service) UpgradeInstance(ctx context.Context, instanceID int64, req *UpgradeInstanceRequest) (*Instance, error) {
	path := fmt.Sprintf("/v1/compute/instances/%d/upgrade", instanceID)

	var resp struct {
		Data []Instance `json:"data"`
	}
	if err := s.client.Post(ctx, path, req, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no instance returned")
	}

	return &resp.Data[0], nil
}

// Instance Actions

// StartInstance starts a stopped instance
func (s *Service) StartInstance(ctx context.Context, instanceID int64) error {
	path := fmt.Sprintf("/v1/compute/instances/%d/actions/start", instanceID)
	return s.client.Post(ctx, path, nil, nil)
}

// StopInstance stops a running instance
func (s *Service) StopInstance(ctx context.Context, instanceID int64) error {
	path := fmt.Sprintf("/v1/compute/instances/%d/actions/stop", instanceID)
	return s.client.Post(ctx, path, nil, nil)
}

// RestartInstance restarts an instance
func (s *Service) RestartInstance(ctx context.Context, instanceID int64) error {
	path := fmt.Sprintf("/v1/compute/instances/%d/actions/restart", instanceID)
	return s.client.Post(ctx, path, nil, nil)
}

// ShutdownInstance gracefully shuts down an instance
func (s *Service) ShutdownInstance(ctx context.Context, instanceID int64) error {
	path := fmt.Sprintf("/v1/compute/instances/%d/actions/shutdown", instanceID)
	return s.client.Post(ctx, path, nil, nil)
}

// RescueInstance puts an instance into rescue mode
func (s *Service) RescueInstance(ctx context.Context, instanceID int64, req *RescueInstanceRequest) error {
	path := fmt.Sprintf("/v1/compute/instances/%d/actions/rescue", instanceID)
	return s.client.Post(ctx, path, req, nil)
}

// ResetPassword resets the root password of an instance
func (s *Service) ResetPassword(ctx context.Context, instanceID int64, req *ResetPasswordRequest) error {
	path := fmt.Sprintf("/v1/compute/instances/%d/actions/resetPassword", instanceID)
	return s.client.Post(ctx, path, req, nil)
}

// Snapshots

// ListSnapshots retrieves snapshots for an instance
func (s *Service) ListSnapshots(ctx context.Context, instanceID int64, opts *ListOptions) (*SnapshotsResponse, error) {
	path := fmt.Sprintf("/v1/compute/instances/%d/snapshots", instanceID)
	if opts != nil {
		path += buildQueryString(opts, nil)
	}

	var resp SnapshotsResponse
	if err := s.client.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetSnapshot retrieves a specific snapshot
func (s *Service) GetSnapshot(ctx context.Context, instanceID int64, snapshotID string) (*Snapshot, error) {
	path := fmt.Sprintf("/v1/compute/instances/%d/snapshots/%s", instanceID, snapshotID)

	var resp struct {
		Data []Snapshot `json:"data"`
	}
	if err := s.client.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("snapshot not found")
	}

	return &resp.Data[0], nil
}

// CreateSnapshot creates a new snapshot of an instance
func (s *Service) CreateSnapshot(ctx context.Context, instanceID int64, req *CreateSnapshotRequest) (*Snapshot, error) {
	path := fmt.Sprintf("/v1/compute/instances/%d/snapshots", instanceID)

	var resp CreateSnapshotResponse
	if err := s.client.Post(ctx, path, req, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no snapshot returned")
	}

	return &resp.Data[0], nil
}

// UpdateSnapshot updates a snapshot
func (s *Service) UpdateSnapshot(ctx context.Context, instanceID int64, snapshotID string, req *PatchSnapshotRequest) (*Snapshot, error) {
	path := fmt.Sprintf("/v1/compute/instances/%d/snapshots/%s", instanceID, snapshotID)

	var resp struct {
		Data []Snapshot `json:"data"`
	}
	if err := s.client.Patch(ctx, path, req, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no snapshot returned")
	}

	return &resp.Data[0], nil
}

// DeleteSnapshot deletes a snapshot
func (s *Service) DeleteSnapshot(ctx context.Context, instanceID int64, snapshotID string) error {
	path := fmt.Sprintf("/v1/compute/instances/%d/snapshots/%s", instanceID, snapshotID)
	return s.client.Delete(ctx, path)
}

// RollbackSnapshot rolls back an instance to a snapshot
func (s *Service) RollbackSnapshot(ctx context.Context, instanceID int64, snapshotID string) (*Instance, error) {
	path := fmt.Sprintf("/v1/compute/instances/%d/snapshots/%s/rollback", instanceID, snapshotID)

	var resp RollbackSnapshotResponse
	if err := s.client.Post(ctx, path, nil, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no instance returned")
	}

	return &resp.Data[0], nil
}

// Images

// ListImages retrieves a list of images
func (s *Service) ListImages(ctx context.Context, opts *ListOptions, standardImage *bool) (*ImagesResponse, error) {
	path := "/v1/compute/images"

	params := make(map[string]string)
	if standardImage != nil {
		if *standardImage {
			params["standardImage"] = "true"
		} else {
			params["standardImage"] = "false"
		}
	}

	if opts != nil || len(params) > 0 {
		path += buildQueryString(opts, params)
	}

	var resp ImagesResponse
	if err := s.client.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetImage retrieves a specific image by ID
func (s *Service) GetImage(ctx context.Context, imageID string) (*Image, error) {
	path := fmt.Sprintf("/v1/compute/images/%s", imageID)

	var resp struct {
		Data []Image `json:"data"`
	}
	if err := s.client.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("image not found")
	}

	return &resp.Data[0], nil
}

// CreateImage creates a new custom image
func (s *Service) CreateImage(ctx context.Context, req *CreateImageRequest) (*Image, error) {
	path := "/v1/compute/images"

	var resp CreateImageResponse
	if err := s.client.Post(ctx, path, req, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no image returned")
	}

	return &resp.Data[0], nil
}

// UpdateImage updates a custom image
func (s *Service) UpdateImage(ctx context.Context, imageID string, req *PatchImageRequest) (*Image, error) {
	path := fmt.Sprintf("/v1/compute/images/%s", imageID)

	var resp struct {
		Data []Image `json:"data"`
	}
	if err := s.client.Patch(ctx, path, req, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no image returned")
	}

	return &resp.Data[0], nil
}

// DeleteImage deletes a custom image
func (s *Service) DeleteImage(ctx context.Context, imageID string) error {
	path := fmt.Sprintf("/v1/compute/images/%s", imageID)
	return s.client.Delete(ctx, path)
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
