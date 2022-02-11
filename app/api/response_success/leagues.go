package response_success

import "tournament_gateway/models/leagues"

type LeaguesResponse struct {
	Leagues []LeagueResponse `json:"leagues"`
}

type LeagueResponse struct {
	Alias          string `json:"alias"`
	Name           string `json:"name"`
	StrengthWeight int    `json:"strengthWeight"`
}

func FromLeaguesResponse(leagues []*leagues.League) LeaguesResponse {
	ls := make([]LeagueResponse, 0, len(leagues))
	for _, l := range leagues {
		ls = append(ls, FromLeagueResponse(l))
	}
	return LeaguesResponse{Leagues: ls}
}

func FromLeagueResponse(l *leagues.League) LeagueResponse {
	return LeagueResponse{
		Alias:          l.Alias,
		Name:           l.Name,
		StrengthWeight: l.StrengthWeight,
	}
}
