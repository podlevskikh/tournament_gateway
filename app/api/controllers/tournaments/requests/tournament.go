package requests

type TournamentRequest struct {
	Alias       string `json:"alias"`
	ShortName   string `json:"shortName"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Gender      string `json:"gender"`
}
