package leagues

import (
	"context"
	"tournament_gateway/models/leagues"
)

type Repository interface {
	GetLeagues(ctx context.Context) ([]*leagues.League, error)
	GetLeague(ctx context.Context, alias string) (*leagues.League, error)
	UpdateLeague(ctx context.Context, l leagues.League) (*leagues.League, error)
	CreateLeague(ctx context.Context, l leagues.League) (*leagues.League, error)
}
