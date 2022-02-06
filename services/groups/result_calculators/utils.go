package result_calculators

import "tournament_gateway/models/groups"

func getWinLosePointsSet(homeOrGuest string, mr groups.MatchResult) (wins, getPoints, winSets, loseSets, winPoints, losePoints int) {
	for _, s := range mr.SetResults {
		if homeOrGuest == groups.HomeTeam {
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
	} else if loseSets == 2 {
		getPoints = 2
		wins = 1
	} else if winSets == 2 {
		getPoints = 1
	}

	return wins, getPoints, winSets, loseSets, winPoints, losePoints
}

func getRatio(wins, loses int) float64 {
	if loses == 0 {
		return float64(wins) * 1000
	}
	return float64(wins) / float64(loses)
}
