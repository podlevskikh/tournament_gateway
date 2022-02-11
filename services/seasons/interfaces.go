package seasons

import (
	"context"
	"tournament_gateway/models/seasons"
)

type Repository interface {
	GetSeasons(ctx context.Context) ([]*seasons.Season, error)
	GetSeasonByAlias(ctx context.Context, alias string) (*seasons.Season, error)
	UpdateStage(ctx context.Context, st *seasons.Stage) error
	CreateStage(ctx context.Context, st *seasons.Stage) error
	UpdateSeason(ctx context.Context, se seasons.Season) error
	GetStageByAlias(ctx context.Context, alias string) (*seasons.Stage, error)
	CreateSeason(ctx context.Context, se seasons.Season) (*seasons.Season, error)
}
