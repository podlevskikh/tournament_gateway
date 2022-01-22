package teams

import "time"

type Gender string

const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
)

type Player struct {
	ID         int
	FirstName  string
	LastName   string
	MiddleName string
	BirthDate  time.Time
	Gender     Gender
	Skill      string
	AvatarUrl  string
}
