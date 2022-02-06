package seasons

import (
	"context"
	"tournament_gateway/models/seasons"
)

type Service interface {
	GetSeasons(ctx context.Context) ([]*seasons.Season, error)
	GetStagesBySeasonAlias(ctx context.Context, alias string) ([]*seasons.Stage, error)
}
