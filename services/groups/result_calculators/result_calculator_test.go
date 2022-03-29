package result_calculators

import (
	"github.com/go-playground/assert/v2"
	"testing"
	"tournament_gateway/models/groups"
)

func TestGetWinLosePointSet(t *testing.T) {
	tests := []struct {
		name                       string
		team                       groups.Team
		mr                         groups.Match
		expectedWins               int
		expectedMaxPossibleWins    int
		expectedPoints             int
		expectedMaxPossiblePoints  int
		expectedThreePointsMatches int
		expectedTwoPointsMatches   int
		expectedOnePointMatches    int
		expectedZeroPointsMatches  int
		expectedWinSets            int
		expectedLoseSets           int
		expectedWinPoints          int
		expectedLosePoints         int
	}{
		{
			name: "not suitable match",
			team: groups.Team{
				ID: 1,
			},
			mr: groups.Match{
				Result: &groups.MatchResult{
					SetResults: []*groups.SetResult{
						{HomeScore: 25, GuestScore: 21},
						{HomeScore: 26, GuestScore: 28},
						{HomeScore: 15, GuestScore: 25},
						{HomeScore: 25, GuestScore: 0},
						{HomeScore: 15, GuestScore: 10},
					},
				},
			},
			expectedWins:               0,
			expectedMaxPossibleWins:    0,
			expectedPoints:             0,
			expectedMaxPossiblePoints:  0,
			expectedThreePointsMatches: 0,
			expectedTwoPointsMatches:   0,
			expectedOnePointMatches:    0,
			expectedZeroPointsMatches:  0,
			expectedWinSets:            0,
			expectedLoseSets:           0,
			expectedWinPoints:          0,
			expectedLosePoints:         0,
		},
		{
			name: "empty results",
			team: groups.Team{
				ID: 1,
			},
			mr: groups.Match{
				GuestTeamID: 1,
				Result:      nil,
			},
			expectedWins:               0,
			expectedMaxPossibleWins:    1,
			expectedPoints:             0,
			expectedMaxPossiblePoints:  3,
			expectedThreePointsMatches: 0,
			expectedTwoPointsMatches:   0,
			expectedOnePointMatches:    0,
			expectedZeroPointsMatches:  0,
			expectedWinSets:            0,
			expectedLoseSets:           0,
			expectedWinPoints:          0,
			expectedLosePoints:         0,
		},
		{
			name: "home win 3:2",
			team: groups.Team{ID: 1},
			mr: groups.Match{
				HomeTeamID: 1,
				Result: &groups.MatchResult{
					SetResults: []*groups.SetResult{
						{HomeScore: 25, GuestScore: 21},
						{HomeScore: 26, GuestScore: 28},
						{HomeScore: 15, GuestScore: 25},
						{HomeScore: 25, GuestScore: 0},
						{HomeScore: 15, GuestScore: 10},
					},
				},
			},
			expectedWins:               1,
			expectedMaxPossibleWins:    1,
			expectedPoints:             2,
			expectedMaxPossiblePoints:  2,
			expectedThreePointsMatches: 0,
			expectedTwoPointsMatches:   1,
			expectedOnePointMatches:    0,
			expectedZeroPointsMatches:  0,
			expectedWinSets:            3,
			expectedLoseSets:           2,
			expectedWinPoints:          106,
			expectedLosePoints:         84,
		},
		{
			name: "guest loose 1:3",
			team: groups.Team{ID: 1},
			mr: groups.Match{
				GuestTeamID: 1,
				Result: &groups.MatchResult{
					SetResults: []*groups.SetResult{
						{HomeScore: 25, GuestScore: 21},
						{HomeScore: 26, GuestScore: 28},
						{HomeScore: 27, GuestScore: 25},
						{HomeScore: 25, GuestScore: 0},
					},
				},
			},
			expectedWins:               0,
			expectedMaxPossibleWins:    0,
			expectedPoints:             0,
			expectedMaxPossiblePoints:  0,
			expectedThreePointsMatches: 0,
			expectedTwoPointsMatches:   0,
			expectedOnePointMatches:    0,
			expectedZeroPointsMatches:  1,
			expectedWinSets:            1,
			expectedLoseSets:           3,
			expectedWinPoints:          74,
			expectedLosePoints:         103,
		},
		{
			name: "guest lose 2:3",
			team: groups.Team{ID: 1},
			mr: groups.Match{
				GuestTeamID: 1,
				Result: &groups.MatchResult{
					SetResults: []*groups.SetResult{
						{HomeScore: 25, GuestScore: 21},
						{HomeScore: 26, GuestScore: 28},
						{HomeScore: 23, GuestScore: 25},
						{HomeScore: 25, GuestScore: 0},
						{HomeScore: 7, GuestScore: 0},
					},
				},
			},
			expectedWins:               0,
			expectedMaxPossibleWins:    0,
			expectedPoints:             1,
			expectedMaxPossiblePoints:  1,
			expectedThreePointsMatches: 0,
			expectedTwoPointsMatches:   0,
			expectedOnePointMatches:    1,
			expectedZeroPointsMatches:  0,
			expectedWinSets:            2,
			expectedLoseSets:           3,
			expectedWinPoints:          74,
			expectedLosePoints:         106,
		},
		{
			name: "guest win 3:0",
			team: groups.Team{ID: 1},
			mr: groups.Match{
				GuestTeamID: 1,
				Result: &groups.MatchResult{
					SetResults: []*groups.SetResult{
						{HomeScore: 12, GuestScore: 21},
						{HomeScore: 26, GuestScore: 28},
						{HomeScore: 23, GuestScore: 25},
					},
				},
			},
			expectedWins:               1,
			expectedMaxPossibleWins:    1,
			expectedPoints:             3,
			expectedMaxPossiblePoints:  3,
			expectedThreePointsMatches: 1,
			expectedTwoPointsMatches:   0,
			expectedOnePointMatches:    0,
			expectedZeroPointsMatches:  0,
			expectedWinSets:            3,
			expectedLoseSets:           0,
			expectedWinPoints:          74,
			expectedLosePoints:         61,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			winScoring := groups.WinsScoringResult{}
			appendResults(&winScoring, tt.team, tt.mr)

			assert.Equal(t, tt.expectedWins, winScoring.Wins)
			assert.Equal(t, tt.expectedMaxPossibleWins, winScoring.MaxPossibleWins)
			assert.Equal(t, tt.expectedPoints, winScoring.Points)
			assert.Equal(t, tt.expectedMaxPossiblePoints, winScoring.MaxPossiblePoints)
			assert.Equal(t, tt.expectedThreePointsMatches, winScoring.ThreePointsMatches)
			assert.Equal(t, tt.expectedTwoPointsMatches, winScoring.TwoPointsMatches)
			assert.Equal(t, tt.expectedOnePointMatches, winScoring.OnePointMatches)
			assert.Equal(t, tt.expectedZeroPointsMatches, winScoring.ZeroPointsMatches)
			assert.Equal(t, tt.expectedWinSets, winScoring.WinSets)
			assert.Equal(t, tt.expectedLoseSets, winScoring.LoseSets)
			assert.Equal(t, tt.expectedWinPoints, winScoring.WinPoints)
			assert.Equal(t, tt.expectedLosePoints, winScoring.LosePoints)
		})
	}
}

