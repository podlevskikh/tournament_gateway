package groups

import (
	"vollyemsk_tournament_gateway/models/leagues"
	"vollyemsk_tournament_gateway/models/seasons"
	"vollyemsk_tournament_gateway/models/tournaments"
)

type Group struct {
	Alias       string
	ShortName   string
	Name        string
	Description string
	Tournament  *tournaments.Tournament
	Season      *seasons.Season
	Stage       *seasons.Stage
	League      *leagues.League
}
