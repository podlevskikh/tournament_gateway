package groups

import (
	"context"
	"github.com/pkg/errors"
	"vollyemsk_tournament_gateway/models/groups"
	leagues "vollyemsk_tournament_gateway/models/leagues"
)

type Service struct {
	rep Repository
}

func NewService(rep Repository) *Service {
	return &Service{rep: rep}
}

func (s *Service) GetGroupsByTournamentSeasonStageLeagues(ctx context.Context, tournamentAlias, seasonAlias, stageAlias, leagueAlias string) ([]*groups.Group, error) {
	return s.rep.GetGroupsByTournamentSeasonStageLeagues(ctx, tournamentAlias, seasonAlias, stageAlias, leagueAlias)
}

func (s *Service) GetLeaguesByTournamentSeasonStage(ctx context.Context, tournamentAlias, seasonAlias, stageAlias string) ([]*leagues.League, error) {
	gr, err := s.rep.GetGroupsByTournamentSeasonStages(ctx, tournamentAlias, seasonAlias, stageAlias)
	if err != nil {
		return nil, errors.Wrap(err, "get groups by tournament, season, stage")
	}

	ls := []*leagues.League{}
	for _, g := range gr {
		if !leagueExists(ls, g.League) {
			ls = append(ls, g.League)
		}
	}
	return ls, nil
}

func leagueExists(ls []*leagues.League, league *leagues.League) bool {
	for _, l :=range ls {
		if l.Alias == league.Alias {
			return true
		}
	}
	return false
}
