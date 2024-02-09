package user

import (
	"context"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/internal/httpclient"
)

const url = "https://api.mercadopago.com/users/me"

// Client contains the method to interact with the User API.
type Client interface {
	// Get get user information.
	// It is a get request to the endpoint: https://api.mercadopago.com/users/me
	Get(ctx context.Context) (*Response, error)
}

// client is the implementation of Client.
type client struct {
	cfg *config.Config
}

// NewClient returns a new User API Client.
func NewClient(c *config.Config) Client {
	return &client{
		cfg: c,
	}
}

func (c *client) Get(ctx context.Context) (*Response, error) {
	res, err := httpclient.Get[Response](ctx, c.cfg, url)
	if err != nil {
		return nil, err
	}

	return res, nil
}
