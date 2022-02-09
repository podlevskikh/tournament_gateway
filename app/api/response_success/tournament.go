package response_success

import "tournament_gateway/models/tournaments"

type TournamentsResponse struct {
	Tournaments []TournamentResponse `json:"tournaments"`
}

type TournamentResponse struct {
	Alias       string `json:"alias"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Gender      string `json:"gender"`
}

func FromTournamentsResponse(ts []*tournaments.Tournament) TournamentsResponse {
	res := make([]TournamentResponse, 0, len(ts))
	for _, t := range ts {
		res = append(res, FromTournamentResponse(t))
	}
	return TournamentsResponse{Tournaments: res}
}

func FromTournamentResponse(t *tournaments.Tournament) TournamentResponse {
	return TournamentResponse{
		Alias:       t.Alias,
		Name:        t.Name,
		Description: t.Description,
		Gender:      string(t.Gender),
	}
}
