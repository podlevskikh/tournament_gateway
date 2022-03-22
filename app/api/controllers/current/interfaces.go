package current

import (
	"context"
	"tournament_gateway/models/groups"
	"tournament_gateway/models/leagues"
	"tournament_gateway/models/seasons"
	"tournament_gateway/models/tournaments"
)

type GroupService interface {
	GetLeaguesByTournamentSeasonStage(ctx context.Context, tournamentAlias, seasonAlias, stageAlias string) ([]*leagues.League, error)
	GetGroupsByTournamentSeasonStageLeagues(ctx context.Context, tournamentAlias, seasonAlias, stageAlias, leagueAlias string) ([]*groups.Group, error)
}

type SeasonService interface {
	GetSeasons(ctx context.Context) ([]*seasons.Season, error)
	GetStagesBySeasonAlias(ctx context.Context, alias string) ([]*seasons.Stage, error)
}

type TournamentService interface {
	GetTournaments(ctx context.Context) ([]*tournaments.Tournament, error)
}