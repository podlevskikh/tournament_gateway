package response_success

import (
	"sort"
	"tournament_gateway/models/leagues"
)

type LeaguesResponse struct {
	Leagues []LeagueResponse `json:"leagues"`
}

type LeagueResponse struct {
	Alias          string `json:"alias"`
	ShortName      string `json:"shortName"`
	Name           string `json:"name"`
	StrengthWeight int    `json:"strengthWeight"`
}

func FromLeaguesResponse(leagues []*leagues.League) LeaguesResponse {
	ls := make([]LeagueResponse, 0, len(leagues))
	for _, l := range leagues {
		ls = append(ls, FromLeagueResponse(l))
	}

	sort.Slice(
		ls,
		func(i, j int) bool {
			return ls[i].StrengthWeight < ls[j].StrengthWeight
		})
	return LeaguesResponse{Leagues: ls}
}

func FromLeagueResponse(l *leagues.League) LeagueResponse {
	return LeagueResponse{
		Alias:          l.Alias,
		ShortName:      l.ShortName,
		Name:           l.Name,
		StrengthWeight: l.StrengthWeight,
	}
}
