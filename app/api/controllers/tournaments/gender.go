package tournaments

import (
	"errors"
	"tournament_gateway/models/tournaments"
)

func getGender(genderRequest string) (tournaments.TournamentGender, error) {
	switch genderRequest {
	case "male":
		return tournaments.TournamentGenderMale, nil
	case "female":
		return tournaments.TournamentGenderFemale, nil
	}
	return tournaments.TournamentGenderUndefined, errors.New("gender undefined")
}

