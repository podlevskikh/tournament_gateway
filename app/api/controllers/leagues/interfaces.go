package leagues

import (
	"context"
	"tournament_gateway/models/leagues"
)

type Service interface {
	GetLeaguesByTournamentSeasonStage(ctx context.Context, tournamentAlias, seasonAlias, stageAlias string) ([]*leagues.League, error)
}
