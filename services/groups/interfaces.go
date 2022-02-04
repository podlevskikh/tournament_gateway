package groups

import (
	"context"
	"vollyemsk_tournament_gateway/models/groups"
)

type Repository interface {
	GetGroupsByTournamentSeasonStageLeagues(ctx context.Context, tournamentAlias, seasonAlias, stageAlias, leagueAlias string) ([]*groups.Group, error)
	GetGroupsByTournamentSeasonStages(ctx context.Context, tournamentAlias, seasonAlias, stageAlias string) ([]*groups.Group, error)
	GetGroupWithTeams(ctx context.Context, groupAlias string) (groups.Group, error)
	GetGroupWithMatches(ctx context.Context, groupAlias string) (groups.Group, error)
	GetGroupWithGroupResultsAndMatches(ctx context.Context, groupAlias string) (groups.Group, error)
}
