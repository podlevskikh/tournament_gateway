package groups

import (
	"context"
	"tournament_gateway/models/groups"
	"tournament_gateway/models/leagues"
	"tournament_gateway/models/seasons"
	"tournament_gateway/models/tournaments"
)

type GroupService interface {
	GetGroupsByTournamentSeasonStageLeagues(ctx context.Context, tournamentAlias, seasonAlias, stageAlias, leagueAlias string) ([]*groups.Group, error)
	GetGroupTeams(ctx context.Context, groupAlias string) ([]groups.Team, error)
	GetGroupMatches(ctx context.Context, groupAlias string) ([]groups.Match, error)
	GetGroupResults(ctx context.Context, groupAlias string) ([]groups.GroupResult, error)
	GetGroup(ctx context.Context, alias string) (*groups.Group, error)
	UpdateGroup(ctx context.Context, g groups.Group) (*groups.Group, error)
	CreateGroup(ctx context.Context, g groups.Group) (*groups.Group, error)
}

type TournamentService interface {
	GetTournament(ctx context.Context, alias string) (*tournaments.Tournament, error)
}

type SeasonService interface {
	GetSeason(ctx context.Context, alias string) (*seasons.Season, error)
}

type LeagueService interface {
	GetLeague(ctx context.Context, alias string) (*leagues.League, error)
}