func TestRatio(t *testing.T) {
	tests := []struct {
		name          string
		wins          int
		looses        int
		expectedRatio float64
	}{
		{
			name:          "no looses",
			wins:          10,
			looses:        0,
			expectedRatio: 10000,
		},
		{
			name:          "win eq looses",
			wins:          10,
			looses:        10,
			expectedRatio: 1,
		},
		{
			name:          "no wins",
			wins:          0,
			looses:        10,
			expectedRatio: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := getRatio(tt.wins, tt.looses)

			assert.Equal(t, tt.expectedRatio, r)
		})
	}
}

func TestGetWinScoring(t *testing.T) {
	handicapWins := 1
	handicapPoints := 2
	tests := []struct {
		name                       string
		team                       groups.Team
		mrs                        []groups.Match
		expectedWins               int
		expectedMaxPossibleWins    int
		expectedHandicapWins       int
		expectedPoints             int
		expectedMaxPossiblePoints  int
		expectedHandicapPoints     int
		expectedThreePointsMatches int
		expectedTwoPointsMatches   int
		expectedOnePointMatches    int
		expectedZeroPointsMatches  int
		expectedWinSets            int
		expectedLoseSets           int
		expectedWinPoints          int
		expectedLosePoints         int
	}{
		{
			name: "scoring",
			team: groups.Team{
				ID:             1,
				HandicapWins:   &handicapWins,
				HandicapPoints: &handicapPoints,
			},
			mrs: []groups.Match{
				{
					Result: &groups.MatchResult{
						SetResults: []*groups.SetResult{
							{HomeScore: 25, GuestScore: 21},
							{HomeScore: 26, GuestScore: 28},
							{HomeScore: 15, GuestScore: 25},
							{HomeScore: 25, GuestScore: 0},
							{HomeScore: 15, GuestScore: 10},
						},
					},
				},
			},
			expectedWins:               1,
			expectedMaxPossibleWins:    1,
			expectedHandicapWins:       1,
			expectedPoints:             2,
			expectedMaxPossiblePoints:  2,
			expectedHandicapPoints:     2,
			expectedThreePointsMatches: 0,
			expectedTwoPointsMatches:   0,
			expectedOnePointMatches:    0,
			expectedZeroPointsMatches:  0,
			expectedWinSets:            0,
			expectedLoseSets:           0,
			expectedWinPoints:          0,
			expectedLosePoints:         0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			winScoring := getWinScoring(tt.team, tt.mrs)

			assert.Equal(t, tt.expectedWins, winScoring.Wins)
			assert.Equal(t, tt.expectedMaxPossibleWins, winScoring.MaxPossibleWins)
			assert.Equal(t, tt.expectedHandicapWins, winScoring.HandicapWins)
			assert.Equal(t, tt.expectedPoints, winScoring.Points)
			assert.Equal(t, tt.expectedMaxPossiblePoints, winScoring.MaxPossiblePoints)
			assert.Equal(t, tt.expectedHandicapPoints, winScoring.HandicapPoints)
			assert.Equal(t, tt.expectedThreePointsMatches, winScoring.ThreePointsMatches)
			assert.Equal(t, tt.expectedTwoPointsMatches, winScoring.TwoPointsMatches)
			assert.Equal(t, tt.expectedOnePointMatches, winScoring.OnePointMatches)
			assert.Equal(t, tt.expectedZeroPointsMatches, winScoring.ZeroPointsMatches)
			assert.Equal(t, tt.expectedWinSets, winScoring.WinSets)
			assert.Equal(t, tt.expectedLoseSets, winScoring.LoseSets)
			assert.Equal(t, tt.expectedWinPoints, winScoring.WinPoints)
			assert.Equal(t, tt.expectedLosePoints, winScoring.LosePoints)
		})
	}
}

