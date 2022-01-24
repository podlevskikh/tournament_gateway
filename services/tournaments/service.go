package tournaments

import (
	"context"
	"vollyemsk_tournament_gateway/models/tournaments"
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
