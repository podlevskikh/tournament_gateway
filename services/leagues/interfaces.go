package leagues

import (
	"context"
	"tournament_gateway/models/leagues"
)

type Repository interface {
	GetLeagues(ctx context.Context) ([]*leagues.League, error)
}
