package teams

import (
	"context"
	"vollyemsk_tournament_gateway/models/groups"
)

type Service struct {
	rep Repository
}

func NewService(rep Repository) *Service {
	return &Service{rep: rep}
}

func (s *Service) GetTeam(ctx context.Context, ID int) (groups.Team, error) {
	return s.rep.GetTeam(ctx, ID)
}

func (s *Service) GetTeamWithGroups(ctx context.Context, ID int) (groups.Team, error) {
	return s.rep.GetTeamWithGroups(ctx, ID)
}
