package secret

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

// Service handles secret-related API operations
type Service struct {
	client Client
}

// NewService creates a new secret service
func NewService(client Client) *Service {
	return &Service{client: client}
}

// ListSecrets retrieves a list of secrets
func (s *Service) ListSecrets(ctx context.Context, opts *ListOptions) (*SecretsResponse, error) {
	path := "/v1/secrets"
	if opts != nil {
		path += buildQueryString(opts, nil)
	}

	var resp SecretsResponse
	if err := s.client.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetSecret retrieves a specific secret by ID
func (s *Service) GetSecret(ctx context.Context, secretID int64) (*Secret, error) {
	path := fmt.Sprintf("/v1/secrets/%d", secretID)

	var resp struct {
		Data []Secret `json:"data"`
	}
	if err := s.client.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("secret not found")
	}

	return &resp.Data[0], nil
}

// CreateSecret creates a new secret
func (s *Service) CreateSecret(ctx context.Context, req *CreateSecretRequest) (*Secret, error) {
	path := "/v1/secrets"

	var resp CreateSecretResponse
	if err := s.client.Post(ctx, path, req, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no secret returned")
	}

	return &resp.Data[0], nil
}

// UpdateSecret updates a secret
func (s *Service) UpdateSecret(ctx context.Context, secretID int64, req *PatchSecretRequest) (*Secret, error) {
	path := fmt.Sprintf("/v1/secrets/%d", secretID)

	var resp struct {
		Data []Secret `json:"data"`
	}
	if err := s.client.Patch(ctx, path, req, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no secret returned")
	}

	return &resp.Data[0], nil
}

// DeleteSecret deletes a secret
func (s *Service) DeleteSecret(ctx context.Context, secretID int64) error {
	path := fmt.Sprintf("/v1/secrets/%d", secretID)
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
