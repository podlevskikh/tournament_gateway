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

func FromGroupsResponse(gs []*groups.Group) GroupsResponse {
	grs := make([]GroupResponse, 0, len(gs))
	for _, g := range gs {
		grs = append(grs, FromGroupResponse(g))
	}
	return GroupsResponse{Groups: grs}
}

func FromGroupResponse(g *groups.Group) GroupResponse {
	return GroupResponse{
		Alias:           g.Alias,
		ShortName:       g.ShortName,
		Name:            g.Name,
		Description:     g.Description,
		TournamentAlias: g.Tournament.Alias,
		SeasonAlias:     g.Season.Alias,
		StageAlias:      g.Stage.Alias,
		LeagueAlias:     g.League.Alias,
	}
}
