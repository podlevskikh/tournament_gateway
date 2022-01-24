package seasons

import (
	"context"
	"vollyemsk_tournament_gateway/models/seasons"
)

type Repository interface {
	GetSeasons(ctx context.Context) ([]*seasons.Season, error)
	GetSeasonByAlias(ctx context.Context, alias string) (*seasons.Season, error)
}
