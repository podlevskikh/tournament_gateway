package matches

import (
	"context"
	"vollyemsk_tournament_gateway/models/groups"
)

type Service interface {
	GetMatch(ctx context.Context, id int) (groups.Match, error)
}
