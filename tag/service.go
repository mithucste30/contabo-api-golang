package tag

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

// Service handles tag-related API operations
type Service struct {
	client Client
}

// NewService creates a new tag service
func NewService(client Client) *Service {
	return &Service{client: client}
}

// ListTags retrieves a list of tags
func (s *Service) ListTags(ctx context.Context, opts *ListOptions) (*TagsResponse, error) {
	path := "/v1/tags"
	if opts != nil {
		path += buildQueryString(opts, nil)
	}

	var resp TagsResponse
	if err := s.client.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetTag retrieves a specific tag by ID
func (s *Service) GetTag(ctx context.Context, tagID int64) (*Tag, error) {
	path := fmt.Sprintf("/v1/tags/%d", tagID)

	var resp struct {
		Data []Tag `json:"data"`
	}
	if err := s.client.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("tag not found")
	}

	return &resp.Data[0], nil
}

// CreateTag creates a new tag
func (s *Service) CreateTag(ctx context.Context, req *CreateTagRequest) (*Tag, error) {
	path := "/v1/tags"

	var resp CreateTagResponse
	if err := s.client.Post(ctx, path, req, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no tag returned")
	}

	return &resp.Data[0], nil
}

// UpdateTag updates a tag
func (s *Service) UpdateTag(ctx context.Context, tagID int64, req *PatchTagRequest) (*Tag, error) {
	path := fmt.Sprintf("/v1/tags/%d", tagID)

	var resp struct {
		Data []Tag `json:"data"`
	}
	if err := s.client.Patch(ctx, path, req, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no tag returned")
	}

	return &resp.Data[0], nil
}

// DeleteTag deletes a tag
func (s *Service) DeleteTag(ctx context.Context, tagID int64) error {
	path := fmt.Sprintf("/v1/tags/%d", tagID)
	return s.client.Delete(ctx, path)
}

// AssignTag assigns a tag to a resource
func (s *Service) AssignTag(ctx context.Context, tagID int64, resourceType string, resourceID string) error {
	path := fmt.Sprintf("/v1/tags/%d/assignments/%s/%s", tagID, resourceType, resourceID)
	return s.client.Put(ctx, path, nil, nil)
}

// UnassignTag unassigns a tag from a resource
func (s *Service) UnassignTag(ctx context.Context, tagID int64, resourceType string, resourceID string) error {
	path := fmt.Sprintf("/v1/tags/%d/assignments/%s/%s", tagID, resourceType, resourceID)
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
