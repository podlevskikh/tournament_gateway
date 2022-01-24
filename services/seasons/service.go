package seasons

import (
	"context"
	"vollyemsk_tournament_gateway/models/seasons"
)

type Service struct {
	rep Repository
}

func NewService(rep Repository) *Service {
	return &Service{rep: rep}
}

func (s *Service) GetSeasons(ctx context.Context) ([]*seasons.Season, error) {
	return s.rep.GetSeasons(ctx)
}

func (s *Service) GetStagesBySeasonAlias(ctx context.Context, alias string) ([]*seasons.Stage, error) {
	season, err := s.rep.GetSeasonByAlias(ctx, alias)
	if err != nil {
		return nil, err
	}
	return season.Stages, nil
}
