package tournaments

import (
	"context"
	"tournament_gateway/models/tournaments"
)

type Service interface {
	GetTournaments(ctx context.Context) ([]*tournaments.Tournament, error)
}
