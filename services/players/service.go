package players

import (
	"context"
	"github.com/pkg/errors"
	"vollyemsk_tournament_gateway/models/groups"
)

type Service struct {
	rep Repository
}

func NewService(rep Repository) *Service {
	return &Service{rep: rep}
}

func (s *Service) GetPlayersByTeamAndGroup(ctx context.Context, teamID int, groupAlias string) ([]*groups.Player, error) {
	tgps, err := s.rep.GetPlayersByTeamAndGroup(ctx, teamID, groupAlias)
	if err != nil {
		return nil, errors.Wrap(err, "get team group players")
	}
	ps := make([]*groups.Player, 0, len(tgps))
	for _, tgp := range tgps {
		ps = append(ps, &tgp.Player)
	}
	return ps, nil
}
