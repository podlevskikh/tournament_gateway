package groups

import (
	"vollyemsk_tournament_gateway/models/leagues"
	"vollyemsk_tournament_gateway/models/seasons"
	"vollyemsk_tournament_gateway/models/tournaments"
)

type Group struct {
	Alias           string
	ShortName       string
	Name            string
	Description     string
	TournamentAlias string
	SeasonAlias     string
	StageAlias      string
	LeagueAlias     string
	Tournament      *tournaments.Tournament `gorm:"ForeignKey:TournamentAlias;AssociationForeignKey:Alias"`
	Season          *seasons.Season         `gorm:"ForeignKey:SeasonAlias;AssociationForeignKey:Alias"`
	Stage           *seasons.Stage          `gorm:"ForeignKey:StageAlias;AssociationForeignKey:Alias"`
	League          *leagues.League         `gorm:"ForeignKey:LeagueAlias;AssociationForeignKey:Alias"`
}
