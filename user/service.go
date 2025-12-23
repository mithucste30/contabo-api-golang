package user

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

// Service handles user-related API operations
type Service struct {
	client Client
}

// NewService creates a new user service
func NewService(client Client) *Service {
	return &Service{client: client}
}

// Users

// ListUsers retrieves a list of users
func (s *Service) ListUsers(ctx context.Context, opts *ListOptions) (*UsersResponse, error) {
	path := "/v1/users"
	if opts != nil {
		path += buildQueryString(opts, nil)
	}

	var resp UsersResponse
	if err := s.client.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetUser retrieves a specific user by ID
func (s *Service) GetUser(ctx context.Context, userID string) (*User, error) {
	path := fmt.Sprintf("/v1/users/%s", userID)

	var resp struct {
		Data []User `json:"data"`
	}
	if err := s.client.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return &resp.Data[0], nil
}

// CreateUser creates a new user
func (s *Service) CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error) {
	path := "/v1/users"

	var resp CreateUserResponse
	if err := s.client.Post(ctx, path, req, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no user returned")
	}

	return &resp.Data[0], nil
}

// UpdateUser updates a user
func (s *Service) UpdateUser(ctx context.Context, userID string, req *PatchUserRequest) (*User, error) {
	path := fmt.Sprintf("/v1/users/%s", userID)

	var resp struct {
		Data []User `json:"data"`
	}
	if err := s.client.Patch(ctx, path, req, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no user returned")
	}

	return &resp.Data[0], nil
}

// DeleteUser deletes a user
func (s *Service) DeleteUser(ctx context.Context, userID string) error {
	path := fmt.Sprintf("/v1/users/%s", userID)
	return s.client.Delete(ctx, path)
}

// Roles

// ListRoles retrieves a list of roles
func (s *Service) ListRoles(ctx context.Context, opts *ListOptions) (*RolesResponse, error) {
	path := "/v1/roles"
	if opts != nil {
		path += buildQueryString(opts, nil)
	}

	var resp RolesResponse
	if err := s.client.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetRole retrieves a specific role by ID
func (s *Service) GetRole(ctx context.Context, roleID int64) (*Role, error) {
	path := fmt.Sprintf("/v1/roles/%d", roleID)

	var resp struct {
		Data []Role `json:"data"`
	}
	if err := s.client.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("role not found")
	}

	return &resp.Data[0], nil
}

// CreateRole creates a new role
func (s *Service) CreateRole(ctx context.Context, req *CreateRoleRequest) (*Role, error) {
	path := "/v1/roles"

	var resp CreateRoleResponse
	if err := s.client.Post(ctx, path, req, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no role returned")
	}

	return &resp.Data[0], nil
}

// UpdateRole updates a role
func (s *Service) UpdateRole(ctx context.Context, roleID int64, req *PatchRoleRequest) (*Role, error) {
	path := fmt.Sprintf("/v1/roles/%d", roleID)

	var resp struct {
		Data []Role `json:"data"`
	}
	if err := s.client.Patch(ctx, path, req, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no role returned")
	}

	return &resp.Data[0], nil
}

// DeleteRole deletes a role
func (s *Service) DeleteRole(ctx context.Context, roleID int64) error {
	path := fmt.Sprintf("/v1/roles/%d", roleID)
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
