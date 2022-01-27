package groups

import (
	"context"
	"vollyemsk_tournament_gateway/models/groups"
)

type Repository interface {
	GetGroupsByTournamentSeasonStageLeagues(ctx context.Context, tournamentAlias, seasonAlias, stageAlias, leagueAlias string) ([]*groups.Group, error)
	GetGroupsByTournamentSeasonStages(ctx context.Context, tournamentAlias, seasonAlias, stageAlias string) ([]*groups.Group, error)
}
