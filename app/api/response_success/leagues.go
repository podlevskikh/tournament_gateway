package response_success

import "vollyemsk_tournament_gateway/models/leagues"

type LeaguesResponse struct {
	Leagues []LeagueResponse `json:"leagues"`
}

type LeagueResponse struct {
	Alias           string `json:"alias"`
	Name            string `json:"name"`
	StrengthWeight  int    `json:"strengthWeight"`
}

func FromLeaguesResponse(leagues []*leagues.League) LeaguesResponse {
	ls := make([]LeagueResponse, 0, len(leagues))
	for _, l := range leagues {
		ls = append(ls, LeagueResponse{
			Alias:           l.Alias,
			Name:            l.Name,
			StrengthWeight:  l.StrengthWeight,
		})
	}
	return LeaguesResponse{Leagues: ls}
}
