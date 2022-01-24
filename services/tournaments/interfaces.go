package tournaments

import (
	"context"
	"vollyemsk_tournament_gateway/models/tournaments"
)

type Repository interface {
	GetTournaments(ctx context.Context) ([]*tournaments.Tournament, error)
}
