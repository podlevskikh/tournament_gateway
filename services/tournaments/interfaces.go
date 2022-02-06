package tournaments

import (
	"context"
	"tournament_gateway/models/tournaments"
)

type Repository interface {
	GetTournaments(ctx context.Context) ([]*tournaments.Tournament, error)
}
