package tournaments

import (
	"context"
	"tournament_gateway/models/tournaments"
)

type Service struct {
	rep Repository
}

func NewService(rep Repository) *Service {
	return &Service{rep: rep}
}

func (s *Service) GetTournaments(ctx context.Context) ([]*tournaments.Tournament, error) {
	return s.rep.GetTournaments(ctx)
}

func (s *Service) GetTournament(ctx context.Context, alias string) (*tournaments.Tournament, error) {
	return s.rep.GetTournament(ctx, alias)
}

func (s *Service) UpdateTournament(ctx context.Context, t tournaments.Tournament) (*tournaments.Tournament, error) {
	return s.rep.UpdateTournament(ctx, t)
}
