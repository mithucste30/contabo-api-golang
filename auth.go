package contabo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// TokenResponse represents the OAuth2 token response
type TokenResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	SessionState     string `json:"session_state"`
	Scope            string `json:"scope"`
}

// AuthManager handles OAuth2 authentication and token management
type AuthManager struct {
	config      *Config
	httpClient  *http.Client
	token       *TokenResponse
	tokenExpiry time.Time
	mu          sync.RWMutex
}

// NewAuthManager creates a new authentication manager
func NewAuthManager(config *Config, httpClient *http.Client) *AuthManager {
	return &AuthManager{
		config:     config,
		httpClient: httpClient,
	}
}

// GetAccessToken returns a valid access token, refreshing if necessary
func (a *AuthManager) GetAccessToken() (string, error) {
	a.mu.RLock()
	// Check if we have a valid token
	if a.token != nil && time.Now().Before(a.tokenExpiry) {
		token := a.token.AccessToken
		a.mu.RUnlock()
		return token, nil
	}
	a.mu.RUnlock()

	// Token is expired or doesn't exist, acquire new token
	a.mu.Lock()
	defer a.mu.Unlock()

	// Double-check after acquiring write lock
	if a.token != nil && time.Now().Before(a.tokenExpiry) {
		return a.token.AccessToken, nil
	}

	// Request new token
	return a.authenticate()
}

// authenticate performs the OAuth2 password grant flow
func (a *AuthManager) authenticate() (string, error) {
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("client_id", a.config.ClientID)
	data.Set("client_secret", a.config.ClientSecret)
	data.Set("username", a.config.Username)
	data.Set("password", a.config.Password)

	req, err := http.NewRequest("POST", a.config.AuthURL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to create auth request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("authentication request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read auth response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%w: status %d, body: %s", ErrAuthenticationFailed, resp.StatusCode, string(body))
	}

	var tokenResp TokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return "", fmt.Errorf("failed to parse token response: %w", err)
	}

	// Store token and calculate expiry (subtract 60 seconds as buffer)
	a.token = &tokenResp
	a.tokenExpiry = time.Now().Add(time.Duration(tokenResp.ExpiresIn-60) * time.Second)

	return tokenResp.AccessToken, nil
}

// Refresh forces a token refresh
func (a *AuthManager) Refresh() error {
	a.mu.Lock()
	defer a.mu.Unlock()

	_, err := a.authenticate()
	return err
}
