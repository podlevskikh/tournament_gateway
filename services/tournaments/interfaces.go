package tournaments

import (
	"context"
	"tournament_gateway/models/tournaments"
)

type Repository interface {
	GetTournaments(ctx context.Context) ([]*tournaments.Tournament, error)
	GetTournament(ctx context.Context, alias string) (*tournaments.Tournament, error)
	UpdateTournament(ctx context.Context, t tournaments.Tournament) (*tournaments.Tournament, error)
}
