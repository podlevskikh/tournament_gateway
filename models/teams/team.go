package teams

import (
	"vollyemsk_tournament_gateway/models/gyms"
	"vollyemsk_tournament_gateway/models/users"
	"time"
)

type Team struct {
	ID             int
	Name           string
	Description    string
	FoundationDate time.Time
	ContactUsers   []users.User
	HomeGyms       []gyms.Gym
}
