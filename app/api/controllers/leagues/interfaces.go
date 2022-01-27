package leagues

import (
	"context"
	"vollyemsk_tournament_gateway/models/leagues"
)

type Service interface {
	GetLeaguesByTournamentSeasonStage(ctx context.Context, tournamentAlias, seasonAlias, stageAlias string) ([]*leagues.League, error)
}
