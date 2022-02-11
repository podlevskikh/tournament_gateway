package seasons

import (
	"context"
	"github.com/pkg/errors"
	"tournament_gateway/models/seasons"
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

func (s *Service) GetSeason(ctx context.Context, alias string) (*seasons.Season, error) {
	season, err := s.rep.GetSeasonByAlias(ctx, alias)
	if err != nil {
		return nil, err
	}
	return season, nil
}

func (s *Service) UpdateSeason(ctx context.Context, se seasons.Season) (*seasons.Season, error) {
	season, err := s.rep.GetSeasonByAlias(ctx, se.Alias)
	if err != nil {
		return nil, errors.Wrap(err, "get season")
	}
	for _, st := range season.Stages {
		if !stageExists(st, se.Stages) {
			return nil, errors.New("delete stage deprecated")
		}
	}

	for _, st := range se.Stages {
		if stageExists(st, season.Stages) {
			err = s.rep.UpdateStage(ctx, st)
			if err != nil {
				return nil, errors.Wrap(err, "update stage")
			}
		} else {
			err = s.rep.CreateStage(ctx, st)
			if err != nil {
				return nil, errors.Wrap(err, "create stage")
			}
		}
	}

	err = s.rep.UpdateSeason(ctx, se)
	if err != nil {
		return nil, errors.Wrap(err, "update season")
	}

	return &se, nil
}

func stageExists(s *seasons.Stage, ss []*seasons.Stage) bool {
	for _, st := range ss {
		if st.Alias == s.Alias {
			return true
		}
	}
	return false
}

func (s *Service) CreateSeason(ctx context.Context, se seasons.Season) (*seasons.Season, error) {
	_, err := s.rep.GetSeasonByAlias(ctx, se.Alias)
	if err == nil {
		return nil, errors.Wrap(err, "season exists")
	}
	//todo check err not found
	for _, st := range se.Stages {
		_, err = s.rep.GetStageByAlias(ctx, st.Alias)
		if err == nil {
			return nil, errors.Wrap(err, "stage exists")
		}
		//todo check err not found
	}

	return s.rep.CreateSeason(ctx, se)
}
