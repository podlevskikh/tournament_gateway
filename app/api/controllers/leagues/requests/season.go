package requests

type LeagueRequest struct {
	Alias          string `json:"alias"`
	ShortName      string `json:"shortName"`
	Name           string `json:"name"`
	StrengthWeight int    `json:"strengthWeight"`
}
