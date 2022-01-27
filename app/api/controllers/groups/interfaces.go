package groups

import (
	"context"
	"vollyemsk_tournament_gateway/models/groups"
)

type Service interface {
	GetGroupsByTournamentSeasonStageLeagues(ctx context.Context, tournamentAlias, seasonAlias, stageAlias, leagueAlias string) ([]*groups.Group, error)
}
