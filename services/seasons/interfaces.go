package seasons

import (
	"context"
	"tournament_gateway/models/seasons"
)

type Repository interface {
	GetSeasons(ctx context.Context) ([]*seasons.Season, error)
	GetSeasonByAlias(ctx context.Context, alias string) (*seasons.Season, error)
}
