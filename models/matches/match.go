package matches

import (
	"time"
	"vollyemsk_tournament_gateway/models/teams"
)

type Match struct {
	ID          int
	Date        time.Time
	HomeTeamID  int
	GuestTeamID int
	HomeTeam    teams.Team   `gorm:"ForeignKey:HomeTeamID;AssociationForeignKey:ID"`
	GuestTeam   teams.Team   `gorm:"ForeignKey:GuestTeamID;AssociationForeignKey:ID"`
	Result      *MatchResult `gorm:"ForeignKey:MatchID;AssociationForeignKey:ID"`
	GroupAlias   string
}
