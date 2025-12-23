package contabo

// Config holds the configuration for the Contabo API client
type Config struct {
	// OAuth2 credentials
	ClientID     string
	ClientSecret string
	Username     string // API User email
	Password     string // API Password

	// API endpoint URLs
	AuthURL string // OAuth2 token endpoint
	BaseURL string // API base URL

	// Optional custom HTTP client
	HTTPClient interface{} // Will be *http.Client
}

// NewConfig creates a new Config with default values
func NewConfig(clientID, clientSecret, username, password string) *Config {
	return &Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Username:     username,
		Password:     password,
		AuthURL:      "https://auth.contabo.com/auth/realms/contabo/protocol/openid-connect/token",
		BaseURL:      "https://api.contabo.com",
	}
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.ClientID == "" {
		return ErrMissingClientID
	}
	if c.ClientSecret == "" {
		return ErrMissingClientSecret
	}
	if c.Username == "" {
		return ErrMissingUsername
	}
	if c.Password == "" {
		return ErrMissingPassword
	}
	return nil
}
