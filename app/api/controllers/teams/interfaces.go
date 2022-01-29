package teams

import (
	"context"
	"vollyemsk_tournament_gateway/models/groups"
)

type TeamsService interface {
	GetTeam(ctx context.Context, ID int) (groups.Team, error)
	GetTeamWithGroups(ctx context.Context, ID int) (groups.Team, error)
}

type PlayersService interface {
	GetPlayersByTeamAndGroup(ctx context.Context, teamID int, groupAlias string) ([]*groups.Player, error)
}
