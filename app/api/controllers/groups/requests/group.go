package requests

type GroupRequest struct {
	Alias           string `json:"alias"`
	ShortName       string `json:"shortName"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	TournamentAlias string `json:"tournamentAlias"`
	SeasonAlias     string `json:"seasonAlias"`
	StageAlias      string `json:"stageAlias"`
	LeagueAlias     string `json:"leagueAlias"`
}
