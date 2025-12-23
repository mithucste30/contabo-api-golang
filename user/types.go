package user

import "time"

// User represents a user with API access
type User struct {
	UserID      string    `json:"userId"`
	TenantID    string    `json:"tenantId"`
	CustomerID  string    `json:"customerId"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	EmailVerified bool    `json:"emailVerified"`
	Enabled     bool      `json:"enabled"`
	TOTP        bool      `json:"totp"`
	Admin       bool      `json:"admin"`
	Roles       []Role    `json:"roles"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `json:"updatedDate"`
}

// Role represents a user role with permissions
type Role struct {
	RoleID      int64     `json:"roleId"`
	TenantID    string    `json:"tenantId"`
	CustomerID  string    `json:"customerId"`
	Name        string    `json:"name"`
	Admin       bool      `json:"admin"`
	AccessAllResources bool `json:"accessAllResources"`
	Type        string    `json:"type"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `json:"updatedDate"`
}

// UsersResponse represents the response for listing users
type UsersResponse struct {
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
	Data []User `json:"data"`
}

// CreateUserRequest represents the request body for creating a user
type CreateUserRequest struct {
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Email     string  `json:"email"`
	Enabled   bool    `json:"enabled"`
	Admin     bool    `json:"admin"`
	Roles     []int64 `json:"roles,omitempty"`
}

// CreateUserResponse represents the response when creating a user
type CreateUserResponse struct {
	Data []User `json:"data"`
}

// PatchUserRequest represents the request body for updating a user
type PatchUserRequest struct {
	FirstName *string  `json:"firstName,omitempty"`
	LastName  *string  `json:"lastName,omitempty"`
	Email     *string  `json:"email,omitempty"`
	Enabled   *bool    `json:"enabled,omitempty"`
	Admin     *bool    `json:"admin,omitempty"`
	Roles     []int64  `json:"roles,omitempty"`
}

// RolesResponse represents the response for listing roles
type RolesResponse struct {
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
	Data []Role `json:"data"`
}

// CreateRoleRequest represents the request body for creating a role
type CreateRoleRequest struct {
	Name               string            `json:"name"`
	Admin              bool              `json:"admin"`
	AccessAllResources bool              `json:"accessAllResources"`
	Type               string            `json:"type"` // "apiPermission" or "resourcePermission"
	Permissions        map[string]string `json:"permissions,omitempty"`
	TagIDs             []int64           `json:"tagIds,omitempty"`
}

// CreateRoleResponse represents the response when creating a role
type CreateRoleResponse struct {
	Data []Role `json:"data"`
}

// PatchRoleRequest represents the request body for updating a role
type PatchRoleRequest struct {
	Name               *string           `json:"name,omitempty"`
	Admin              *bool             `json:"admin,omitempty"`
	AccessAllResources *bool             `json:"accessAllResources,omitempty"`
	Permissions        map[string]string `json:"permissions,omitempty"`
	TagIDs             []int64           `json:"tagIds,omitempty"`
}
