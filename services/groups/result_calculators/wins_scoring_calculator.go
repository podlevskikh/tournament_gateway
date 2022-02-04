package result_calculators

import (
	"sort"
	"vollyemsk_tournament_gateway/models/groups"
)

func CalculateWinScoringResults(ts []groups.Team, ms []groups.Match) []groups.TeamResult {
	res := make([]groups.TeamResult, 0, len(ts))
	for _, t := range ts {
		winsScoring := groups.WinsScoringResult{
			Wins:               *t.HandicapWins,
			MaxPossibleWins:    *t.HandicapWins,
			Points:             *t.HandicapPoints,
			MaxPossiblePoints:  *t.HandicapPoints,
			HandicapWins:       *t.HandicapWins,
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
				winsScoring.MaxPossibleWins++
				winsScoring.MaxPossiblePoints += 3
				continue
			}
			homeOrGuest := groups.HomeTeam
			if t.ID == m.GuestTeamID {
				homeOrGuest = groups.GuestTeam
			}
			wins, getPoints, winSets, loseSets, winPoints, losePoints := getWinLosePointsSet(homeOrGuest, *m.Result)

			winsScoring.Wins += wins
			winsScoring.MaxPossibleWins += wins
			winsScoring.Points += getPoints
			winsScoring.MaxPossiblePoints += getPoints
			if getPoints == 3 {
				winsScoring.ThreePointsMatches += 1
			} else if getPoints == 2 {
				winsScoring.TwoPointsMatches += 1
			} else if getPoints == 1 {
				winsScoring.OnePointMatches += 1
			} else {
				winsScoring.ZeroPointsMatches += 1
			}
			winsScoring.WinSets += winSets
			winsScoring.LoseSets += loseSets
			winsScoring.WinPoints += winPoints
			winsScoring.LosePoints += losePoints
		}
		res = append(res, groups.TeamResult{
			Team:              t,
			Place:             0,
			WinsScoringResult: &winsScoring,
		})
	}

	sort.SliceStable(res, func(i, j int) bool {
		if res[i].WinsScoringResult.Wins != res[j].WinsScoringResult.Wins {
			return res[i].WinsScoringResult.Wins > res[j].WinsScoringResult.Wins
		}
		if res[i].WinsScoringResult.Points != res[j].WinsScoringResult.Points {
			return res[i].WinsScoringResult.Points > res[j].WinsScoringResult.Points
		}
		winToLoseSetsI := getRatio(res[i].WinsScoringResult.WinSets, res[i].WinsScoringResult.LoseSets)
		winToLoseSetsJ := getRatio(res[j].WinsScoringResult.WinSets, res[j].WinsScoringResult.LoseSets)
		if winToLoseSetsI != winToLoseSetsJ {
			return winToLoseSetsI > winToLoseSetsJ
		}
		winToLosePointsI := getRatio(res[i].WinsScoringResult.WinPoints, res[i].WinsScoringResult.LosePoints)
		winToLosePointsJ := getRatio(res[j].WinsScoringResult.WinPoints, res[j].WinsScoringResult.LosePoints)
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
