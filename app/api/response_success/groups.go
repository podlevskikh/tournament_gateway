package response_success

import (
	"tournament_gateway/models/groups"
)

type GroupsResponse struct {
	Groups []GroupResponse `json:"groups"`
}

type GroupResponse struct {
	Alias           string `json:"alias"`
	ShortName       string `json:"shortName"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	TournamentAlias string `json:"tournamentAlias"`
	SeasonAlias     string `json:"seasonAlias"`
	StageAlias      string `json:"stageAlias"`
	LeagueAlias     string `json:"leagueAlias"`
}

func FromGroupsResponse(grs []*groups.Group) GroupsResponse {
	gs := make([]GroupResponse, 0, len(grs))
	for _, g := range grs {
		gs = append(gs, GroupResponse{
			Alias:           g.Alias,
			ShortName:       g.ShortName,
			Name:            g.Name,
			Description:     g.Description,
			TournamentAlias: g.Tournament.Alias,
			SeasonAlias:     g.Season.Alias,
			StageAlias:      g.Stage.Alias,
			LeagueAlias:     g.League.Alias,
		})
	}
	return GroupsResponse{Groups: gs}
}
