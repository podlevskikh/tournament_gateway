package requests

type TournamentRequest struct {
	Alias       string `json:"alias"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Gender      string `json:"gender"`
}