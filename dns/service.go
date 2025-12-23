package dns

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

// Service handles DNS-related API operations
type Service struct {
	client Client
}

// NewService creates a new DNS service
func NewService(client Client) *Service {
	return &Service{client: client}
}

// Zones

// ListZones retrieves a list of DNS zones
func (s *Service) ListZones(ctx context.Context, opts *ListOptions) (*ZonesResponse, error) {
	path := "/v1/dns/zones"
	if opts != nil {
		path += buildQueryString(opts, nil)
	}

	var resp ZonesResponse
	if err := s.client.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetZone retrieves a specific DNS zone by name
func (s *Service) GetZone(ctx context.Context, zoneName string) (*Zone, error) {
	path := fmt.Sprintf("/v1/dns/zones/%s", zoneName)

	var resp struct {
		Data []Zone `json:"data"`
	}
	if err := s.client.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("zone not found")
	}

	return &resp.Data[0], nil
}

// CreateZone creates a new DNS zone
func (s *Service) CreateZone(ctx context.Context, req *CreateZoneRequest) (*Zone, error) {
	path := "/v1/dns/zones"

	var resp CreateZoneResponse
	if err := s.client.Post(ctx, path, req, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no zone returned")
	}

	return &resp.Data[0], nil
}

// DeleteZone deletes a DNS zone
func (s *Service) DeleteZone(ctx context.Context, zoneName string) error {
	path := fmt.Sprintf("/v1/dns/zones/%s", zoneName)
	return s.client.Delete(ctx, path)
}

// Records

// ListRecords retrieves DNS records for a zone
func (s *Service) ListRecords(ctx context.Context, zoneName string, opts *ListOptions) (*RecordsResponse, error) {
	path := fmt.Sprintf("/v1/dns/zones/%s/records", zoneName)
	if opts != nil {
		path += buildQueryString(opts, nil)
	}

	var resp RecordsResponse
	if err := s.client.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetRecord retrieves a specific DNS record
func (s *Service) GetRecord(ctx context.Context, zoneName, recordID string) (*Record, error) {
	path := fmt.Sprintf("/v1/dns/zones/%s/records/%s", zoneName, recordID)

	var resp struct {
		Data []Record `json:"data"`
	}
	if err := s.client.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("record not found")
	}

	return &resp.Data[0], nil
}

// CreateRecord creates a new DNS record
func (s *Service) CreateRecord(ctx context.Context, zoneName string, req *CreateRecordRequest) (*Record, error) {
	path := fmt.Sprintf("/v1/dns/zones/%s/records", zoneName)

	var resp CreateRecordResponse
	if err := s.client.Post(ctx, path, req, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no record returned")
	}

	return &resp.Data[0], nil
}

// UpdateRecord updates a DNS record
func (s *Service) UpdateRecord(ctx context.Context, zoneName, recordID string, req *PatchRecordRequest) (*Record, error) {
	path := fmt.Sprintf("/v1/dns/zones/%s/records/%s", zoneName, recordID)

	var resp struct {
		Data []Record `json:"data"`
	}
	if err := s.client.Patch(ctx, path, req, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no record returned")
	}

	return &resp.Data[0], nil
}

// DeleteRecord deletes a DNS record
func (s *Service) DeleteRecord(ctx context.Context, zoneName, recordID string) error {
	path := fmt.Sprintf("/v1/dns/zones/%s/records/%s", zoneName, recordID)
	return s.client.Delete(ctx, path)
}

// PTR Records

// ListPTRRecords retrieves a list of PTR records
func (s *Service) ListPTRRecords(ctx context.Context, opts *ListOptions) (*PTRRecordsResponse, error) {
	path := "/v1/dns/ptrs"
	if opts != nil {
		path += buildQueryString(opts, nil)
	}

	var resp PTRRecordsResponse
	if err := s.client.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetPTRRecord retrieves a specific PTR record by IP address
func (s *Service) GetPTRRecord(ctx context.Context, ipAddress string) (*PTRRecord, error) {
	path := fmt.Sprintf("/v1/dns/ptrs/%s", ipAddress)

	var resp struct {
		Data []PTRRecord `json:"data"`
	}
	if err := s.client.Get(ctx, path, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("PTR record not found")
	}

	return &resp.Data[0], nil
}

// UpdatePTRRecord updates a PTR record
func (s *Service) UpdatePTRRecord(ctx context.Context, ipAddress string, req *PatchPTRRequest) (*PTRRecord, error) {
	path := fmt.Sprintf("/v1/dns/ptrs/%s", ipAddress)

	var resp struct {
		Data []PTRRecord `json:"data"`
	}
	if err := s.client.Patch(ctx, path, req, &resp); err != nil {
		return nil, err
	}

	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("no PTR record returned")
	}

	return &resp.Data[0], nil
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
