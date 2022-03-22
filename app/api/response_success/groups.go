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
	Tournament      *TournamentResponse    `json:"tournament,omitempty"`
	Season          *SeasonResponse        `json:"season,omitempty"`
	Stage           *StageResponse         `json:"stage,omitempty"`
	League          *LeagueResponse        `json:"league,omitempty"`
	Results         []GroupResultResponse `json:"results,omitempty"`
}

func FromGroupsResponse(gs []*groups.Group, withInnerObjects bool) GroupsResponse {
	grs := make([]GroupResponse, 0, len(gs))
	for _, g := range gs {
		grs = append(grs, FromGroupResponse(g, withInnerObjects))
	}
	return GroupsResponse{Groups: grs}
}

func FromGroupResponse(g *groups.Group, withInnerObjects bool) GroupResponse {
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
	if withInnerObjects {
		if g.Tournament != nil {
			t := FromTournamentResponse(g.Tournament)
			res.Tournament = &t
		}
		if g.Season != nil {
			se := FromSeasonResponse(g.Season)
			res.Season = &se
		}
		if g.Stage != nil {
			st := FromStageResponse(g.Stage)
			res.Stage = &st
		}
		if g.League != nil {
			l := FromLeagueResponse(g.League)
			res.League = &l
		}
	}
	return res
}
