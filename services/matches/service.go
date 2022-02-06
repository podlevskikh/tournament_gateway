package matches

import (
	"context"
	"tournament_gateway/models/groups"
)

type Service struct {
	rep Repository
}

func NewService(rep Repository) *Service {
	return &Service{rep: rep}
}

func (s *Service) GetMatch(ctx context.Context, matchID int) (groups.Match, error) {
	return s.rep.GetMatch(ctx, matchID)
}
