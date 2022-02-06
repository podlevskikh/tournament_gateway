package players

import (
	"context"
	"tournament_gateway/models/groups"
)

type Repository interface {
	GetPlayersByTeamAndGroup(ctx context.Context, teamID int, groupAlias string) ([]*groups.TeamGroupPlayer, error)
}
