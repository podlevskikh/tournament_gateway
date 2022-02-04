package response_success

import "vollyemsk_tournament_gateway/models/groups"

type GroupResultsResponse struct {
	GroupResults []GroupResultResponse `json:"group_results"`
}

type GroupResultResponse struct {
	TeamResults []TeamResultResponse `json:"teamResults"`
	Name        string               `json:"name"`
	SoringType  string               `json:"scoringType"`
}

type TeamResultResponse struct {
	Team                TeamResponse                 `json:"team"`
	Place               int                          `json:"place"`
	WinsScoringResult   *WinsScoringResultResponse   `json:"winsScoringResult,omitempty"`
	PointsScoringResult *PointsScoringResultResponse `json:"pointsScoringResult,omitempty"`
}

type WinsScoringResultResponse struct {
	Wins               int `json:"wins"`
	MaxPossibleWins    int `json:"maxPossibleWins"`
	Points             int `json:"points"`
	MaxPossiblePoints  int `json:"maxPossiblePoints"`
	HandicapWins       int `json:"handicapWins"`
	HandicapPoints     int `json:"handicapPoints"`
	ThreePointsMatches int `json:"threePointsMatches"`
	TwoPointsMatches   int `json:"twoPointsMatches"`
	OnePointMatches    int `json:"onePointMatches"`
	ZeroPointsMatches  int `json:"zeroPointsMatches"`
	WinSets            int `json:"winSets"`
	LoseSets           int `json:"loseSets"`
	WinPoints          int `json:"winPoints"`
	LosePoints         int `json:"losePoints"`
}

type PointsScoringResultResponse struct {
	Points             int `json:"points"`
	MaxPossiblePoints  int `json:"maxPossiblePoints"`
	HandicapPoints     int `json:"handicapPoints"`
	ThreePointsMatches int `json:"threePointsMatches"`
	TwoPointsMatches   int `json:"twoPointsMatches"`
	OnePointMatches    int `json:"onePointMatches"`
	ZeroPointsMatches  int `json:"zeroPointsMatches"`
	WinSets            int `json:"winSets"`
	LoseSets           int `json:"loseSets"`
	WinPoints          int `json:"winPoints"`
	LosePoints         int `json:"losePoints"`
}

func FromGroupResultsResponse(grs []groups.GroupResult) GroupResultsResponse {
	grrs := make([]GroupResultResponse, 0, len(grs))
	for _, gr := range grs {
		grrs = append(grrs, GroupResultResponse{
			TeamResults: FromTeamResultsResponse(gr.TeamResults),
			Name:        gr.Name,
			SoringType:  gr.ScoringType,
		})
	}
	return GroupResultsResponse{GroupResults: grrs}
}

func FromTeamResultsResponse(trs []groups.TeamResult) []TeamResultResponse {
	trrs := make([]TeamResultResponse, 0, len(trs))
	for _, tr := range trs {
		trr := TeamResultResponse{
			Team:  FromTeamResponse(tr.Team),
			Place: tr.Place,
		}
		if tr.WinsScoringResult != nil {
			trr.WinsScoringResult = &WinsScoringResultResponse{
				Wins:               tr.WinsScoringResult.Wins,
				MaxPossibleWins:    tr.WinsScoringResult.MaxPossibleWins,
				Points:             tr.WinsScoringResult.Points,
				MaxPossiblePoints:  tr.WinsScoringResult.MaxPossiblePoints,
				HandicapWins:       tr.WinsScoringResult.HandicapWins,
				HandicapPoints:     tr.WinsScoringResult.HandicapPoints,
				ThreePointsMatches: tr.WinsScoringResult.ThreePointsMatches,
				TwoPointsMatches:   tr.WinsScoringResult.TwoPointsMatches,
				OnePointMatches:    tr.WinsScoringResult.OnePointMatches,
				ZeroPointsMatches:  tr.WinsScoringResult.ZeroPointsMatches,
				WinSets:            tr.WinsScoringResult.WinSets,
				LoseSets:           tr.WinsScoringResult.LoseSets,
				WinPoints:          tr.WinsScoringResult.WinPoints,
				LosePoints:         tr.WinsScoringResult.LosePoints,
			}
		}
		if tr.PointsScoringResult != nil {
			trr.PointsScoringResult = &PointsScoringResultResponse{
				Points:             tr.PointsScoringResult.Points,
				MaxPossiblePoints:  tr.PointsScoringResult.MaxPossiblePoints,
				HandicapPoints:     tr.PointsScoringResult.HandicapPoints,
				ThreePointsMatches: tr.PointsScoringResult.ThreePointsMatches,
				TwoPointsMatches:   tr.PointsScoringResult.TwoPointsMatches,
				OnePointMatches:    tr.PointsScoringResult.OnePointMatches,
				ZeroPointsMatches:  tr.PointsScoringResult.ZeroPointsMatches,
				WinSets:            tr.PointsScoringResult.WinSets,
				LoseSets:           tr.PointsScoringResult.LoseSets,
				WinPoints:          tr.PointsScoringResult.WinPoints,
				LosePoints:         tr.PointsScoringResult.LosePoints,
			}
		}
		trrs = append(trrs, trr)
	}
	return trrs
}
