package response_success

import "tournament_gateway/models/referees"

type RefereeResponse struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	MiddleName  string `json:"middleName"`
	Phone       string `json:"phone"`
	Skill       string `json:"skill"`
	Description string `json:"description"`
	Price       string `json:"price"`
}

func FromRefereeResponse(r referees.Referee) RefereeResponse {
	return RefereeResponse{
		ID:          r.ID,
		FirstName:   r.FirstName,
		LastName:    r.LastName,
		MiddleName:  r.MiddleName,
		Phone:       r.Phone,
		Skill:       r.Skill,
		Description: r.Description,
		Price:       r.Price,
	}
}
