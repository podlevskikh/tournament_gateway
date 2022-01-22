package matches

import (
	"vollyemsk_tournament_gateway/models/referees"
	"vollyemsk_tournament_gateway/models/teams"
	"time"
)

type Winner string

const (
	WinnerHome  Winner = "home"
	WinnerGuest Winner = "guest"
)

type MatchResult struct {
	MatchID       int
	MatchDatetime time.Time

	HomeTeam              teams.Team
	HomePoints            int
	HomeTeamPlayers       []teams.Player
	HomeBestPlayer        teams.Player
	HomeRefereeEvaluation int

	GuestTeam              teams.Team
	GuestPoints            int
	GuestTeamPlayers       []teams.Player
	GuestBestPlayer        teams.Player
	GuestRefereeEvaluation int

	Winner     Winner
	Referee    referees.Referee
	SetResults []SetResult
}
