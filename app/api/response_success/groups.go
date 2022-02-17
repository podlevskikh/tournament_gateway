package response_success

import (
	"tournament_gateway/models/groups"
)

type GroupsResponse struct {
	Groups []GroupResponse `json:"groups"`
}

type GroupResponse struct {
	Alias           string                `json:"alias"`
	ShortName       string                `json:"shortName"`
	Name            string                `json:"name"`
	Description     string                `json:"description"`
	TournamentAlias string                `json:"tournamentAlias"`
	SeasonAlias     string                `json:"seasonAlias"`
	StageAlias      string                `json:"stageAlias"`
	LeagueAlias     string                `json:"leagueAlias"`
	Tournament      TournamentResponse    `json:"tournament,omitempty"`
	Season          SeasonResponse        `json:"season,omitempty"`
	Stage           StageResponse         `json:"stage,omitempty"`
	League          LeagueResponse        `json:"league,omitempty"`
	Results         []GroupResultResponse `json:"results,omitempty"`
}

func FromGroupsResponse(gs []*groups.Group) GroupsResponse {
	grs := make([]GroupResponse, 0, len(gs))
	for _, g := range gs {
		grs = append(grs, FromGroupResponse(g))
	}
	return GroupsResponse{Groups: grs}
}

func FromGroupResponse(g *groups.Group) GroupResponse {
	res := GroupResponse{
		Alias:           g.Alias,
		ShortName:       g.ShortName,
		Name:            g.Name,
		Description:     g.Description,
		TournamentAlias: g.TournamentAlias,
		SeasonAlias:     g.SeasonAlias,
		StageAlias:      g.StageAlias,
		LeagueAlias:     g.LeagueAlias,
		Results:         FromGroupResultsResponse(g.GroupResults).GroupResults,
	}
	if g.Tournament != nil {
		res.Tournament = FromTournamentResponse(g.Tournament)
	}
	if g.Season != nil {
		res.Season = FromSeasonResponse(g.Season)
	}
	if g.Stage != nil {
		res.Stage = FromStageResponse(g.Stage)
	}
	if g.League != nil {
		res.League = FromLeagueResponse(g.League)
	}
	return res
}
