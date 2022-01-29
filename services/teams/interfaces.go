package teams

import (
	"context"
	"vollyemsk_tournament_gateway/models/groups"
)

type Repository interface {
	GetTeam(ctx context.Context, teamID int) (groups.Team, error)
	GetTeamWithGroups(ctx context.Context, teamID int) (groups.Team, error)
}
