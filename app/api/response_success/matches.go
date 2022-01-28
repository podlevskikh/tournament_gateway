package response_success

import (
	"vollyemsk_tournament_gateway/models/matches"
)

type MatchesResponse struct {
	Matches []MatchResponse `json:"matches"`
}

type MatchResponse struct {
	ID          int                  `json:"id"`
	Date        string               `json:"date"`
	HomeTeamID  int                  `json:"homeTeamID"`
	GuestTeamID int                  `json:"guestTeamID"`
	Result      *MatchResultResponse `json:"result,omitempty"`
}

type MatchResultResponse struct {
	MatchDatetime string              `json:"matchDatetime"`
	Winner        string              `json:"winner"`
	HomePoints    int                 `json:"homePoints"`
	GuestPoints   int                 `json:"guestPoints"`
	SetResults    []SetResultResponse `json:"setResults"`
	//HomeTeamPlayers        int    `json:"homeTeamPlayers"`
	//GuestTeamPlayers       int    `json:"guestTeamPlayers"`
	HomeBestPlayerId  int `json:"homeBestPlayerId"`
	GuestBestPlayerId int `json:"guestBestPlayerId"`
	//Referee                int    `json:"referee"`
	HomeRefereeEvaluation  int `json:"homeRefereeEvaluation"`
	GuestRefereeEvaluation int `json:"guestRefereeEvaluation"`
}

type SetResultResponse struct {
	SetNumber  int `json:"setNumber"`
	HomeScore  int `json:"homeScore"`
	GuestScore int `json:"guestScore"`
}

func FromMatchesResponse(ms []matches.Match) MatchesResponse {
	mrs := make([]MatchResponse, 0, len(ms))
	for _, m := range ms {
		mr := MatchResponse{
			ID:          m.ID,
			Date:        m.Date.Format("2006-01-02"),
			HomeTeamID:  m.HomeTeamID,
			GuestTeamID: m.GuestTeamID,
		}
		if m.Result != nil {
			setResults := make([]SetResultResponse, 0, len(m.Result.SetResults))
			for _, sr := range m.Result.SetResults {
				setResults = append(setResults, SetResultResponse{
					SetNumber:  sr.SetNumber,
					HomeScore:  sr.HomeScore,
					GuestScore: sr.GuestScore,
				})
			}
			res := MatchResultResponse{
				MatchDatetime:          m.Result.MatchDatetime.Format("2006-01-02 15:04:05"),
				Winner:                 string(m.Result.Winner),
				HomePoints:             m.Result.HomePoints,
				GuestPoints:            m.Result.GuestPoints,
				SetResults:             setResults,
				HomeBestPlayerId:       m.Result.HomeBestPlayerID,
				GuestBestPlayerId:      m.Result.GuestBestPlayerID,
				HomeRefereeEvaluation:  m.Result.HomeRefereeEvaluation,
				GuestRefereeEvaluation: m.Result.GuestRefereeEvaluation,
			}
			mr.Result = &res
		}
		mrs = append(mrs, mr)
	}
	return MatchesResponse{Matches: mrs}
}
