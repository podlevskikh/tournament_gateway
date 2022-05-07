package response_success

import (
	"tournament_gateway/models/groups"
)

type MatchesResponse struct {
	Matches []MatchResponse `json:"matches"`
}

type MatchResponse struct {
	ID        int                  `json:"id"`
	DateTime  string               `json:"datetime"`
	HomeTeam  TeamResponse         `json:"homeTeam"`
	GuestTeam TeamResponse         `json:"guestTeam"`
	Result    *MatchResultResponse `json:"result,omitempty"`
	Gym       *GymResponse         `json:"gym,omitempty"`
}

type MatchResultResponse struct {
	MatchDatetime          string              `json:"matchDatetime"`
	Winner                 string              `json:"winner"`
	HomePoints             int                 `json:"homePoints"`
	GuestPoints            int                 `json:"guestPoints"`
	HomeSetsWin            int                 `json:"homeSetsWin"`
	GuestSetsWin           int                 `json:"guestSetsWin"`
	SetResults             []SetResultResponse `json:"setResults"`
	HomeTeamPlayers        []PlayerResponse    `json:"homeTeamPlayers,omitempty"`
	GuestTeamPlayers       []PlayerResponse    `json:"guestTeamPlayers,omitempty"`
	HomeBestPlayerId       int                 `json:"homeBestPlayerId"`
	GuestBestPlayerId      int                 `json:"guestBestPlayerId"`
	Referees               []*RefereeResponse  `json:"referees"`
	HomeRefereeEvaluation  int                 `json:"homeRefereeEvaluation"`
	GuestRefereeEvaluation int                 `json:"guestRefereeEvaluation"`
}

type SetResultResponse struct {
	SetNumber  int `json:"setNumber"`
	HomeScore  int `json:"homeScore"`
	GuestScore int `json:"guestScore"`
}

func FromMatchesResponse(ms []groups.Match) MatchesResponse {
	mrs := make([]MatchResponse, 0, len(ms))
	for _, m := range ms {
		mrs = append(mrs, FromMatchResponse(m))
	}
	return MatchesResponse{Matches: mrs}
}

func FromMatchResponse(m groups.Match) MatchResponse {
	mr := MatchResponse{
		ID:        m.ID,
		DateTime:  m.Date.Format("2006-01-02 15:04:05"),
		HomeTeam:  FromTeamResponse(m.HomeTeam),
		GuestTeam: FromTeamResponse(m.GuestTeam),
	}
	if m.Gym != nil {
		g := FromGymResponse(m.Gym)
		mr.Gym = &g
	}
	if m.Result != nil {
		homeSetsWin := 0
		guestSetsWin := 0
		setResults := make([]SetResultResponse, 0, len(m.Result.SetResults))
		for _, sr := range m.Result.SetResults {
			setResults = append(setResults, SetResultResponse{
				SetNumber:  sr.SetNumber,
				HomeScore:  sr.HomeScore,
				GuestScore: sr.GuestScore,
			})
			if sr.HomeScore > sr.GuestScore {
				homeSetsWin++
			} else {
				guestSetsWin++
			}
		}
		rs := make([]*RefereeResponse, 0, len(m.Result.Referees))
		for _, r := range m.Result.Referees {
			rr := FromRefereeResponse(*r)
			rs = append(rs, &rr)
		}
		res := MatchResultResponse{
			MatchDatetime:          m.Result.MatchDatetime.Format("2006-01-02 15:04:05"),
			Winner:                 string(m.Result.Winner),
			HomePoints:             m.Result.HomePoints,
			GuestPoints:            m.Result.GuestPoints,
			HomeSetsWin:            homeSetsWin,
			GuestSetsWin:           guestSetsWin,
			SetResults:             setResults,
			HomeTeamPlayers:        FromPlayersResponse(m.Result.HomeTeamPlayers).Players,
			GuestTeamPlayers:       FromPlayersResponse(m.Result.GuestTeamPlayers).Players,
			HomeBestPlayerId:       m.Result.HomeBestPlayerID,
			GuestBestPlayerId:      m.Result.GuestBestPlayerID,
			HomeRefereeEvaluation:  m.Result.HomeRefereeEvaluation,
			GuestRefereeEvaluation: m.Result.GuestRefereeEvaluation,
			Referees:               rs,
		}
		mr.Result = &res
	}
	return mr
}
