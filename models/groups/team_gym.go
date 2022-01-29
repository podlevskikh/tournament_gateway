package groups

import "vollyemsk_tournament_gateway/models/gyms"

type WeekDay int

const (
	WeekDayMonday    WeekDay = 1
	WeekDayTuesday   WeekDay = 2
	WeekDayWednesday WeekDay = 3
	WeekDayThursday  WeekDay = 4
	WeekDayFriday    WeekDay = 5
	WeekDaySaturday  WeekDay = 6
	WeekDaySunday    WeekDay = 7
)

type TeamGym struct {
	TeamID     int
	GymID      int
	Gym        *gyms.Gym `gorm:"ForeignKey:GymID;AssociationForeignKey:ID"`
	WeekDay    WeekDay
	TimeFrom   string
	TimeTo     string
	TimeWarmUp string
}

func (tg *TeamGym) TableName() string {
	return "team2gym"
}
