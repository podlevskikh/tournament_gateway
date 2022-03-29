package result_calculators

import (
	"tournament_gateway/models/groups"
)

func GetPointScoringResult(t groups.Team, ms []groups.Match) groups.TeamResult {
	winsScoring := getWinScoring(t, ms)
	return groups.TeamResult{Team: t, PointsScoringResult: &winsScoring.PointsScoringResult}
}
