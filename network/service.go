package network

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

// Service handles network-related API operations
type Service struct {
	client Client
}

// NewService creates a new network service
func NewService(client Client) *Service {
	return &Service{client: client}
}

// ListPrivateNetworks retrieves a list of private networks
func (s *Service) ListPrivateNetworks(ctx context.Context, opts *ListOptions) (*PrivateNetworksResponse, error) {
	path := "/v1/private-networks"
	if opts != nil {
		path += buildQueryString(opts, nil)
	}

	var resp PrivateNetworksResponse
	if err := s.client.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetPrivateNetwork retrieves a specific private network by ID
func (s *Service) GetPrivateNetwork(ctx context.Context, privateNetworkID int64) (*PrivateNetwork, error) {
	path := fmt.Sprintf("/v1/private-networks/%d", privateNetworkID)

	var resp struct {
		Data []PrivateNetwork `json:"data"`
	}
	if err := s.client.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("private network not found")
	}

	return &resp.Data[0], nil
}

// CreatePrivateNetwork creates a new private network
func (s *Service) CreatePrivateNetwork(ctx context.Context, req *CreatePrivateNetworkRequest) (*PrivateNetwork, error) {
	path := "/v1/private-networks"

	var resp CreatePrivateNetworkResponse
	if err := s.client.Post(ctx, path, req, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no private network returned")
	}

	return &resp.Data[0], nil
}

// UpdatePrivateNetwork updates a private network
func (s *Service) UpdatePrivateNetwork(ctx context.Context, privateNetworkID int64, req *PatchPrivateNetworkRequest) (*PrivateNetwork, error) {
	path := fmt.Sprintf("/v1/private-networks/%d", privateNetworkID)

	var resp struct {
		Data []PrivateNetwork `json:"data"`
	}
	if err := s.client.Patch(ctx, path, req, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no private network returned")
	}

	return &resp.Data[0], nil
}

// DeletePrivateNetwork deletes a private network
func (s *Service) DeletePrivateNetwork(ctx context.Context, privateNetworkID int64) error {
	path := fmt.Sprintf("/v1/private-networks/%d", privateNetworkID)
	return s.client.Delete(ctx, path)
}

// AssignInstances assigns instances to a private network
func (s *Service) AssignInstances(ctx context.Context, privateNetworkID int64, req *AssignInstanceRequest) error {
	path := fmt.Sprintf("/v1/private-networks/%d/instances", privateNetworkID)
	return s.client.Post(ctx, path, req, nil)
}

// UnassignInstances unassigns instances from a private network
func (s *Service) UnassignInstances(ctx context.Context, privateNetworkID int64, req *UnassignInstanceRequest) error {
	path := fmt.Sprintf("/v1/private-networks/%d/instances", privateNetworkID)
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
