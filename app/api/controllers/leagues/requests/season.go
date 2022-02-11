package requests

type LeagueRequest struct {
	Alias          string `json:"alias"`
	Name           string `json:"name"`
	StrengthWeight int    `json:"strengthWeight"`
}
