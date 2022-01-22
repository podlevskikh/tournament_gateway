package tournaments

type TournamentGender = string

const (
	TournamentGenderMale   TournamentGender = "male"
	TournamentGenderFemale TournamentGender = "female"
)

type Tournament struct {
	Alias       string
	Name        string
	Description string
	Gender      TournamentGender
}
