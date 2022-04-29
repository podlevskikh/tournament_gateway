package response_success

import (
	"sort"
	"tournament_gateway/models/groups"
	"tournament_gateway/models/leagues"
	"tournament_gateway/models/seasons"
	"tournament_gateway/models/tournaments"
)

type CurrentResponse struct {
	CurrentStage       StageResponse       `json:"currentStage"`
	CurrentSeason      SeasonResponse      `json:"currentSeason"`
	CurrentTournaments []CurrentTournament `json:"currentTournaments"`
}

type CurrentTournament struct {
	Tournament     TournamentResponse `json:"tournament"`
	CurrentLeagues []CurrentLeague    `json:"currentLeagues"`
}

type CurrentLeague struct {
	League        LeagueResponse  `json:"league"`
	CurrentGroups []GroupResponse `json:"groups"`
}

func FromCurrentTournamentResponse(t *tournaments.Tournament, gs []*groups.Group) CurrentTournament {
	ls := getLeagues(gs, t)

	cls := make([]CurrentLeague, 0, len(ls))
	for _, l := range ls {
		cls = append(cls, FromLeagueGroupsResponse(t, l, gs))
	}

	sort.Slice(cls, func(i,j int) bool {
		return cls[i].League.StrengthWeight < cls[j].League.StrengthWeight
	})

	return CurrentTournament{
		Tournament:     FromTournamentResponse(t),
		CurrentLeagues: cls,
	}
}

func FromLeagueGroupsResponse(t *tournaments.Tournament, l *leagues.League, gs []*groups.Group) CurrentLeague {
	cgs := make([]*groups.Group, 0, len(gs))
	for _, g := range gs {
		if l.Alias == g.League.Alias && t.Alias == g.Tournament.Alias {
			cgs = append(cgs, g)
		}
	}
	return CurrentLeague{
		League:        FromLeagueResponse(l),
		CurrentGroups: FromGroupsResponse(cgs, false).Groups,
	}
}

func FromCurrentResponse(gs []*groups.Group) CurrentResponse {
	ts, se, st, cgs := filterCurrentObjects(gs)

	cts := make([]CurrentTournament, 0, len(ts))
	for _, t := range ts {
		cts = append(cts, FromCurrentTournamentResponse(t, cgs))
	}
	sort.Slice(cts, func(i,j int) bool {
		return cts[i].Tournament.Gender == "male" //todo пока сортируем по полу
	})
	return CurrentResponse{
		CurrentStage:       FromStageResponse(st),
		CurrentSeason:      FromSeasonResponse(se),
		CurrentTournaments: cts,
	}
}

func filterCurrentObjects(gs []*groups.Group) ([]*tournaments.Tournament, *seasons.Season, *seasons.Stage, []*groups.Group) {
	cgs := make([]*groups.Group, 0, len(gs))
	cts := make([]*tournaments.Tournament, 0, len(gs))
	for _, g := range gs {
		if g.Stage.IsCurrent {
			cgs = append(cgs, g)
			if !tournamentExists(g.Tournament, cts) {
				cts = append(cts, g.Tournament)
			}
		}
	}
	return cts, cgs[0].Season, cgs[0].Stage, cgs
}

func tournamentExists(t *tournaments.Tournament, ts []*tournaments.Tournament) bool {
	for _, tr := range ts {
		if t.Alias == tr.Alias {
			return true
		}
	}
	return false
}

func getLeagues(gs []*groups.Group, t *tournaments.Tournament) []*leagues.League {
	ls := make([]*leagues.League, 0, len(gs))
	for _, g := range gs {
		if g.Tournament.Alias == t.Alias && !leagueExists(g.League, ls) {
			ls = append(ls, g.League)
		}
	}
	return ls
}

func leagueExists(l *leagues.League, ls []*leagues.League) bool {
	for _, lr := range ls {
		if l.Alias == lr.Alias {
			return true
		}
	}
	return false
}