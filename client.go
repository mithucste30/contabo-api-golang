package contabo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Client is the main Contabo API client
type Client struct {
	config      *Config
	httpClient  *http.Client
	authManager *AuthManager
	baseURL     *url.URL
}

// NewClient creates a new Contabo API client
func NewClient(config *Config) (*Client, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	if config.HTTPClient != nil {
		if hc, ok := config.HTTPClient.(*http.Client); ok {
			httpClient = hc
		}
	}

	baseURL, err := url.Parse(config.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}

	authManager := NewAuthManager(config, httpClient)

	return &Client{
		config:      config,
		httpClient:  httpClient,
		authManager: authManager,
		baseURL:     baseURL,
	}, nil
}

// NewRequest creates a new HTTP request with authentication and required headers
func (c *Client) NewRequest(ctx context.Context, method, path string, body interface{}) (*http.Request, error) {
	// Parse the path
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	u := c.baseURL.ResolveReference(rel)

	var buf io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		buf = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	// Get access token
	token, err := c.authManager.GetAccessToken()
	if err != nil {
		return nil, err
	}

	// Set required headers
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("x-request-id", uuid.New().String())
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Optional trace ID from context
	if traceID, ok := ctx.Value("x-trace-id").(string); ok && traceID != "" {
		req.Header.Set("x-trace-id", traceID)
	}

	return req, nil
}

// Do executes an HTTP request and handles the response
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check for errors
	if resp.StatusCode >= 400 {
		apiErr := NewAPIError(
			resp.StatusCode,
			string(body),
			req.Header.Get("x-request-id"),
			req.Header.Get("x-trace-id"),
		)
		return resp, apiErr
	}

	// Decode response if v is provided
	if v != nil && len(body) > 0 {
		if err := json.Unmarshal(body, v); err != nil {
			return resp, fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return resp, nil
}

// Get performs a GET request
func (c *Client) Get(ctx context.Context, path string, v interface{}) error {
	req, err := c.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return err
	}

	_, err = c.Do(req, v)
	return err
}

// Post performs a POST request
func (c *Client) Post(ctx context.Context, path string, body, v interface{}) error {
	req, err := c.NewRequest(ctx, "POST", path, body)
	if err != nil {
		return err
	}

	_, err = c.Do(req, v)
	return err
}

// Put performs a PUT request
func (c *Client) Put(ctx context.Context, path string, body, v interface{}) error {
	req, err := c.NewRequest(ctx, "PUT", path, body)
	if err != nil {
		return err
	}

	_, err = c.Do(req, v)
	return err
}

// Patch performs a PATCH request
func (c *Client) Patch(ctx context.Context, path string, body, v interface{}) error {
	req, err := c.NewRequest(ctx, "PATCH", path, body)
	if err != nil {
		return err
	}

	_, err = c.Do(req, v)
	return err
}

// Delete performs a DELETE request
func (c *Client) Delete(ctx context.Context, path string) error {
	req, err := c.NewRequest(ctx, "DELETE", path, nil)
	if err != nil {
		return err
	}

	_, err = c.Do(req, nil)
	return err
}

// BuildQueryString builds a query string from ListOptions and additional parameters
func BuildQueryString(opts *ListOptions, params map[string]string) string {
	values := url.Values{}

	if opts != nil {
		if opts.Page > 0 {
			values.Set("page", fmt.Sprintf("%d", opts.Page))
		}
		if opts.Size > 0 {
			values.Set("size", fmt.Sprintf("%d", opts.Size))
		}
		if len(opts.OrderBy) > 0 {
			for _, order := range opts.OrderBy {
				values.Add("orderBy", order)
			}
		}
	}

	for k, v := range params {
		if v != "" {
			values.Set(k, v)
		}
	}

	if len(values) == 0 {
		return ""
	}

	return "?" + values.Encode()
}

// AddQueryParams adds query parameters to a path
func AddQueryParams(path string, params map[string]string) string {
	if len(params) == 0 {
		return path
	}

	values := url.Values{}
	for k, v := range params {
		if v != "" {
			values.Set(k, v)
		}
	}

	if len(values) == 0 {
		return path
	}

	if strings.Contains(path, "?") {
		return path + "&" + values.Encode()
	}
	return path + "?" + values.Encode()
}
