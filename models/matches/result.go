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
	HomeTeamPlayers       []*teams.Player
	HomeBestPlayer        *teams.Player
	HomeRefereeEvaluation int

	GuestPoints            int
	GuestTeamPlayers       []*teams.Player
	GuestBestPlayer        *teams.Player
	GuestRefereeEvaluation int

	Winner     Winner
	Referee    *referees.Referee
	SetResults []*SetResult
	Approved   bool
}
