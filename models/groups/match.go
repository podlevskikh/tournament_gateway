package groups

import (
	"time"
)

type Match struct {
	ID          int
	Date        time.Time
	HomeTeamID  int
	GuestTeamID int
	HomeTeam    Team         `gorm:"ForeignKey:HomeTeamID;AssociationForeignKey:ID"`
	GuestTeam   Team         `gorm:"ForeignKey:GuestTeamID;AssociationForeignKey:ID"`
	Result      *MatchResult `gorm:"ForeignKey:MatchID;AssociationForeignKey:ID"`
	GroupAlias  string
}
