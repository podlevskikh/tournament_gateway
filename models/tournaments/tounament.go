package tournaments

type TournamentGender string

const (
	TournamentGenderUndefined TournamentGender = "undefined"
	TournamentGenderMale      TournamentGender = "male"
	TournamentGenderFemale    TournamentGender = "female"
)

type Tournament struct {
	//ID          string `gorm:"column:alias"`
	Alias       string
	ShortName   string
	Name        string
	Description string
	Gender      TournamentGender
}
