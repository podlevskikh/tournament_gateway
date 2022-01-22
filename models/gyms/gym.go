package gyms

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

type Gym struct {
	Place      Place
	WeekDay    WeekDay
	TimeFrom   string
	TimeTo     string
	TimeWarmUp string
}
