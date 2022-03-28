package sql

import (
	"context"

	"github.com/ory/hydra/client"
	"github.com/ory/x/sqlcon"
)

func (p *Persister) CreateClientSecret(ctx context.Context, c *client.ClientSecret) error {
	return sqlcon.HandleError(p.Connection(ctx).Create(c, ""))
}

func (p *Persister) GetClientSecret(ctx context.Context, clientID string) (*client.ClientSecret, error) {
	var cl client.ClientSecret
	return &cl, sqlcon.HandleError(p.Connection(ctx).Where("client_id = ?", clientID).First(&cl))
}
