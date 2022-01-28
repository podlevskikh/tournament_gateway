package response_success

import "vollyemsk_tournament_gateway/models/teams"

type TeamsResponse struct {
	Teams []TeamResponse `json:"teams"`
}

type TeamResponse struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	FoundationDate string `json:"foundationDate"`
	HandicapWins   int    `json:"handicapWins,omitempty"`
	HandicapPoints int    `json:"handicapPoints,omitempty"`
}

func FromTeamsResponse(ts []teams.Team) TeamsResponse {
	trs := make([]TeamResponse, 0, len(ts))
	for _, t := range ts {
		tr := TeamResponse{
			ID:             t.ID,
			Name:           t.Name,
			Description:    t.Description,
			FoundationDate: t.FoundationDate.Format("2006-01-02"),
		}
		if t.HandicapWins != nil {
			tr.HandicapWins = *t.HandicapWins
		}
		if t.HandicapPoints != nil {
			tr.HandicapPoints = *t.HandicapPoints
		}
		trs = append(trs, tr)
	}
	return TeamsResponse{Teams: trs}
}
