package matches

import (
	"context"
	"vollyemsk_tournament_gateway/models/groups"
)

type Repository interface {
	GetMatch(ctx context.Context, matchID int) (groups.Match, error)
}
