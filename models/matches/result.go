package matches

import (
	"time"
	"vollyemsk_tournament_gateway/models/referees"
	"vollyemsk_tournament_gateway/models/teams"
)

type Winner string

const (
	WinnerHome  Winner = "home"
	WinnerGuest Winner = "guest"
)

type MatchResult struct {
	MatchID       int
	MatchDatetime time.Time

	HomePoints            int
	HomeTeamPlayers       []*teams.Player `gorm:"many2many:match_result2player;foreignKey:MatchID;joinTableForeignKey:MatchID;associationForeignKey:ID;associationJoinTableForeignKey:PlayerID"`
	HomeBestPlayerID      int
	HomeBestPlayer        *teams.Player `gorm:"ForeignKey:HomeBestPlayerID;AssociationForeignKey:ID"`
	HomeRefereeEvaluation int

	GuestPoints            int
	GuestTeamPlayers       []*teams.Player `gorm:"many2many:match_result2player;foreignKey:MatchID;joinTableForeignKey:MatchID;associationForeignKey:ID;associationJoinTableForeignKey:PlayerID"`
	GuestBestPlayerID      int
	GuestBestPlayer        *teams.Player `gorm:"ForeignKey:GuestBestPlayerID;AssociationForeignKey:ID"`
	GuestRefereeEvaluation int

	Winner     Winner
	RefereeID  int
	Referee    *referees.Referee `gorm:"ForeignKey:RefereeID;AssociationForeignKey:ID"`
	SetResults []*SetResult      `gorm:"ForeignKey:MatchID;AssociationForeignKey:MatchID"`
	Approved   bool
}
