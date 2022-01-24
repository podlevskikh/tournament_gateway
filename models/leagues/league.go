package leagues

import (
	"vollyemsk_tournament_gateway/models/seasons"
	"vollyemsk_tournament_gateway/models/tournaments"
)

type League struct {
	Alias          string
	Name           string
	StrengthWeight int
	Tournament     *tournaments.Tournament
	Season         *seasons.Season
	Stage          *seasons.Stage
}
