package players

import (
	"context"
	"vollyemsk_tournament_gateway/models/groups"
)

type Repository interface {
	GetPlayersByTeamAndGroup(ctx context.Context, teamID int, groupAlias string) ([]*groups.TeamGroupPlayer, error)
}
