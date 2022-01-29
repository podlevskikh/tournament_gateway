package response_success

import (
	"vollyemsk_tournament_gateway/models/groups"
)

type PlayersResponse struct {
	Players []PlayerResponse `json:"players"`
}

type PlayerResponse struct {
	Id         int    `json:"id"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	MiddleName string `json:"middleName"`
	BirthDate  string `json:"birthDate"`
	Gender     string `json:"gender"`
	Skill      string `json:"skill"`
	AvatarUrl  string `json:"avatarUrl"`
}

func FromPlayersResponse(ps []*groups.Player) PlayersResponse {
	psr := make([]PlayerResponse, 0, len(ps))
	for _, p := range ps {
		psr = append(psr, PlayerResponse{
			Id:         p.ID,
			FirstName:  p.FirstName,
			LastName:   p.LastName,
			MiddleName: p.MiddleName,
			BirthDate:  p.BirthDate.Format("2006-01-02"),
			Gender:     string(p.Gender),
			Skill:      p.Skill,
			AvatarUrl:  p.AvatarUrl,
		})
	}
	return PlayersResponse{Players: psr}
}
