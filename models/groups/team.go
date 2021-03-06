package groups

import (
	"tournament_gateway/models/users"
)

type Team struct {
	ID             int
	Name           string
	Description    string
	Foundation     string
	Groups         []*Group      `gorm:"many2many:team2group;foreignKey:ID;joinTableForeignKey:TeamID;associationForeignKey:Alias;associationJoinTableForeignKey:GroupAlias"`
	ContactUsers   []*users.User `gorm:"many2many:team2user;foreignKey:ID;joinTableForeignKey:TeamID;associationForeignKey:ID;associationJoinTableForeignKey:UserID"`
	HomeGyms       []*TeamGym    `gorm:"ForeignKey:TeamID;AssociationForeignKey:ID"`
	HandicapWins   *int
	HandicapPoints *int
}
