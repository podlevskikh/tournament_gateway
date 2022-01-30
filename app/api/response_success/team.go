package response_success

import "vollyemsk_tournament_gateway/models/groups"

type TeamsResponse struct {
	Teams []TeamResponse `json:"teams"`
}

type TeamResponse struct {
	ID             int               `json:"id"`
	Name           string            `json:"name"`
	Description    string            `json:"description"`
	Foundation     string            `json:"foundation"`
	HandicapWins   int               `json:"handicapWins,omitempty"`
	HandicapPoints int               `json:"handicapPoints,omitempty"`
	ContactUsers   []UserResponse    `json:"contactUsers,omitempty"`
	HomeGyms       []TeamGymResponse `json:"homeGyms,omitempty"`
}

type TeamGymResponse struct {
	Place      GymResponse `json:"place"`
	WeekDay    int         `json:"weekDay"`
	TimeFrom   string      `json:"timeFrom"`
	TimeTo     string      `json:"timeTo"`
	TimeWarmUp string      `json:"timeWarmUp"`
}

func FromTeamsResponse(ts []groups.Team) TeamsResponse {
	trs := make([]TeamResponse, 0, len(ts))
	for _, t := range ts {
		trs = append(trs, FromTeamResponse(t))
	}
	return TeamsResponse{Teams: trs}
}

func FromTeamResponse(t groups.Team) TeamResponse {
	tr := TeamResponse{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
		Foundation:  t.Foundation,
	}
	tr.HomeGyms = make([]TeamGymResponse, 0, len(t.HomeGyms))
	for _, hg := range t.HomeGyms {
		tr.HomeGyms = append(tr.HomeGyms, TeamGymResponse{
			Place:      FromGymResponse(hg.Gym),
			WeekDay:    int(hg.WeekDay),
			TimeFrom:   hg.TimeFrom,
			TimeTo:     hg.TimeTo,
			TimeWarmUp: hg.TimeWarmUp,
		})
	}
	usr := FromUsersResponse(t.ContactUsers)
	tr.ContactUsers = usr.Users
	if t.HandicapWins != nil {
		tr.HandicapWins = *t.HandicapWins
	}
	if t.HandicapPoints != nil {
		tr.HandicapPoints = *t.HandicapPoints
	}
	return tr
}
