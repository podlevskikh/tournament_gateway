package leagues

import (
	"context"
	"vollyemsk_tournament_gateway/models/leagues"
)

type Service struct {
	rep Repository
}

func NewService(rep Repository) *Service {
	return &Service{rep: rep}
}

func (s *Service) GetLeagues(ctx context.Context) ([]*leagues.League, error) {
	return s.rep.GetLeagues(ctx)
}