func TestSortAndSetPlaces(t *testing.T) {
	tests := []struct {
		name          string
		res           []groups.TeamResult
		expectedOrder map[int]int
	}{
		{
			name:          "empty results",
			res:           []groups.TeamResult{},
			expectedOrder: map[int]int{},
		},
		{
			name: "single team",
			res: []groups.TeamResult{
				groups.TeamResult{
					Team:                groups.Team{ID: 7},
					Place:               0,
					WinsScoringResult:   nil,
					PointsScoringResult: nil,
				},
			},
			expectedOrder: map[int]int{1: 7},
		},
		{
			name: "sort point scoring",
			res: []groups.TeamResult{
				groups.TeamResult{
					Team:  groups.Team{ID: 7},
					Place: 0,
					PointsScoringResult: &groups.PointsScoringResult{
						Points:     1,
						WinSets:    1,
						LoseSets:   1,
						WinPoints:  1,
						LosePoints: 1,
					},
				},
				groups.TeamResult{
					Team:  groups.Team{ID: 6},
					Place: 0,
					PointsScoringResult: &groups.PointsScoringResult{
						Points:     2,
						WinSets:    1,
						LoseSets:   1,
						WinPoints:  1,
						LosePoints: 1,
					},
				},
				groups.TeamResult{
					Team:  groups.Team{ID: 5},
					Place: 0,
					PointsScoringResult: &groups.PointsScoringResult{
						Points:     2,
						WinSets:    2,
						LoseSets:   1,
						WinPoints:  1,
						LosePoints: 1,
					},
				},
				groups.TeamResult{
					Team:  groups.Team{ID: 4, Name: "2"},
					Place: 0,
					PointsScoringResult: &groups.PointsScoringResult{
						Points:     2,
						WinSets:    4,
						LoseSets:   2,
						WinPoints:  2,
						LosePoints: 1,
					},
				},
				groups.TeamResult{
					Team:  groups.Team{ID: 3, Name: "1"},
					Place: 0,
					PointsScoringResult: &groups.PointsScoringResult{
						Points:     2,
						WinSets:    2,
						LoseSets:   1,
						WinPoints:  4,
						LosePoints: 2,
					},
				},
			},
			expectedOrder: map[int]int{1: 3, 2: 4, 3: 5, 4: 6, 5: 7},
		},
		{
			name: "sort win scoring",
			res: []groups.TeamResult{
				groups.TeamResult{
					Team:  groups.Team{ID: 7},
					Place: 0,
					WinsScoringResult: &groups.WinsScoringResult{
						Wins: 1,
					},
				},
				groups.TeamResult{
					Team:  groups.Team{ID: 6},
					Place: 0,
					WinsScoringResult: &groups.WinsScoringResult{
						Wins: 2,
						PointsScoringResult: groups.PointsScoringResult{
							Points: 1,
						},
					},
				},
				groups.TeamResult{
					Team:  groups.Team{ID: 5},
					Place: 0,
					WinsScoringResult: &groups.WinsScoringResult{
						Wins: 2,
						PointsScoringResult: groups.PointsScoringResult{
							Points: 2,
						},
					},
				},
			},
			expectedOrder: map[int]int{1: 5, 2: 6, 3: 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := sortAndSetPlaces(tt.res)

			for i, r := range res {
				assert.Equal(t, i+1, r.Place)
				assert.Equal(t, tt.expectedOrder[i+1], r.Team.ID)
			}
		})
	}
}

