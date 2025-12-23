package dns

import "time"

// Zone represents a DNS zone
type Zone struct {
	ZoneID      string    `json:"zoneId"`
	TenantID    string    `json:"tenantId"`
	CustomerID  string    `json:"customerId"`
	Name        string    `json:"name"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `json:"updatedDate"`
}

// ZonesResponse represents the response for listing DNS zones
type ZonesResponse struct {
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
	Data []Zone `json:"data"`
}

// CreateZoneRequest represents the request body for creating a DNS zone
type CreateZoneRequest struct {
	Name string `json:"name"`
}

// CreateZoneResponse represents the response when creating a DNS zone
type CreateZoneResponse struct {
	Data []Zone `json:"data"`
}

// Record represents a DNS record
type Record struct {
	RecordID string `json:"recordId"`
	TenantID string `json:"tenantId"`
	CustomerID string `json:"customerId"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Content  string `json:"content"`
	TTL      int    `json:"ttl"`
	Priority *int   `json:"priority,omitempty"`
}

// RecordsResponse represents the response for listing DNS records
type RecordsResponse struct {
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
	Data []Record `json:"data"`
}

// CreateRecordRequest represents the request body for creating a DNS record
type CreateRecordRequest struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Content  string `json:"content"`
	TTL      int    `json:"ttl"`
	Priority *int   `json:"priority,omitempty"`
}

// CreateRecordResponse represents the response when creating a DNS record
type CreateRecordResponse struct {
	Data []Record `json:"data"`
}

// PatchRecordRequest represents the request body for updating a DNS record
type PatchRecordRequest struct {
	Name     *string `json:"name,omitempty"`
	Type     *string `json:"type,omitempty"`
	Content  *string `json:"content,omitempty"`
	TTL      *int    `json:"ttl,omitempty"`
	Priority *int    `json:"priority,omitempty"`
}

// PTRRecord represents a PTR (reverse DNS) record
type PTRRecord struct {
	IPAddress string `json:"ipAddress"`
	TenantID  string `json:"tenantId"`
	CustomerID string `json:"customerId"`
	PTR       string `json:"ptr"`
}

// PTRRecordsResponse represents the response for listing PTR records
type PTRRecordsResponse struct {
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
	Data []PTRRecord `json:"data"`
}

// PatchPTRRequest represents the request body for updating a PTR record
type PatchPTRRequest struct {
	PTR string `json:"ptr"`
}
