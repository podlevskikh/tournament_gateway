package leagues

import (
	"context"
	"vollyemsk_tournament_gateway/models/leagues"
)

type Repository interface {
	GetLeagues(ctx context.Context) ([]*leagues.League, error)
}
