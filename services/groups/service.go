package groups

import (
	"context"
	"github.com/pkg/errors"
	"tournament_gateway/models/groups"
	"tournament_gateway/models/leagues"
	"tournament_gateway/services/groups/result_calculators"
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

func (s *Service) GetGroupTeams(ctx context.Context, groupAlias string) ([]groups.Team, error) {
	gr, err := s.rep.GetGroupWithTeams(ctx, groupAlias)
	if err != nil {
		return []groups.Team{}, errors.Wrap(err, "get group with teams")
	}
	return gr.Teams, nil
}

func (s *Service) GetGroupMatches(ctx context.Context, groupAlias string) ([]groups.Match, error) {
	gr, err := s.rep.GetGroupWithMatches(ctx, groupAlias)
	if err != nil {
		return []groups.Match{}, errors.Wrap(err, "get group with matches")
	}
	return gr.Matches, nil
}

func (s *Service) GetGroupResults(ctx context.Context, groupAlias string) ([]groups.GroupResult, error) {
	gr, err := s.rep.GetGroupWithGroupResultsAndMatches(ctx, groupAlias)
	if err != nil {
		return []groups.GroupResult{}, errors.Wrap(err, "get group with matches")
	}
	res := gr.GroupResults
	for i, r := range res {
		sr, err := getScoringCalculator(r.ScoringType)
		if err != nil {
			return []groups.GroupResult{}, errors.Wrap(err, "scoring calculator")
		}
		res[i].TeamResults = result_calculators.CalculateResults(r.Teams, gr.Matches, sr)
	}
	return res, nil
}

func (s *Service) GetGroup(ctx context.Context, alias string) (*groups.Group, error) {
	return s.rep.GetGroup(ctx, alias)
}

func (s *Service) UpdateGroup(ctx context.Context, g groups.Group) (*groups.Group, error) {
	return s.rep.UpdateGroup(ctx, g)
}

func (s *Service) CreateGroup(ctx context.Context, g groups.Group) (*groups.Group, error) {
	return s.rep.CreateGroup(ctx, g)
}

func leagueExists(ls []*leagues.League, league *leagues.League) bool {
	for _, l := range ls {
		if l.Alias == league.Alias {
			return true
		}
	}
	return false
}

func getScoringCalculator(scoringType string) (func(t groups.Team, ms []groups.Match) groups.TeamResult, error) {
	switch scoringType {
	case groups.WinsScoring:
		return result_calculators.GetWinScoringResult, nil
	case groups.PointsScoring:
		return result_calculators.GetPointScoringResult, nil
	}
	return func(t groups.Team, ms []groups.Match) groups.TeamResult { return groups.TeamResult{} }, errors.New("calculator not supported")
}
