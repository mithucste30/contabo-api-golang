package contabo

// PaginationMeta represents pagination metadata in API responses
type PaginationMeta struct {
	Size          int   `json:"size"`
	TotalElements int64 `json:"totalElements"`
	TotalPages    int   `json:"totalPages"`
	Number        int   `json:"number"`
}

// Links represents hypermedia links in API responses
type Links struct {
	Self     string `json:"self,omitempty"`
	First    string `json:"first,omitempty"`
	Previous string `json:"previous,omitempty"`
	Next     string `json:"next,omitempty"`
	Last     string `json:"last,omitempty"`
}

// ListResponse is a generic response structure for list endpoints
type ListResponse struct {
	Pagination *PaginationMeta `json:"_pagination,omitempty"`
	Links      *Links          `json:"_links,omitempty"`
}

// ListOptions represents common query parameters for list operations
type ListOptions struct {
	Page    int      // Page number
	Size    int      // Number of elements per page
	OrderBy []string // Ordering specifications (e.g., "name:asc")
}

// AuditResponse represents audit log entries
type AuditResponse struct {
	ID            string                 `json:"id"`
	Action        string                 `json:"action"`
	Timestamp     string                 `json:"timestamp"`
	TenantID      string                 `json:"tenantId"`
	CustomerID    string                 `json:"customerId"`
	ChangedBy     string                 `json:"changedBy"`
	Username      string                 `json:"username"`
	RequestID     string                 `json:"requestId"`
	TraceID       string                 `json:"traceId"`
	ResourceID    string                 `json:"resourceId"`
	ResourceType  string                 `json:"resourceType"`
	Changes       map[string]interface{} `json:"changes,omitempty"`
}
