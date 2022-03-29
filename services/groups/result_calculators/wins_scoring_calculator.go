package result_calculators

import (
	"tournament_gateway/models/groups"
)

func GetWinScoringResult(t groups.Team, ms []groups.Match) groups.TeamResult {
	winsScoring := getWinScoring(t, ms)
	return groups.TeamResult{Team: t, WinsScoringResult: &winsScoring}
}
