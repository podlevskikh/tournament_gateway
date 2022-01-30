package gyms

type Gym struct {
	ID          int
	Address     string
	Name        string
	Description string
	Metros      []*Metro `gorm:"many2many:gym2metro;foreignKey:ID;joinTableForeignKey:GymID;associationForeignKey:ID;associationJoinTableForeignKey:MetroID"`
}
