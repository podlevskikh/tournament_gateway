package matches

import (
	"context"
	"tournament_gateway/models/groups"
)

type Repository interface {
	GetMatch(ctx context.Context, matchID int) (groups.Match, error)
}
