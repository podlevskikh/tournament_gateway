package response_success

import "tournament_gateway/models/gyms"

type GymResponse struct {
	ID          int             `json:"id"`
	Address     string          `json:"address"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Metros      []MetroResponse `json:"metros"`
}

type MetroResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FromGymResponse(g *gyms.Gym) GymResponse {
	ms := make([]MetroResponse, 0, len(g.Metros))
	for _, m := range g.Metros {
		ms = append(ms, MetroResponse{
			ID:   m.ID,
			Name: m.Name,
		})
	}
	return GymResponse{
		ID:          g.ID,
		Address:     "", //todo check privilegues g.Address,
		Name:        "", //todo check privilegues g.Name,
		Description: g.Description,
		Metros:      ms,
	}
}
