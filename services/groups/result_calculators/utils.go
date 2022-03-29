package result_calculators

import (
	"sort"
	"tournament_gateway/models/groups"
)

func CalculateResults(ts []groups.Team, ms []groups.Match, f func (groups.Team, []groups.Match) groups.TeamResult) []groups.TeamResult {
	res := make([]groups.TeamResult, 0, len(ts))
	for _, t := range ts {
		res = append(res, f(t, ms))
	}
	return sortAndSetPlaces(res)
}


func appendResults(winsScoring *groups.WinsScoringResult, t groups.Team, m groups.Match) {
	if m.HomeTeamID != t.ID && m.GuestTeamID != t.ID {
		return
	}
	if m.Result == nil {
		winsScoring.MaxPossibleWins++
		winsScoring.MaxPossiblePoints += 3
		return
	}

	var wins, getPoints, threePointsMatches, twoPointsMatches, onePointMatches, zeroPointsMatches, winSets, loseSets, winPoints, losePoints int
	for _, s := range m.Result.SetResults {
		if t.ID == m.HomeTeamID {
			winPoints += s.HomeScore
			losePoints += s.GuestScore
			if s.HomeScore > s.GuestScore {
				winSets++
			} else {
				loseSets++
			}
		} else {
			winPoints += s.GuestScore
			losePoints += s.HomeScore
			if s.HomeScore < s.GuestScore {
				winSets++
			} else {
				loseSets++
			}
		}
	}

	if loseSets <= 1 {
		getPoints = 3
		wins = 1
		threePointsMatches = 1
	} else if loseSets == 2 {
		getPoints = 2
		wins = 1
		twoPointsMatches = 1
	} else if winSets == 2 {
		getPoints = 1
		onePointMatches = 1
	} else {
		zeroPointsMatches = 1
	}

	winsScoring.Wins += wins
	winsScoring.MaxPossibleWins += wins
	winsScoring.Points += getPoints
	winsScoring.MaxPossiblePoints += getPoints
	winsScoring.ThreePointsMatches += threePointsMatches
	winsScoring.TwoPointsMatches += twoPointsMatches
	winsScoring.OnePointMatches += onePointMatches
	winsScoring.ZeroPointsMatches += zeroPointsMatches
	winsScoring.WinSets += winSets
	winsScoring.LoseSets += loseSets
	winsScoring.WinPoints += winPoints
	winsScoring.LosePoints += losePoints

	return
}

func getRatio(wins, loses int) float64 {
	if loses == 0 {
		return float64(wins) * 1000
	}
	return float64(wins) / float64(loses)
}

func getWinScoring(t groups.Team, ms []groups.Match) groups.WinsScoringResult {
	winsScoring := groups.WinsScoringResult{
		Wins:            *t.HandicapWins,
		MaxPossibleWins: *t.HandicapWins,
		HandicapWins:    *t.HandicapWins,
		PointsScoringResult: groups.PointsScoringResult{
			Points:             *t.HandicapPoints,
			MaxPossiblePoints:  *t.HandicapPoints,
			HandicapPoints:     *t.HandicapPoints,
			ThreePointsMatches: 0,
			TwoPointsMatches:   0,
			OnePointMatches:    0,
			ZeroPointsMatches:  0,
			WinSets:            0,
			LoseSets:           0,
			WinPoints:          0,
			LosePoints:         0,
		},
	}
	for _, m := range ms {
		appendResults(&winsScoring, t, m)
	}

	return winsScoring
}

func sortAndSetPlaces(res []groups.TeamResult) []groups.TeamResult {
	sort.SliceStable(res, func(i, j int) bool {
		var iPointScoringResult, jPointScoringResult groups.PointsScoringResult
		if res[i].WinsScoringResult != nil {
			if res[i].WinsScoringResult.Wins != res[j].WinsScoringResult.Wins {
				return res[i].WinsScoringResult.Wins > res[j].WinsScoringResult.Wins
			}
			iPointScoringResult = res[i].WinsScoringResult.PointsScoringResult
			jPointScoringResult = res[j].WinsScoringResult.PointsScoringResult
		} else {
			iPointScoringResult = *res[i].PointsScoringResult
			jPointScoringResult = *res[j].PointsScoringResult
		}
		if iPointScoringResult.Points != jPointScoringResult.Points {
			return iPointScoringResult.Points > jPointScoringResult.Points
		}
		winToLoseSetsI := getRatio(iPointScoringResult.WinSets, iPointScoringResult.LoseSets)
		winToLoseSetsJ := getRatio(jPointScoringResult.WinSets, jPointScoringResult.LoseSets)
		if winToLoseSetsI != winToLoseSetsJ {
			return winToLoseSetsI > winToLoseSetsJ
		}
		winToLosePointsI := getRatio(iPointScoringResult.WinPoints, iPointScoringResult.LosePoints)
		winToLosePointsJ := getRatio(jPointScoringResult.WinPoints, jPointScoringResult.LosePoints)
		if winToLosePointsI != winToLosePointsJ {
			return winToLosePointsI > winToLosePointsJ
		}

		return res[i].Team.Name < res[j].Team.Name
	})

	for i := range res {
		res[i].Place = i + 1
	}
	return res
}
