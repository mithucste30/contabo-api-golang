package secret

import "time"

// Secret represents a stored secret (SSH key or password)
type Secret struct {
	SecretID    int64     `json:"secretId"`
	TenantID    string    `json:"tenantId"`
	CustomerID  string    `json:"customerId"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`  // "ssh" or "password"
	Value       string    `json:"value"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `json:"updatedDate"`
}

// SecretsResponse represents the response for listing secrets
type SecretsResponse struct {
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
	Data []Secret `json:"data"`
}

// CreateSecretRequest represents the request body for creating a secret
type CreateSecretRequest struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

// CreateSecretResponse represents the response when creating a secret
type CreateSecretResponse struct {
	Data []Secret `json:"data"`
}

// PatchSecretRequest represents the request body for updating a secret
type PatchSecretRequest struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}
