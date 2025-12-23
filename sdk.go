package contabo

import (
	"github.com/mithucste30/contabo-sdk-golang/compute"
	"github.com/mithucste30/contabo-sdk-golang/dns"
	"github.com/mithucste30/contabo-sdk-golang/network"
	"github.com/mithucste30/contabo-sdk-golang/secret"
	"github.com/mithucste30/contabo-sdk-golang/storage"
	"github.com/mithucste30/contabo-sdk-golang/tag"
	"github.com/mithucste30/contabo-sdk-golang/user"
)

// SDK provides access to all Contabo API services
type SDK struct {
	Client  *Client
	Compute *compute.Service
	Storage *storage.Service
	Network *network.Service
	DNS     *dns.Service
	Secret  *secret.Service
	Tag     *tag.Service
	User    *user.Service
}

// NewSDK creates a new Contabo SDK instance with all services initialized
func NewSDK(config *Config) (*SDK, error) {
	client, err := NewClient(config)
	if err != nil {
		return nil, err
	}

	return &SDK{
		Client:  client,
		Compute: compute.NewService(client),
		Storage: storage.NewService(client),
		Network: network.NewService(client),
		DNS:     dns.NewService(client),
		Secret:  secret.NewService(client),
		Tag:     tag.NewService(client),
		User:    user.NewService(client),
	}, nil
}
