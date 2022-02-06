package result_calculators

import (
	"sort"
	"tournament_gateway/models/groups"
)

func CalculatePointScoringResults(ts []groups.Team, ms []groups.Match) []groups.TeamResult {
	res := make([]groups.TeamResult, 0, len(ts))
	for _, t := range ts {
		pointResult := groups.PointsScoringResult{
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
		}
		for _, m := range ms {
			if m.HomeTeamID != t.ID && m.GuestTeamID != t.ID {
				continue
			}
			if m.Result == nil {
				pointResult.MaxPossiblePoints += 3
				continue
			}
			homeOrGuest := groups.HomeTeam
			if t.ID == m.GuestTeamID {
				homeOrGuest = groups.GuestTeam
			}
			_, getPoints, winSets, loseSets, winPoints, losePoints := getWinLosePointsSet(homeOrGuest, *m.Result)

			pointResult.Points += getPoints
			pointResult.MaxPossiblePoints += getPoints
			if getPoints == 3 {
				pointResult.ThreePointsMatches += 1
			} else if getPoints == 2 {
				pointResult.TwoPointsMatches += 1
			} else if getPoints == 1 {
				pointResult.OnePointMatches += 1
			} else {
				pointResult.ZeroPointsMatches += 1
			}
			pointResult.WinSets += winSets
			pointResult.LoseSets += loseSets
			pointResult.WinPoints += winPoints
			pointResult.LosePoints += losePoints
		}

		res = append(res, groups.TeamResult{
			Team:                t,
			Place:               0,
			PointsScoringResult: &pointResult,
		})
	}

	sort.SliceStable(res, func(i, j int) bool {
		if res[i].PointsScoringResult.Points != res[j].PointsScoringResult.Points {
			return res[i].PointsScoringResult.Points > res[j].PointsScoringResult.Points
		}
		winToLoseSetsI := getRatio(res[i].PointsScoringResult.WinSets, res[i].PointsScoringResult.LoseSets)
		winToLoseSetsJ := getRatio(res[j].PointsScoringResult.WinSets, res[j].PointsScoringResult.LoseSets)
		if winToLoseSetsI != winToLoseSetsJ {
			return winToLoseSetsI > winToLoseSetsJ
		}
		winToLosePointsI := getRatio(res[i].PointsScoringResult.WinPoints, res[i].PointsScoringResult.LosePoints)
		winToLosePointsJ := getRatio(res[j].PointsScoringResult.WinPoints, res[j].PointsScoringResult.LosePoints)
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
