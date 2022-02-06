package groups

import (
	"time"
	"tournament_gateway/models/referees"
)

type Winner string

const (
	HomeTeam           = "home"
	GuestTeam          = "guest"
	WinnerHome  Winner = HomeTeam
	WinnerGuest Winner = GuestTeam
)

type MatchResult struct {
	MatchID       int
	MatchDatetime time.Time

	HomePoints            int
	HomeTeamPlayers       []*Player `gorm:"many2many:match_result2player;foreignKey:MatchID;joinTableForeignKey:MatchResultMatchID;associationForeignKey:ID;associationJoinTableForeignKey:PlayerID"`
	HomeBestPlayerID      int
	HomeBestPlayer        *Player `gorm:"ForeignKey:HomeBestPlayerID;AssociationForeignKey:ID"`
	HomeRefereeEvaluation int

	GuestPoints            int
	GuestTeamPlayers       []*Player `gorm:"many2many:match_result2player;foreignKey:MatchID;joinTableForeignKey:MatchResultMatchID;associationForeignKey:ID;associationJoinTableForeignKey:PlayerID"`
	GuestBestPlayerID      int
	GuestBestPlayer        *Player `gorm:"ForeignKey:GuestBestPlayerID;AssociationForeignKey:ID"`
	GuestRefereeEvaluation int

	Winner     Winner
	Referees   []*referees.Referee `gorm:"many2many:match_result2referee;foreignKey:MatchID;joinTableForeignKey:MatchResultMatchID;associationForeignKey:ID;associationJoinTableForeignKey:RefereeID"`
	SetResults []*SetResult        `gorm:"ForeignKey:ResultMatchID;AssociationForeignKey:MatchID"`
	Approved   bool
}
