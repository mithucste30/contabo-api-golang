package tag

import "time"

// Tag represents a resource tag
type Tag struct {
	TagID      int64     `json:"tagId"`
	TenantID   string    `json:"tenantId"`
	CustomerID string    `json:"customerId"`
	Name       string    `json:"name"`
	Color      string    `json:"color,omitempty"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `json:"updatedDate"`
}

// TagsResponse represents the response for listing tags
type TagsResponse struct {
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
	Data []Tag `json:"data"`
}

// CreateTagRequest represents the request body for creating a tag
type CreateTagRequest struct {
	Name  string `json:"name"`
	Color string `json:"color,omitempty"`
}

// CreateTagResponse represents the response when creating a tag
type CreateTagResponse struct {
	Data []Tag `json:"data"`
}

// PatchTagRequest represents the request body for updating a tag
type PatchTagRequest struct {
	Name  *string `json:"name,omitempty"`
	Color *string `json:"color,omitempty"`
}
