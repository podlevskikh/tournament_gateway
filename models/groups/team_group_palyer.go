package groups

type TeamGroupPlayer struct {
	TeamID int
	GroupAlias string
	PlayerID int
	Player Player `gorm:"ForeignKey:PlayerID;AssociationForeignKey:ID"`
}

func (tgp *TeamGroupPlayer) TableName() string {
	return "team_group_player"
}
