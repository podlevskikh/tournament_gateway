package matches

import (
	"time"
)

type Match struct {
	ID          int
	Date        time.Time
	HomeTeamId  int
	GuestTeamId int
	Result      *MatchResult
}