func TestCalculateResults(t *testing.T) {
	handicapWins := 1
	handicapPoints := 2
	te := groups.Team{
		ID:             1,
		Name:           "name",
		HandicapWins:   &handicapWins,
		HandicapPoints: &handicapPoints,
	}
	res := groups.TeamResult{
		Team:  te,
		Place: 1,
		PointsScoringResult: &groups.PointsScoringResult{
			Points:             1,
			MaxPossiblePoints:  2,
			HandicapPoints:     3,
			ThreePointsMatches: 4,
			TwoPointsMatches:   5,
			OnePointMatches:    6,
			ZeroPointsMatches:  7,
			WinSets:            8,
			LoseSets:           9,
			WinPoints:          0,
			LosePoints:         10,
		},
	}
	tests := []struct {
		name        string
		ts          []groups.Team
		ms          []groups.Match
		f           func(groups.Team, []groups.Match) groups.TeamResult
		expectedTrs []groups.TeamResult
	}{
		{
			name: "calc result",
			ts:   []groups.Team{te},
			ms:   []groups.Match{{HomeTeamID: 1}},
			f: func(groups.Team, []groups.Match) groups.TeamResult {
				return res
			},
			expectedTrs: []groups.TeamResult{res},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := CalculateResults(tt.ts, tt.ms, tt.f)

			assert.Equal(t, tt.expectedTrs, r)
		})
	}
}

func TestGetPointScoringResult(t *testing.T) {
	handicapWins := 1
	handicapPoints := 2
	te := groups.Team{
		ID:             1,
		Name:           "name",
		HandicapWins:   &handicapWins,
		HandicapPoints: &handicapPoints,
	}
	res := groups.TeamResult{
		Team: te,
		PointsScoringResult: &groups.PointsScoringResult{
			Points:            2,
			MaxPossiblePoints: 2,
			HandicapPoints:    2,
		},
	}
	tests := []struct {
		name        string
		t           groups.Team
		ms          []groups.Match
		expectedTrs groups.TeamResult
	}{
		{
			name:        "empty results",
			t:           te,
			ms:          []groups.Match{},
			expectedTrs: res,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := GetPointScoringResult(tt.t, tt.ms)

			assert.Equal(t, tt.expectedTrs, r)
		})
	}
}

func TestGetWinScoringResult(t *testing.T) {
	handicapWins := 1
	handicapPoints := 2
	te := groups.Team{
		ID:             1,
		Name:           "name",
		HandicapWins:   &handicapWins,
		HandicapPoints: &handicapPoints,
	}
	res := groups.TeamResult{
		Team: te,
		WinsScoringResult: &groups.WinsScoringResult{
			Wins:            1,
			MaxPossibleWins: 1,
			HandicapWins:    1,
			PointsScoringResult: groups.PointsScoringResult{
				Points:            2,
				MaxPossiblePoints: 2,
				HandicapPoints:    2,
			},
		},
	}
	tests := []struct {
		name        string
		t           groups.Team
		ms          []groups.Match
		expectedTrs groups.TeamResult
	}{
		{
			name:        "empty results",
			t:           te,
			ms:          []groups.Match{},
			expectedTrs: res,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := GetWinScoringResult(tt.t, tt.ms)

			assert.Equal(t, tt.expectedTrs, r)
		})
	}
}
