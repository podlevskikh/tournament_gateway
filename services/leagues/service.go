package leagues

import (
	"context"
	"tournament_gateway/models/leagues"
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

func (s *Service) GetLeague(ctx context.Context, alias string) (*leagues.League, error) {
	return s.rep.GetLeague(ctx, alias)
}

func (s *Service) UpdateLeague(ctx context.Context, l leagues.League) (*leagues.League, error) {
	return s.rep.UpdateLeague(ctx, l)
}

func (s *Service) CreateLeague(ctx context.Context, l leagues.League) (*leagues.League, error) {
	return s.rep.CreateLeague(ctx, l)
}
