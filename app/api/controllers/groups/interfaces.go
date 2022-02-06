package groups

import (
	"context"
	"tournament_gateway/models/groups"
)

type Service interface {
	GetGroupsByTournamentSeasonStageLeagues(ctx context.Context, tournamentAlias, seasonAlias, stageAlias, leagueAlias string) ([]*groups.Group, error)
	GetGroupTeams(ctx context.Context, groupAlias string) ([]groups.Team, error)
	GetGroupMatches(ctx context.Context, groupAlias string) ([]groups.Match, error)
	GetGroupResults(ctx context.Context, groupAlias string) ([]groups.GroupResult, error)
}
