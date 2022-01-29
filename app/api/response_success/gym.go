package response_success

import "vollyemsk_tournament_gateway/models/gyms"

type GymResponse struct {
	ID           int    `json:"id"`
	MetroStation string `json:"metroStation"`
	Address      string `json:"address"`
}

func FromGymResponse(g *gyms.Gym) GymResponse {
	return GymResponse{
		ID:           g.ID,
		MetroStation: g.MetroStation,
		Address:      g.Address,
	}
}
