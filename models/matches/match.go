package matches

import (
	"time"
	"vollyemsk_tournament_gateway/models/teams"
)

type Match struct {
	ID        int
	Date      time.Time
	HomeTeam  teams.Team
	GuestTeam teams.Team
	Result    *MatchResult
}
